import { useEffect, useState } from "react";
import { Stepper, Group, Button, TextInput, rem } from "@mantine/core";
import { IconAlertTriangle, IconLoader } from '@tabler/icons-react';
import classes from './classes.module.css';
import { useDebouncedValue } from '@mantine/hooks';

interface Props { onClose: (username: string) => void; }

export const LoginStepper = ({ onClose }: Props): JSX.Element => {
  const [active, setActive] = useState(0);
  const nextStep = () =>
  setActive((current) => (current < 3 ? current + 1 : current));
  const prevStep = () =>
  setActive((current) => (current > 0 ? current - 1 : current));
  const [value, setValue] = useState("");
  const [debounced] = useDebouncedValue(value, 500);
  const [isLoading, setLoading] = useState(false);
  const [error, setError] = useState<string | undefined>();
  useEffect(() => {
    if(debounced === "") {
      setError("Username cannot be empty");
      setLoading(false);
      return;
    }
    fetch(`http://localhost:3000/nameisunique/${debounced}`)
   .then(response => response.json())
   .then(data => {
      if (data === "false") {
        setError('Username already taken');
      } else if(data === "true") {
        setError(undefined);
      }
   }); 
    setLoading(false);
  }, [debounced]);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLoading(true);
    setValue(event.currentTarget.value);
  }

  const getInputIcon = () => {
    if(isLoading) {
      return <IconLoader />;
    }
    if (error) {
      return (
        <IconAlertTriangle
          stroke={1.5}
          style={{ width: rem(18), height: rem(18) }}
          className={classes.icon}
        />
      );
    }
  }
  return (
    <>
      <Stepper active={active} onStepClick={setActive}>
        <Stepper.Step>
          This is a minimalist clone of r/place, where you can draw pixels on a
          canvas. Select a color and click on the canvas to draw a pixel. You
          can color a pixel every 30 seconds.
        </Stepper.Step>
        <Stepper.Step>
          <TextInput
            error={error}
            label="Username"
            required
            value={value}
            onChange={handleChange}
            mt="md"
            autoComplete="nope"
            rightSection={getInputIcon()}
          />
        </Stepper.Step>
        <Stepper.Completed>
          Enjoy !
        </Stepper.Completed>
      </Stepper>

      <Group justify="center" mt="xl">
        <Button variant="default" onClick={prevStep}>
          Back
        </Button>
        {active === 2 && (
          <Button onClick={() => onClose(value)}>Enter game</Button>
        )}
        {active !== 2 && (
          <Button disabled={!!error && active !== 0} onClick={nextStep}>Next</Button>
        )}
      </Group>
    </>
  );
};

import { useEffect, useState } from "react";
import { Modal } from "@mantine/core";
import { PixelsPane } from "./components/PixelsPane/PixelsPane";
import { css } from "@emotion/css";
import { LoginStepper } from "./components/LoginStepper/LoginStepper";
import useWebSocket, { ReadyState } from "react-use-websocket";

function App() {
  const width = 50;
  const height = 50;
  const [pixelData, setPixelData] = useState<{x:number,y:number,color:string}[][]>(
    Array(width).fill(Array(height).fill("white"))
  );
  const [userName, setUserName] = useState("");
  const [opened, setOpened] = useState(true);
  const { sendMessage, lastMessage, readyState } = useWebSocket(
    `ws://localhost:3000/pixelwar/${userName}`
  );

  useEffect(() => {
    if (readyState === ReadyState.OPEN) {
      setPixelData(JSON.parse(lastMessage?.data).pixelMatrix);
    }
  }, [lastMessage]);

  const handleColorChange = (x: number, y: number) => {
    sendMessage(
      JSON.stringify({ action: "changePixel", payload: { x, y, color: "red" } })
    );
  };

  return (
    <>
      <div className={styles.wrapper}>
        <PixelsPane
          onDraw={handleColorChange}
          width={width}
          height={height}
          pixelColors={pixelData}
        />
      </div>
      <Modal
        opened={opened}
        onClose={() => setOpened(false)}
        title="Before entering the game"
      >
        <LoginStepper
          onClose={(name) => {
            setUserName(name);
            setOpened(false);
          }}
        />
      </Modal>
    </>
  );
}

export default App;

const styles = {
  wrapper: css`
    display: flex;
    flex-direction: column;
    align-items: center;
  `,
};

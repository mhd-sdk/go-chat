import { useEffect, useState } from "react";
import { Modal } from "@mantine/core";
import { PixelsPane } from "./components/PixelsPane/PixelsPane";
import { css } from "@emotion/css";
import { LoginStepper } from "./components/LoginStepper/LoginStepper";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { Chat } from "./components/Chat/Chat";

function App() {
  const width = 50;
  const height = 50;
  const [pixelData, setPixelData] = useState<{ x: number, y: number, color: string }[][]>(
    Array(width).fill(Array(height).fill("white"))
  );
  const [color, setColor] = useState("red");
  const [userName, setUserName] = useState("");
  const [opened, setOpened] = useState(true);
  const { sendMessage, lastMessage, readyState } = useWebSocket(
    `ws://localhost:3000/pixelwar/${userName}`
  );
  const [messages, setMessages] = useState<{ author: string, content: string }[]>([]);
  const [currentLoggedUsers, setCurrentLoggedUsers] = useState<string[]>([]);

  useEffect(() => {
    if (readyState === ReadyState.OPEN) {
      console.log(JSON.parse(lastMessage?.data));
      setPixelData(JSON.parse(lastMessage?.data).pixelMatrix);
      setMessages(JSON.parse(lastMessage?.data).messages);
      setCurrentLoggedUsers(JSON.parse(lastMessage?.data).loggedUsers);
    }
  }, [lastMessage]);

  const handleOnDraw = (x: number, y: number) => {
    sendMessage(
      JSON.stringify({ action: "changePixel", payload: { x, y, color } })
    );
  };

  const handleColorChange = (color: string) => {
    setColor(color);
  }

  const handleSendMessage = (message: string) => {
    sendMessage(
      JSON.stringify({ action: "addMessage", payload: message })
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.wrapper}>
        <PixelsPane
          onDraw={handleOnDraw}
          width={width}
          height={height}
          pixelColors={pixelData}
          disabled={readyState !== ReadyState.OPEN} onColorChange={handleColorChange} />
      </div>
      <div className={styles.wrapper}>
        <Chat messages={messages} currentLoggedUsers={currentLoggedUsers} onSendMessage={handleSendMessage} />
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
    </div>
  );
}

export default App;

const styles = {
  wrapper: css`
    display: flex;
    flex-direction: column;
    align-items: center;
    // fill the space left
    flex: 1;
  `,
  container: css`
    display: flex;
    flex-direction: row;
    padding: 50px 200px;
    // spac
    gap: 20px;
  `
};

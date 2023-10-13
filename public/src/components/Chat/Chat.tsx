import { css } from '@emotion/css';
import { TextInput } from '@mantine/core';

interface Message {
    author: string;
    content: string;
}
interface Props { messages: Message[], currentLoggedUsers: string[], onSendMessage: (message: string) => void }

export const Chat = ({ currentLoggedUsers, messages, onSendMessage }: Props): JSX.Element => {
    const getLoggedUsers = currentLoggedUsers.map((user) => user).join(", ");

    const renderMessages = () => {
        return messages.map((message, index) => {
            return (
                <div key={index}>
                    <div>{`${message.author}: ${message.content}`}</div>
                </div>)
        })
    }

    return (
        <div className={styles.container}>
            <div className={styles.scrollableChat}>

                Logged users :
                {getLoggedUsers}
                {renderMessages()}
            </div>
            <div className={styles.input}>
                <TextInput radius="xl" label="My input" onKeyDown={
                    (event) => {
                        if (event.key === "Enter") {
                            onSendMessage(event.currentTarget.value);
                            event.currentTarget.value = "";
                        }
                    }
                } />
            </div>
        </div>);
}

const styles = {
    container: css`
    width: 100%;
    
    height: 100%;
    display: flex;
    flex-direction: column;
    `,
    scrollableChat: css`
    // take all the space left
    flex: 1;
    overflow-y: auto;
    `,
    input: css`
    // stick to bottom
    bottom: 0;
    `

}
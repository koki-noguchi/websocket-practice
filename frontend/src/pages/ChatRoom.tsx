import React from 'react';
import { useWebSocket } from '../hooks/useWebSocket';
import ChatBox from '../components/ChatBox';
import ChatInput from '../components/ChatInput';

type Props = {
    roomName: string;
}

const ChatRoom: React.FC<Props> = ({ roomName }) => {
    const { messages, sendMessage } = useWebSocket(roomName);

    return (
        <div style={{ padding: 20 }}>
            <h2>Chat Room: {roomName}</h2>
            <ChatBox messages={messages} />
            <ChatInput onSend={sendMessage} />
        </div>
    );
};

export default ChatRoom;

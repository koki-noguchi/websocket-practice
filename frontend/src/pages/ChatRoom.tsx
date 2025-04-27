import React from 'react';
import { useWebSocket } from '../hooks/useWebSocket';
import ChatBox from '../components/ChatBox';
import ChatInput from '../components/ChatInput';

const ChatRoom = () => {
    const { messages, sendMessage } = useWebSocket('ws://localhost:8080/ws');

    return (
        <div style={{ padding: 20 }}>
            <h2>Chat Room</h2>
            <ChatBox messages={messages} />
            <ChatInput onSend={sendMessage} />
        </div>
    );
};

export default ChatRoom;

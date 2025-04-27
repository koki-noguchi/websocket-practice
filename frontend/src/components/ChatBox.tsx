import React from 'react';
import { Message } from '../types/Message';
import '../styles/ChatBox.css';

type ChatBoxProps = {
    messages: Message[];
};

const ChatBox: React.FC<ChatBoxProps> = ({ messages }) => {
    return (
        <div className="chat-box">
            {messages.map((msg, idx) => (
                <div key={idx} className="chat-message">{msg.text}</div>
            ))}
        </div>
    );
};

export default ChatBox;
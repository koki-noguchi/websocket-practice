import React, { useState } from 'react';
import '../styles/CahtInput.css';

type ChatInputProps = {
    onSend: (text: string) => void;
};

const ChatInput: React.FC<ChatInputProps> = ({ onSend }) => {
    const [input, setInput] = useState('');

    const handleSend = (e?: React.FormEvent<HTMLFormElement> | React.MouseEvent<HTMLButtonElement>) => {
        if (e) e.preventDefault();
        if (input.trim() !== '') {
            onSend(input);
            setInput('');
        }
    };

    const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            e.preventDefault();
            handleSend();
        }
    };

    return (
        <form onSubmit={handleSend} className="chat-input">
            <input
                type="text"
                value={input}
                onChange={(e) => setInput(e.target.value)}
                onKeyDown={handleKeyPress}
                placeholder="Type a message..."
            />
            <button type="submit">Send</button>
        </form>
    );
};

export default ChatInput;

import React, {useState} from 'react';

type Props = {
    onJoin: (roomName: string) => void;
};

const JoinRoom: React.FC<Props> = ({ onJoin }) => {
    const [input, setInput] = useState('');

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (input.trim()) {
            onJoin(input.trim());
        }
    };

    return (
        <div style={{ padding: 20 }}>
            <h2>Join a Room</h2>
            <form onSubmit={handleSubmit}>
                <input value={input} onChange={(e) => setInput(e.target.value)} />
                <button type="submit">Join</button>
            </form>
        </div>
    );
};

export default JoinRoom;

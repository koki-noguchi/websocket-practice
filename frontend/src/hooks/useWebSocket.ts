import { useEffect, useRef, useState } from 'react';
import { Message } from '../types/Message';

export const useWebSocket = (url: string) => {
    const ws = useRef<WebSocket | null>(null);
    const [messages, setMessages] = useState<Message[]>([]);

    useEffect(() => {
        ws.current = new WebSocket(url);

        ws.current.onmessage = (event) => {
            const msg = { text: event.data };
            setMessages(prev => [...prev, msg]);
        };

        ws.current.onerror = (err) => {
            console.error('WebSocket error:', err);
        };

        return () => {
            ws.current?.close();
        };
    }, [url]);

    const sendMessage = (text: string) => {
        if (ws.current && ws.current.readyState === WebSocket.OPEN) {
            ws.current.send(text);
        }
    };

    return { messages, sendMessage };
};

import { useEffect, useRef, useState } from 'react';
import { Message } from '../types/Message';

export const useWebSocket = (roomName: string) => {
    const ws = useRef<WebSocket | null>(null);
    const [messages, setMessages] = useState<Message[]>([]);

    useEffect(() => {
        const connect = () => {
            if (ws.current && (ws.current.readyState === WebSocket.OPEN || ws.current.readyState === WebSocket.CONNECTING)) {
                console.log("[WebSocket] Closing existing socket...");
                ws.current.close();
            }

            const socket = new WebSocket(`ws://localhost:8080/ws`);
            ws.current = socket;

            socket.onopen = () => {
                console.log("[WebSocket] Connected:", roomName);
                socket.send(roomName);
            };

            socket.onmessage = (event) => {
                try {
                    const data: Message = JSON.parse(event.data);
                    setMessages((prev) => [...prev, data]);
                } catch (err) {
                    console.error("[WebSocket] Parse error", event.data);
                }
            };

            socket.onclose = () => {
                console.log("[WebSocket] Closed");
            };

            socket.onerror = (error) => {
                console.error("[WebSocket] Error", error);
            };
        };

        connect();

        return () => {
            console.log("[WebSocket] Cleanup");
            ws.current?.close();
        };
    }, [roomName]);

    const sendMessage = (text: string) => {
        if (ws.current && ws.current.readyState === WebSocket.OPEN) {
            ws.current.send(text);
        }
    };

    return { messages, sendMessage };
};

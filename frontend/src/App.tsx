import React, {useState} from 'react';
import ChatRoom from './pages/ChatRoom';
import JoinRoom from "./pages/JoinRoom";

function App() {
  const [room, setRoom] = useState<string>('')

  if (!room) {
      return <JoinRoom onJoin={setRoom} />;
  }

  return <ChatRoom roomName={room} />;
}

export default App;

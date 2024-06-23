import { useEffect, useRef, useState } from "react";
import { socket } from "./socket";

interface IMessage {
  sentBy: string;
  content: string;
}

export default function Chat() {
  const [messages, setMessages] = useState<IMessage[]>([]);
  const chatRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleMessage = (message: IMessage) => {
      setMessages((prevMessages) => [...prevMessages, message]);
    };

    socket.on("message", handleMessage);

    return () => {
      socket.off("message");
    };
  }, []);

  useEffect(() => {
    if (chatRef.current) {
      chatRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);

  return (
    <div className=" overflow-y-auto p-2" style={{ height: "88vh" }}>
      {messages.map((msg) => (
        <div
          key={Math.random() + msg.sentBy + Date.now()}
          className="px-2 py-2"
        >
          <p className="text-sm text-gray-600 dark:text-gray-400 font-thin">
            {msg.sentBy || "Unknown"}: {msg.content}
          </p>
        </div>
      ))}
      <div ref={chatRef} />
    </div>
  );
}

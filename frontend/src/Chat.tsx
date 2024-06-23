import { Disclosure } from "@headlessui/react";
import { useState } from "react";

interface IMessage {
  sentBy: string;
  content: string;
}

export default function Chat() {
  const [messages, setMessages] = useState<IMessage[]>([
    { sentBy: "Alice", content: "Hello!" },
    { sentBy: "Bob", content: "Hi!" },
    { sentBy: "Alice", content: "How are you?" },
    { sentBy: "Bob", content: "Good!" },
    { sentBy: "Alice", content: "Great!" },
  ]);

  return (
    <Disclosure>
      {messages.map((msg) => (
        <div
          key={Math.random() + msg.sentBy + Date.now()}
          className="px-2 py-2"
        >
          <p className="text-sm text-gray-600 dark:text-gray-400 font-thin">
            {msg.sentBy}: {msg.content}
          </p>
        </div>
      ))}
    </Disclosure>
  );
}

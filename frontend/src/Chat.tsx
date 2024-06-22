import { Disclosure } from "@headlessui/react";

const msgs = [
  { sentBy: "Alice", content: "Hello!" },
  { sentBy: "Bob", content: "Hi!" },
  { sentBy: "Alice", content: "How are you?" },
  { sentBy: "Bob", content: "Good!" },
  { sentBy: "Alice", content: "Great!" },
];

export default function Chat() {
  return (
    <Disclosure>
      {msgs.map((msg) => (
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

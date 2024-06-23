import { useRef } from "react";
import { Disclosure, Button } from "@headlessui/react";
import { PaperAirplaneIcon } from "@heroicons/react/24/outline";
import { socket } from "./socket";

export default function Chat() {
  const inputRef = useRef<HTMLInputElement>(null);

  function sendMsg(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const target = event.target as typeof event.target & {
      msg: { value: string };
    };
    const msg = target.msg.value;
    socket.emit("message", msg);
    if (inputRef.current) inputRef.current.value = "";
  }

  return (
    <Disclosure>
      <form onSubmit={sendMsg}>
        <div className="flex flex-row justify-stretch">
          <input
            ref={inputRef}
            type="text"
            name="msg"
            placeholder="Type a message..."
            className="grow p-2 bg-gray-200 dark:bg-gray-800 dark:text-gray-200"
          />
          <Button
            type="submit"
            className="grow-0 rounded dark:bg-gray-700 py-2 px-4 text-sm text-white dark:data-[hover]:bg-gray-600 data-[active]:bg-sky-700"
          >
            <PaperAirplaneIcon className="h-6 w-auto" />
          </Button>
        </div>
      </form>
    </Disclosure>
  );
}

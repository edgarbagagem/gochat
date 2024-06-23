import { Disclosure } from "@headlessui/react";
import Chat from "./Chat";
import MessageField from "./MessageField";

export default function ChatGroup() {
  return (
    <Disclosure>
      <div className="mx-1 border border-gray-400 dark:border-gray-800 rounded-md bg-gray-100 dark:bg-gray-900">
        <Chat />
        <div className="sticky bottom-0 w-full bg-gray-100 dark:bg-gray-900">
          <MessageField />
        </div>
      </div>
    </Disclosure>
  );
}

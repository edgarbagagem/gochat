import { Disclosure } from "@headlessui/react";
import { useEffect, useState } from "react";
import { socket } from "./socket";
import { onlyUnique } from "./utils/unique";

export default function OnlineUsers() {
  const [onlineUsers, setOnlineUsers] = useState<string[]>([]);

  useEffect(() => {
    const handleEmitUsers = (
      users: { socketID: string; username: string }[]
    ) => {
      setOnlineUsers(users.map((user) => user.username).filter(onlyUnique));
    };

    socket.on("updateUsers", handleEmitUsers);

    // Cleanup listener on component unmount
    return () => {
      socket.off("updateUsers", handleEmitUsers);
    };
  }, []);

  return (
    <Disclosure>
      <div className="mx-1 border border-gray-400 dark:border-gray-800 rounded-md bg-gray-100 dark:bg-gray-900 max-h-screen">
        <div className="rounded-md py-2 px-2 bg-gray-200 dark:bg-gray-950">
          <p className="text-base text-gray-800 dark:text-gray-200 font-thin">
            Online Users
          </p>
        </div>

        <div className="px-2 py-2">
          <ul>
            {onlineUsers.map((user, idx) => (
              <li key={idx}>
                <a
                  href="#"
                  className="block text-sm text-gray-600 dark:text-gray-400 font-thin"
                >
                  {user}
                </a>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </Disclosure>
  );
}

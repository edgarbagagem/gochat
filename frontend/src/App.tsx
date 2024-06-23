import NavBar from "./NavBar";
import OnlineUsers from "./OnlineUsers";
import ChatGroup from "./ChatGroup";
import { isAuthenticated } from "./utils/auth";
import { Link } from "react-router-dom";
import { useEffect } from "react";
import { socket } from "./socket";

function App() {
  useEffect(() => {
    if (isAuthenticated()) {
      socket.connect();
      socket.emit("register", sessionStorage.getItem("username"));
    }

    return () => {
      if (!socket.disconnected) socket.disconnect();
    };
  }, []);

  return (
    <>
      <div className="dark:bg-gray-950 bg-white flex flex-col h-screen">
        <NavBar />
        <div className="flex flex-row flex-wrap justify-stretch items-stretch mt-1 ">
          {isAuthenticated() ? (
            <>
              <div className="basis-1/4 h-fit max-h-screen">
                <OnlineUsers />
              </div>
              <div className="basis-3/4 h-fit max-h-screen">
                <ChatGroup />
              </div>
            </>
          ) : (
            <div className="flex flex-col items-center justify-center w-full h-full text-center">
              <p className="text-gray-700 dark:text-gray-300 text-lg">
                You need to log in or create an account to access this content.
              </p>
              <div className="mt-4">
                <Link
                  to="/login"
                  className="px-4 py-2 rounded-md text-sm font-medium text-white bg-blue-500 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700"
                >
                  Login
                </Link>
                <Link
                  to="/login"
                  className="ml-4 px-4 py-2 rounded-md text-sm font-medium text-white bg-green-500 hover:bg-green-600 dark:bg-green-600 dark:hover:bg-green-700"
                >
                  Create Account
                </Link>
              </div>
            </div>
          )}
        </div>
      </div>
    </>
  );
}

export default App;

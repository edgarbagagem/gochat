import NavBar from "./NavBar";
import OnlineUsers from "./OnlineUsers";
import ChatGroup from "./ChatGroup";

function App() {
  return (
    <>
      <div className="dark:bg-gray-950 bg-white h-screen flex flex-col">
          <NavBar/>
        <div className="flex flex-row flex-wrap justify-stretch items-stretch h-full mt-1">
          <div className="basis-1/4">
            <OnlineUsers/>
          </div>
          <div className="basis-3/4">
            <ChatGroup/>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;

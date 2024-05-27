import NavBar from "./NavBar";
import OnlineUsers from "./OnlineUsers";

function App() {
  return (
    <>
      <div className="dark:bg-gray-950 bg-white">
        <NavBar/>
        <OnlineUsers/>
      </div>
    </>
  );
}

export default App;

const { io } = require("socket.io-client");
const fetch = require("node-fetch");

const socket = io(process.env.WEBSOCKET_SERVER);

const apiUrl = "https://official-joke-api.appspot.com/random_joke";

socket.on("connect", async () => {
  try {
    const response = await fetch(apiUrl);
    const data = await response.json();

    const message = `${data.setup} - ${data.punchline}`;

    socket.emit("message", message);

    socket.close();
  } catch (error) {
    console.error("Error fetching data or sending message:", error);
  }
});

socket.on("connect_error", (error) => {
  console.error("Socket.IO connection error:", error);
});

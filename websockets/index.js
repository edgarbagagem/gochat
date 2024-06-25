import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";

const app = express();
const server = createServer(app);
const io = new Server(server, {
  cors: {
    origin: process.env.CORS_ORIGIN || "http://localhost:5173",
  },
});
app.get("/", (req, res) => {
  res.send("<h1>Gochat Websocket Server</h1>");
});

let connectedUsers = [];

io.on("connection", (socket) => {
  console.log("a user connected", socket.id);

  const emitUserList = () => {
    io.emit("updateUsers", connectedUsers);
  };

  socket.on("disconnect", () => {
    connectedUsers = connectedUsers.filter((user) => user.socket != socket.id);
    console.log("user disconnected");
    emitUserList();
  });
  socket.on("register", (username) => {
    connectedUsers.push({ socket: socket.id, username: username });
    console.log("Registering user ", username);
    emitUserList();
  });

  socket.on("message", (message) => {
    console.log("Message received: ", message);
    io.emit("message", {
      sentBy: connectedUsers.find((user) => user.socket === socket.id)
        ?.username,
      content: message,
    });
  });
});

server.listen(3000, () => {
  console.log("server running at port 3000");
});

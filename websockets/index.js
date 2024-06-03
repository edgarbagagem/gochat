const express = require("express");
const http = require("http");
const { Server } = require("socket.io");

const app = express();
const server = http.createServer(app);
const io = new Server(server);

const connectedUsers = {};

io.on("connection", (socket) => {
  console.log("a user connected", socket.id);

  // Add the user to the connected users list
  connectedUsers[socket.id] = { id: socket.id };

  // Notify all clients about the updated user list
  io.emit("userList", Object.values(connectedUsers));

  // Broadcast a message to all clients
  socket.on("chatMessage", (msg) => {
    io.emit("chatMessage", { id: socket.id, msg });
  });

  // Handle user disconnect
  socket.on("disconnect", () => {
    console.log("a user disconnected", socket.id);
    delete connectedUsers[socket.id];
    io.emit("userList", Object.values(connectedUsers));
  });
});

server.listen(3000, () => {
  console.log("listening on *:3000");
});

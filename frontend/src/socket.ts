/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import { io } from "socket.io-client";

// "undefined" means the URL will be computed from the `window.location` object
const URL = import.meta.env.VITE_WS_URL;

// eslint-disable-next-line @typescript-eslint/no-unsafe-argument
export const socket = io(URL, {
  autoConnect: false,
});

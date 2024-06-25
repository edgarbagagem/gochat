import axios from "axios";

const instance = axios.create({
  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    "Content-Type": "application/json",
    Authorization: sessionStorage.getItem("jwtToken") || "",
  },
});

instance.interceptors.request.use(
  (config) => {
    const token = sessionStorage.getItem("jwtToken");
    if (token) {
      config.headers.Authorization = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default instance;

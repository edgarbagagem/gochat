import { useState } from "react";
import axios from "../api/axios";
import { AxiosError, AxiosResponse } from "axios";
import { LoginResponse } from "../types/auth";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const response: AxiosResponse<LoginResponse> = await axios.post(
        "/login",
        { username, password }
      );
      sessionStorage.setItem("jwtToken", response.data.token);
      sessionStorage.setItem("username", response.data.username);
      toast.success("Login successful");
      navigate("/");
    } catch (error) {
      if (error instanceof AxiosError) {
        const err = error as AxiosError;
        toast.error(error.message);
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        //@ts-expect-error
        toast.error(`Login failed, ${err.response?.data?.error}`);
      }
      console.error("Login failed:", error);
    }
  };

  const handleRegister = async () => {
    try {
      await axios.post("/register", { username, password });
      toast.success("Registration successfull! You can login now.");
    } catch (error) {
      if (error instanceof AxiosError) {
        const err = error as AxiosError;
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        //@ts-expect-error
        toast.error(`Registration failed, ${err.response?.data.error}`);
      }
      console.error("Login failed:", error);
    }
  };

  return (
    <div>
      <div className="flex items-center justify-center min-h-screen bg-gray-900">
        <div className="bg-gray-800 p-8 rounded-lg shadow-md">
          <h2 className="text-center text-2xl font-bold text-white mb-6">
            Welcome to GoChat
          </h2>
          <div className="mb-4">
            <input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full p-2 mb-2 rounded bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full p-2 rounded bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
          </div>
          <div className="flex justify-between">
            <button
              className="w-1/2 mr-2 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              // eslint-disable-next-line @typescript-eslint/no-misused-promises
              onClick={handleLogin}
            >
              Login
            </button>
            <button
              className="w-1/2 ml-2 py-2 bg-green-500 text-white rounded hover:bg-green-600"
              // eslint-disable-next-line @typescript-eslint/no-misused-promises
              onClick={handleRegister}
            >
              Sign up
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;

import {
  Disclosure,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  Transition,
} from "@headlessui/react";
import { isAuthenticated } from "./utils/auth";
import { useEffect, useState } from "react";
import axios from "./api/axios";
import { GetUserResponse } from "./types/auth";
import { AxiosError } from "axios";
import toast from "react-hot-toast";

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}

export default function NavBar() {
  const [username, setUsername] = useState<string>("");
  const [userPhoto, setUserPhoto] = useState<string>("");

  useEffect(() => {
    const username = sessionStorage.getItem("username");
    setUsername(username || "");
    if (!username) return;

    axios
      .get<GetUserResponse>(`/profile-photo/${username}`)
      .then((resp) => {
        const user: GetUserResponse = resp.data;
        setUserPhoto(user.photo);
      })
      .catch((error) => {
        const axiosError = error as AxiosError;
        if (axiosError.response) {
          toast.error(
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            //@ts-expect-error
            `Error getting user info (photo): ${axiosError.response?.data?.error}`
          );
          console.error(
            "Error getting user info (photo):",
            axiosError.response.data
          );
        } else {
          toast.error("Error getting user info (photo):");
          console.error("Error getting user info (photo):", axiosError.message);
        }
      });
  }, []);

  const handleSignOut = () => {
    sessionStorage.clear();
  };

  return (
    <Disclosure as="nav" className="bg-white dark:bg-gray-800">
      {() => (
        <>
          <div className="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
            <div className="relative flex h-16 items-center justify-between">
              <div className="flex flex-1 items-center justify-start sm:items-stretch sm:justify-start">
                <div className="flex flex-shrink-0 items-center">
                  <img
                    className="h-8 w-auto"
                    src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
                    alt="GoChat"
                  />
                </div>
              </div>
              <div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
                {isAuthenticated() && (
                  <Menu as="div" className="relative ml-3">
                    <div className="flex items-center space-x-3">
                      <span className="text-white">{username}</span>
                      <MenuButton className="relative flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-gray-800 dark:focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
                        <span className="absolute -inset-1.5" />
                        <span className="sr-only">Open user menu</span>
                        <img
                          className="h-8 w-8 rounded-full"
                          src={userPhoto}
                          alt=""
                        />
                      </MenuButton>
                    </div>
                    <Transition
                      enter="transition ease-out duration-100"
                      enterFrom="transform opacity-0 scale-95"
                      enterTo="transform opacity-100 scale-100"
                      leave="transition ease-in duration-75"
                      leaveFrom="transform opacity-100 scale-100"
                      leaveTo="transform opacity-0 scale-95"
                    >
                      <MenuItems className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                        <MenuItem>
                          {({ active }) => (
                            <a
                              href="/upload-photo"
                              className={classNames(
                                active ? "bg-gray-100" : "",
                                "block px-4 py-2 text-sm text-gray-700"
                              )}
                            >
                              Change your photo
                            </a>
                          )}
                        </MenuItem>
                        <MenuItem>
                          {({ active }) => (
                            <a
                              href="/"
                              onClick={handleSignOut}
                              className={classNames(
                                active ? "bg-gray-100" : "",
                                "block px-4 py-2 text-sm text-gray-700"
                              )}
                            >
                              Sign out
                            </a>
                          )}
                        </MenuItem>
                      </MenuItems>
                    </Transition>
                  </Menu>
                )}
              </div>
            </div>
          </div>
        </>
      )}
    </Disclosure>
  );
}

import { Disclosure } from '@headlessui/react';

const onlineUsers = [{id: 1, name: 'Alice'}, {id: 2, name: 'Bob'}, {id: 3, name: 'Charlie'}, {id: 4, name: 'David'}, {id: 5, name: 'Eve'}]

export default function OnlineUsers() {
return (
    <Disclosure>
        <div className="mx-1 border border-gray-400 dark:border-gray-800 rounded-md bg-gray-100 dark:bg-gray-900 h-full">
            <div className="rounded-md py-2 px-2 bg-gray-200 dark:bg-gray-950">
                <p className="text-base text-gray-800 dark:text-gray-200 font-thin">Online Users</p>
            </div>
            
            <div className="px-2 py-2">
                <ul>
                    {onlineUsers.map((user) => (
                        <li key={user.id}>
                        <a href="#" className="block text-sm text-gray-600 dark:text-gray-400 font-thin">
                            {user.name}
                        </a>
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    </Disclosure>
)
}

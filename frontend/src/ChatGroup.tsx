import { Disclosure } from '@headlessui/react';
import Chat from './Chat';
import MessageField from './MessageField';

export default function ChatGroup() {
return (
    <Disclosure>
        <div className="flex flex-col relative mx-1 border border-gray-400 dark:border-gray-800 rounded-md bg-gray-100 dark:bg-gray-900 h-full">
            <Chat/>
            <div className='absolute bottom-0 w-full'>
               <MessageField/> 
            </div>
            
        </div>
    </Disclosure>
)
}

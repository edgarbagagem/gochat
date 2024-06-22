// src/components/UploadProfilePhotoPage.tsx

import React, { useRef, useState } from "react";
import toast from "react-hot-toast";

const UploadProfilePhoto: React.FC = () => {
  const [file, setFile] = useState<File | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setFile(e.target.files[0]);
    }
  };

  const handleClear = () => {
    setFile(null);
    if (inputRef.current) {
      inputRef.current.value = "";
    }
  };

  const handleUpload = () => {
    if (!file) {
      toast.error("Please select a file");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-900">
      <div className="bg-gray-800 p-8 rounded-lg shadow-md">
        <h2 className="text-center text-2xl font-bold text-white mb-6">
          Upload Profile Photo
        </h2>
        <div className="mb-4">
          <input
            ref={inputRef}
            type="file"
            accept="image/*"
            onChange={handleFileChange}
            className="w-full p-2 mb-2 rounded bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-600"
          />
          {file && (
            <div className="mt-4">
              <img
                src={URL.createObjectURL(file)}
                alt="Preview"
                className="h-32 w-32 object-cover rounded-full mx-auto"
              />
              <button
                onClick={handleClear}
                className="mt-2 w-full py-2 bg-red-500 text-white rounded hover:bg-red-600"
              >
                Clear
              </button>
            </div>
          )}
        </div>

        <button
          onClick={handleUpload}
          className="w-full py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          Upload
        </button>
      </div>
    </div>
  );
};

export default UploadProfilePhoto;

import React from 'react';
import Image from 'next/image';
import logoImage from '@/app/logo.png';

const LargeScreenMessage: React.FC = () => {
  return (
    <div className="flex h-screen w-screen items-center justify-center bg-gray-100">
      <div className="text-center max-w-2xl mx-auto p-8 bg-white rounded-lg shadow-lg">
        <h1 className="text-3xl font-bold mb-4 text-gray-800">
          Welcome to MealMate
        </h1>
        <p className="text-xl text-gray-600 mb-6">
          For the best experience, please visit our site using a mobile device.
        </p>
        <div className="flex justify-center mb-6">
          <Image
            src={logoImage}
            alt="MealMate Logo"
            width={120}
            height={120}
            className="rounded-full"
          />
        </div>
        <p className="mt-6 text-gray-500">
          Our mobile version offers optimized features and a seamless browsing
          experience for planning and enjoying your meals.
        </p>
      </div>
    </div>
  );
};

export default LargeScreenMessage;

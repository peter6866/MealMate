'use client';

import { title } from '@/components/primitives';

import { Input } from '@nextui-org/react';
import { Button } from '@nextui-org/button';
import { Image } from '@nextui-org/react';
import NextImage from 'next/image';
import MenuItemList from './menuItemList';

const filterCategories = ['All', 'Italian', 'Indian', 'Salads', 'Pizza'];

export default function HomePage() {
  return (
    <div className="min-h-screen bg-content1 p-4">
      {/* <div className="mb-4">
        <Input
          contentLeft={
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-5 h-5"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z"
              />
            </svg>
          }
          placeholder="Search for dishes..."
          width="100%"
        />
      </div> */}

      <div className="mb-4 overflow-x-auto whitespace-nowrap">
        {filterCategories.map((category) => (
          <Button key={category} size="sm" className="mr-2">
            {category}
          </Button>
        ))}
      </div>

      <MenuItemList />
    </div>
  );
}

import { Button } from '@nextui-org/button';
import React, { Suspense } from 'react';
import MenuItemList from './menuItemList';
import HomeSkeleton from '@/components/homeSkeleton';

const filterCategories = ['All', 'Italian', 'Indian', 'Salads', 'Pizza'];

export default function HomePage() {
  return (
    <div className="min-h-screen bg-content1 p-4">
      <div className="mb-4 overflow-x-auto whitespace-nowrap">
        {filterCategories.map((category) => (
          <Button key={category} size="sm" className="mr-2">
            {category}
          </Button>
        ))}
      </div>

      <Suspense fallback={<HomeSkeleton />}>
        <MenuItemList />
      </Suspense>
    </div>
  );
}

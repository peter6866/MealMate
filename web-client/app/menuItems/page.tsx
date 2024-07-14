import { Button } from '@nextui-org/button';
import React, { Suspense } from 'react';
import MenuItemList from './menuItemList';
import HomeSkeleton from '@/components/homeSkeleton';
import Filter from '@/components/Filter';

interface searchParamsProps {
  category: string;
  spiceLevel: string;
  alcoholContent: string;
}

export default function HomePage({
  searchParams,
}: {
  searchParams: searchParamsProps;
}) {
  return (
    <div className="min-h-screen bg-content1 p-4">
      <Filter />

      <Suspense
        fallback={<HomeSkeleton />}
        key={`${searchParams.category}${searchParams.alcoholContent}${searchParams.spiceLevel}`}
      >
        <MenuItemList filter={searchParams} />
      </Suspense>
    </div>
  );
}

import React, { Suspense } from 'react';
import MenuItemList from './menuItemList';
import HomeSkeleton from '@/components/homeSkeleton';
import Filter from '@/components/Filters/Filter';
import { PlusIcon } from '@heroicons/react/24/outline';
import { Button } from '@nextui-org/button';
import Link from 'next/link';
import { cookies } from 'next/headers';
import Search from '@/components/Search';

interface searchParamsProps {
  category: string;
  spiceLevel: string;
  alcoholContent: string;
  query: string;
}

export default async function HomePage({
  searchParams,
}: {
  searchParams: searchParamsProps;
}) {
  const cookieStore = cookies();
  const isChef = cookieStore.get('isChef')?.value === 'true';

  return (
    <div className="min-h-screen bg-content2 dark:bg-content1 p-4 relative">
      <div className="flex justify-between items-center gap-4 mb-4">
        <Search />
        <Filter />
      </div>

      <Suspense
        fallback={<HomeSkeleton />}
        key={`${searchParams.category}${searchParams.alcoholContent}${searchParams.spiceLevel}${searchParams.query}`}
      >
        <MenuItemList filter={searchParams} />
      </Suspense>

      {isChef && (
        <Link href="menuItems/create" passHref>
          <Button
            isIconOnly
            color="primary"
            aria-label="Create"
            className="fixed bottom-[6.5rem] right-10 z-50 shadow-lg rounded-full w-12 h-12 bg-mainLight dark:bg-mainDark"
          >
            <PlusIcon className="h-8 w-8" />
          </Button>
        </Link>
      )}
    </div>
  );
}

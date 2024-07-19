import React, { Suspense } from 'react';
import MenuItemList from './menuItemList';
import HomeSkeleton from '@/components/homeSkeleton';
import Filter from '@/components/Filters/Filter';
import { PlusIcon } from '@heroicons/react/24/outline';
import { Button } from '@nextui-org/button';
import Link from 'next/link';
import { cookies } from 'next/headers';
import axios from 'axios';

interface searchParamsProps {
  category: string;
  spiceLevel: string;
  alcoholContent: string;
}

interface User {
  isChef: boolean;
}

export default async function HomePage({
  searchParams,
}: {
  searchParams: searchParamsProps;
}) {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;
  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/auth/getUser`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const user = response.data as User;

  return (
    <div className="min-h-screen bg-content1 p-4 relative">
      <Filter />

      <Suspense
        fallback={<HomeSkeleton />}
        key={`${searchParams.category}${searchParams.alcoholContent}${searchParams.spiceLevel}`}
      >
        <MenuItemList filter={searchParams} />
      </Suspense>

      {user.isChef && (
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

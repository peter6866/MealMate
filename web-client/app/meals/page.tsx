import { Button } from '@nextui-org/button';

import Link from 'next/link';
import { PlusIcon } from '@heroicons/react/24/outline';
import MealsList from '@/components/Meals/MealsList';
import MealsFilter from '@/components/Meals/MealsFilter';
import axios from 'axios';
import { cookies } from 'next/headers';

import { Meal } from '@/types/meals';

export const revalidate = 0;

interface MealsParams {
  startDate: string;
  endDate: string;
}

export default async function MealsPage({
  searchParams,
}: {
  searchParams: MealsParams;
}) {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  let { startDate, endDate } = searchParams;

  const today = new Date();
  const tmr = new Date(today);
  tmr.setDate(tmr.getDate() + 1);

  if (!startDate || !endDate) {
    startDate = today.toISOString().split('T')[0];
    endDate = tmr.toISOString().split('T')[0];
  }

  const response = await axios.post(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/meals/date`,
    {
      startDate,
      endDate,
    },
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const { meals }: { meals: Meal[] } = response.data;

  return (
    <div className="flex flex-col bg-content1 dark:bg-content2 min-h-[95vh] pt-4 relative">
      <header className="mb-4">
        <h1 className="text-2xl font-bold text-center text-default-800 mb-4">
          Your Meals
        </h1>

        <MealsFilter />
      </header>

      <div className="bg-content2 dark:bg-content1 flex flex-col flex-1 p-4 rounded-t-2xl">
        <MealsList meals={meals} />

        <Link href="#" passHref className="mt-3 flex">
          <Button color="primary" variant="light" size="md">
            View All Meals
          </Button>
        </Link>
      </div>

      <Link href="/meals/create" passHref>
        <Button
          color="primary"
          aria-label="Create"
          className="fixed bottom-20 right-4 z-50 shadow-lg rounded-full bg-mainLight dark:bg-mainDark"
          startContent={<PlusIcon className="h-6 w-6" />}
        >
          Log a meal
        </Button>
      </Link>
    </div>
  );
}

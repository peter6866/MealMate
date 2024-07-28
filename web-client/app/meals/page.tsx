import { Button } from '@nextui-org/button';

import Link from 'next/link';
import { PlusIcon } from '@heroicons/react/24/outline';
import MealsList from '@/components/Meals/MealsList';
import MealsFilter from '@/components/Meals/MealsFilter';
import axios from 'axios';
import { cookies } from 'next/headers';

export default async function MealsPage() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  const today = new Date();
  const tmr = new Date(today);
  tmr.setDate(tmr.getDate() + 1);

  const startDate = today.toISOString().split('T')[0];
  const endDate = tmr.toISOString().split('T')[0];

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

  const { meals } = response.data;

  return (
    <div className="bg-content2 dark:bg-content1 min-h-[95vh] p-4 relative">
      <header className="mb-6">
        <h1 className="text-2xl font-bold text-center text-default-800 mb-4">
          Your Meals
        </h1>

        <MealsFilter />
      </header>

      <MealsList meals={meals} />

      <Link href="#" passHref className="mt-3 flex">
        <Button color="primary" variant="light" size="md">
          View All Meals
        </Button>
      </Link>

      <Link href="#" passHref>
        <Button
          color="primary"
          aria-label="Create"
          className="fixed bottom-20 right-4 z-50 shadow-lg rounded-full bg-mainLight dark:bg-mainDark"
          startContent={<PlusIcon className="h-6 w-6" />}
        >
          Log your meal
        </Button>
      </Link>
    </div>
  );
}

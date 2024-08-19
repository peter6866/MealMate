import { Card } from '@nextui-org/card';
import { Avatar, AvatarGroup } from '@nextui-org/avatar';
import Image from 'next/image';
import moment from 'moment';

import { Meal, MealItem } from '@/types/meals';
import MealsDropdown from './MealsDropdown';

export default function MealsList({ meals }: { meals: Meal[] }) {
  if (!meals) {
    return (
      <Card className="p-8 text-center" shadow="sm">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-12 h-12 mx-auto mb-4 text-default-400"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M11.35 3.836c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75 2.25 2.25 0 00-.1-.664m-5.8 0A2.251 2.251 0 0113.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m8.9-4.414c.376.023.75.05 1.124.08 1.131.094 1.976 1.057 1.976 2.192V16.5A2.25 2.25 0 0118 18.75h-2.25m-7.5-10.5H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V18.75m-7.5-10.5h6.375c.621 0 1.125.504 1.125 1.125v9.375m-8.25-3l1.5 1.5 3-3.75"
          />
        </svg>

        <h2 className="text-xl font-semibold mb-2 text-default-700">
          No Meals Logged
        </h2>
        <p className="text-default-500">
          You haven&apos;t logged any meals for the selected date. Start logging
          your meals now!
        </p>
      </Card>
    );
  }

  return (
    <div className="space-y-4">
      {meals.map((meal: Meal) => (
        <Card
          key={meal.id}
          className="bg-content1 dark:bg-content2 border-none"
          shadow="sm"
        >
          <div className="p-3">
            <div className="flex justify-between items-center mb-1">
              <h2 className="text-lg font-semibold text-default-700">
                {meal.mealType}
              </h2>
              <MealsDropdown mealId={meal.id} />
            </div>
            <div className="relative h-60 mb-2">
              <Image
                src={meal.photoURL}
                alt={meal.mealType}
                fill
                className="rounded-lg object-cover"
              />

              <AvatarGroup className="absolute bottom-4 right-4">
                <Avatar
                  size="sm"
                  src="https://lh3.googleusercontent.com/a/ACg8ocKmVcEZmM7sgNr-LQbnfswtwOsIqq5haBoLsivizWQpclxI6mU=s96-c"
                  isBordered
                />
                <Avatar
                  size="sm"
                  src="https://lh3.googleusercontent.com/a/ACg8ocKc8PIW9NYpaVe7iCPbrFkCTMSJkmVuFeezcaeAXyzEjmd2JEcr=s96-c"
                  isBordered
                />
              </AvatarGroup>
            </div>
            <h3 className="text-md font-semibold mb-2 text-default-700">
              What&apos;s on the plate:
            </h3>
            <div className="grid grid-cols-2 gap-y-1">
              {meal.items.map((item: MealItem) => (
                <div key={item.id} className="flex items-center mb-1">
                  <Image
                    src={item.imageUrl}
                    alt={item.name}
                    width={45}
                    height={45}
                    className="rounded-lg object-cover"
                  />
                  <p className="ml-2 text-sm text-default-600">{item.name}</p>
                </div>
              ))}
            </div>
            <p className="text-xs text-default-400 mt-2">
              {moment(meal.mealDate.split('T')[0]).format('MMMM D, YYYY')}
            </p>
          </div>
        </Card>
      ))}
    </div>
  );
}

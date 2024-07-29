import { Card } from '@nextui-org/card';
import { Avatar, AvatarGroup } from '@nextui-org/avatar';
import Image from 'next/image';

import { Meal, MealItem } from '@/types/meals';

export default function MealsList({ meals }: { meals: Meal[] }) {
  if (!meals) {
    return (
      <p className="text-center text-lg mt-2 text-default-700">
        Start logging your meals for today!
      </p>
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
          <div className="p-4">
            <div className="flex justify-between items-center mb-3">
              <h2 className="text-lg font-semibold text-default-700">
                {meal.mealType}
              </h2>
              <AvatarGroup>
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
            <div className="relative h-52 mb-3">
              <Image
                src={meal.photoURL}
                alt={meal.mealType}
                fill
                className="rounded-lg object-cover"
              />
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
                    width={50}
                    height={50}
                    className="rounded-lg object-cover"
                  />
                  <p className="ml-2 text-default-600">{item.name}</p>
                </div>
              ))}
            </div>
          </div>
        </Card>
      ))}
    </div>
  );
}

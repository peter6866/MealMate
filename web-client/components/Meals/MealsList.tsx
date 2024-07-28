import { Card } from '@nextui-org/card';
import Image from 'next/image';

export default function MealsList({ meals }: { meals: any }) {
  if (!meals) {
    return <p>No meals found</p>;
  }

  return (
    <div className="space-y-4">
      {meals.map((meal: any) => (
        <Card
          key={meal.id}
          className="bg-content1 dark:bg-content2 border-none"
          shadow="none"
        >
          <div className="p-4">
            <div className="flex justify-between items-center mb-2">
              <h2 className="text-lg font-semibold text-default-700">
                {meal.mealType}
              </h2>
              <p className="text-sm text-default-600">
                {new Date(meal.mealDate).toLocaleTimeString([], {
                  hour: '2-digit',
                  minute: '2-digit',
                })}
              </p>
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
            <ul className="space-y-1">
              {meal.items.map((item: any) => (
                <li
                  key={item.id}
                  className="flex items-center text-default-600 text-sm"
                >
                  <span className="mr-2">🍽️</span>
                  {item.name}
                </li>
              ))}
            </ul>
          </div>
        </Card>
      ))}
    </div>
  );
}

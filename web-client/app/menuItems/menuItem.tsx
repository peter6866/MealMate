import { PlusCircleIcon } from '@heroicons/react/24/outline';
import { Card, CardFooter } from '@nextui-org/card';
import NextImage from 'next/image';
import Link from 'next/link';

interface Dish {
  id: string;
  name: string;
  categoryName: string;
  imageUrl: string;
}

export default function MenuItem({ dish }: { dish: Dish }) {
  return (
    <Card className="border-none" shadow="md">
      <Link href={`/menuItems/${dish.id}`} className="relative aspect-square">
        <NextImage
          src={dish.imageUrl}
          alt={dish.name}
          fill
          className="object-cover rounded-xl"
          sizes="(min-width: 640px) 50vw, 100vw"
        />
      </Link>
      <CardFooter className="justify-between">
        <p className="font-medium text-default-800">{dish.name}</p>
        <PlusCircleIcon className="w-6 h-6 text-mainLight" />
      </CardFooter>
    </Card>
  );
}

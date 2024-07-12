'use client';

import { Button } from '@nextui-org/button';
import { Card, CardBody } from '@nextui-org/card';
import NextImage from 'next/image';

export default function MenuItem({ dish }: { dish: any; index: number }) {
  return (
    <div className="flex flex-col">
      <Card>
        <CardBody>
          <div className="relative aspect-square">
            <NextImage
              src={dish.image}
              alt={dish.name}
              fill
              className="object-cover rounded-xl"
            />
          </div>
          <p className="mt-2 text-center font-semibold">{dish.name}</p>
        </CardBody>
      </Card>
    </div>
  );
}

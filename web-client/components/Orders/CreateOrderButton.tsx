'use client';

import React, { useMemo, useTransition } from 'react';
import { Button } from '@nextui-org/button';
import createOrder from './CreateOrderAction';

interface CartItem {
  id: string;
  name: string;
  categoryId: string;
  imageUrl: string;
  alcoholContent: string;
  createdBy: string;
  createdAt: string;
  updatedAt: string;
}

interface Item {
  id: string;
  name: string;
  imageUrl: string;
}

export default function CreateOrderButton({
  cartItems,
}: {
  cartItems: CartItem[];
}) {
  const items: Item[] = useMemo(() => {
    return cartItems.map(({ id, name, imageUrl }) => ({ id, name, imageUrl }));
  }, [cartItems]);
  const [isPending, startTransition] = useTransition();

  function handleClick() {
    startTransition(() => createOrder(items));
  }

  return (
    <div className="fixed bottom-16 left-0 right-0 p-4 bg-content1 dark:bg-content1">
      <Button
        color="primary"
        fullWidth
        className="bg-mainLight dark:bg-mainDark text-lg"
        onClick={handleClick}
        isLoading={isPending}
      >
        Create Order
      </Button>
    </div>
  );
}

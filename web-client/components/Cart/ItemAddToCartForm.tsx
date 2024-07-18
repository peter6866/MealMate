'use client';

import { Button } from '@nextui-org/button';
import { useFormStatus } from 'react-dom';
import { addToCart } from './AddToCartAction';

function AddToCartButton() {
  const { pending } = useFormStatus();
  return (
    <Button
      type="submit"
      color="primary"
      size="lg"
      className="w-full bg-mainLight dark:bg-mainDark text-white text-lg py-3"
      isLoading={pending}
    >
      Add to Cart
    </Button>
  );
}

export default function ItemAddToCartForm({
  menuItemId,
}: {
  menuItemId: string;
}) {
  return (
    <form action={addToCart}>
      <input type="hidden" name="menuItemId" value={menuItemId} />
      <AddToCartButton />
    </form>
  );
}

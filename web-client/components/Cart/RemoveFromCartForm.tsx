'use client';

import { TrashIcon } from '@heroicons/react/24/outline';
import { Button } from '@nextui-org/button';
import { useFormStatus } from 'react-dom';
import { removeFromCart } from './RemoveFromCartAction';

function DeleteButton() {
  const { pending } = useFormStatus();
  return (
    <Button
      isIconOnly
      color="danger"
      variant="light"
      type="submit"
      isLoading={pending}
    >
      <TrashIcon className="h-6 w-6" />
    </Button>
  );
}

export default function RemoveFromCart({ menuItemId }: { menuItemId: string }) {
  return (
    <form action={removeFromCart}>
      <input type="hidden" name="menuItemId" value={menuItemId} />
      <DeleteButton />
    </form>
  );
}

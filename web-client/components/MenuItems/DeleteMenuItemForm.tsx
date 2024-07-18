'use client';

import { Button } from '@nextui-org/button';
import { TrashIcon } from '@heroicons/react/24/outline';
import { useFormStatus } from 'react-dom';
import { deleteMenuItem } from './DeleteMenuItemAction';

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
      <TrashIcon className="h-7 w-7" />
    </Button>
  );
}

export default function DeleteMenuItemForm({
  menuItemId,
}: {
  menuItemId: string;
}) {
  return (
    <form action={deleteMenuItem}>
      <input type="hidden" name="menuItemId" value={menuItemId} />
      <DeleteButton />
    </form>
  );
}

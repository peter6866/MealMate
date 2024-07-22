'use client';

import { Button } from '@nextui-org/button';
import { useFormStatus } from 'react-dom';
import { orderComplete } from './OrderCompleteAction';

function CompleteButton() {
  const { pending } = useFormStatus();

  return (
    <Button
      type="submit"
      className="text-sm"
      color="primary"
      radius="full"
      size="sm"
      variant="bordered"
      isLoading={pending}
    >
      Complete Order
    </Button>
  );
}

export default function OrderCompleteButton({ orderID }: { orderID: string }) {
  return (
    <form action={orderComplete} className="flex justify-end mt-3">
      <input type="hidden" name="orderID" value={orderID} />
      <CompleteButton />
    </form>
  );
}

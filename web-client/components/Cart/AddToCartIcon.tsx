'use client';

import { addToCart } from './AddToCartAction';
import { PlusCircleIcon } from '@heroicons/react/24/outline';
import { useTheme } from 'next-themes';
import React, { useEffect, useState, useTransition } from 'react';
import { toast } from 'react-toastify';
import { Spinner } from '@nextui-org/spinner';

export default function AddToCartIconForm({
  menuItemId,
  itemName,
}: {
  menuItemId: string;
  itemName: string;
}) {
  const [isPending, startTransition] = useTransition();
  const [shouldShowToast, setShouldShowToast] = useState(false);
  const { theme } = useTheme();

  function handleClick() {
    const formData = new FormData();
    formData.append('menuItemId', menuItemId);
    setShouldShowToast(true);
    startTransition(() => addToCart(formData));
  }

  useEffect(() => {
    if (shouldShowToast && !isPending) {
      toast.success(`${itemName} has been added to cart`, {
        position: 'top-center',
        autoClose: 2000,
        hideProgressBar: true,
        closeOnClick: true,
        pauseOnHover: false,
        draggable: true,
        progress: undefined,
        theme: theme,
      });
      setShouldShowToast(false);
    }
  }, [shouldShowToast, isPending, theme]);

  return (
    <>
      {isPending ? (
        <Spinner size="sm" color="default" />
      ) : (
        <PlusCircleIcon
          className="w-6 h-6 text-mainLight"
          onClick={handleClick}
        />
      )}
    </>
  );
}

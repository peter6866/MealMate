'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';
import { redirect } from 'next/navigation';

interface Item {
  id: string;
  name: string;
  imageUrl: string;
}

export default async function createOrder(items: Item[]) {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  try {
    await axios.post(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/orders`,
      { items },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    revalidatePath('/orders');
    revalidatePath('/cart');
  } catch (error) {
    throw error;
  }
  redirect('/orders');
}

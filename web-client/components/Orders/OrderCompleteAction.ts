'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';

export async function orderComplete(formData: FormData) {
  const orderID = formData.get('orderID') as string;

  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  console.log('orderID', orderID);

  try {
    await axios.put(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/orders/${orderID}`,
      {},
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    revalidatePath('/orders');
  } catch (error) {
    throw error;
  }
}

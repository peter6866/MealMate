'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';

export async function removeFromCart(formData: FormData) {
  const menuItemId = formData.get('menuItemId') as string;

  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  try {
    await axios.delete(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/cart/${menuItemId}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    revalidatePath('/cart');
  } catch (error) {
    throw error;
  }
}

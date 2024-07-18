'use server';

import axios from 'axios';
import { revalidatePath } from 'next/cache';
import { cookies } from 'next/headers';

export async function addToCart(formData: FormData) {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  try {
    await axios.post(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/cart`,
      {
        menuItemID: formData.get('menuItemId'),
      },
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

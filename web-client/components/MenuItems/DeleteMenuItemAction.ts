'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';
import { redirect } from 'next/navigation';

export async function deleteMenuItem(formData: FormData) {
  const menuItemId = formData.get('menuItemId') as string;

  const cookieStore = cookies();
  const token = cookieStore.get('token')!.value;

  try {
    await axios.delete(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/menuItems/${menuItemId}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    revalidatePath('/menuItems');
  } catch (error) {
    throw error;
  }

  redirect('/menuItems');
}

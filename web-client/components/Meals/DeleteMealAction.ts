'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';
import { redirect } from 'next/navigation';

export async function deleteMeal(mealId: string) {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  try {
    await axios.delete(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/meals/${mealId}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    revalidatePath('/meals');
  } catch (error) {
    throw error;
  }

  redirect('/meals');
}

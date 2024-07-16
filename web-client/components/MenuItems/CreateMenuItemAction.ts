'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';
import { redirect } from 'next/navigation';

interface MenuItem {
  name: string;
  alcoholContent?: string;
  spiceLevel?: string;
  referenceLink?: string;
}

export async function createMenuItem(prevState: any, formData: FormData) {
  const menuItem: MenuItem = {
    name: formData.get('name') as string,
    alcoholContent: formData.get('alcoholContent') as string,
    spiceLevel: formData.get('spiceLevel') as string,
    referenceLink: formData.get('referenceLink') as string,
  };

  const jsonData = JSON.stringify(menuItem);

  const imageFile = formData.get('image') as File;

  const form = new FormData();
  form.append('image', imageFile);
  form.append('json', jsonData);
  form.append('categoryId', formData.get('categoryId') as string);

  const cookieStore = cookies();
  const token = cookieStore.get('token')!.value;

  try {
    await axios.post(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/menuItems`,
      form,
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'multipart/form-data',
        },
      }
    );

    revalidatePath('/menuItems');
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        success: false,
        message: error.response?.data.error,
      };
    } else {
      return {
        success: false,
        message: 'An error occurred. Please try again',
      };
    }
  }
  redirect('/menuItems');
}

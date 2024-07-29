'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';
import { redirect } from 'next/navigation';
import { MealItem } from '@/types/meals';

interface MealRequest {
  mealType: string;
  mealDate: string;
  items: MealItem[];
  withPartner: boolean;
}

export async function createMeal(prevState: any, formData: FormData) {
  const formDate = formData.get('date') as string;
  const mealType = formData.get('mealType') as string;

  if (mealType === '') {
    return {
      success: false,
      message: 'Please select a meal type',
    };
  }

  let date;
  if (mealType === 'Breakfast') {
    date = `${formDate}T08:00:00Z`;
  } else if (mealType === 'Lunch') {
    date = `${formDate}T12:00:00Z`;
  } else if (mealType === 'Dinner') {
    date = `${formDate}T18:00:00Z`;
  } else {
    date = `${formDate}T20:00:00Z`;
  }

  const items: MealItem[] = JSON.parse(formData.get('items') as string);
  if (items.length === 0) {
    return {
      success: false,
      message: 'Please select at least one item',
    };
  }

  const meal: MealRequest = {
    mealType,
    mealDate: date,
    items,
    withPartner: formData.get('withPartner') === 'true',
  };

  const jsonData = JSON.stringify(meal);
  const imageFile = formData.get('image') as File;

  if (imageFile.name === 'undefined') {
    return {
      success: false,
      message: 'Please upload an image',
    };
  }

  const form = new FormData();
  form.append('json', jsonData);
  form.append('image', imageFile);

  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  try {
    await axios.post(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/meals`, form, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'multipart/form-data',
      },
    });

    revalidatePath('/meals');
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
  redirect('/meals');
}

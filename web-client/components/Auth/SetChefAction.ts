'use server';

import { cookies } from 'next/headers';
import axios from 'axios';
import { revalidatePath } from 'next/cache';
import { redirect } from 'next/navigation';

export async function setChef(prevState: any, formData: any) {
  const email = formData.get('partner-email');
  const isChef = formData.get('is-chef');

  if (!email.match(/^[a-zA-Z0-9._%+-]+$/)) {
    return {
      success: false,
      message: 'Invalid email address',
    };
  }

  const setData = {
    isChef: isChef === 'true',
    partnerEmail: `${email}@gmail.com`,
  };

  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  try {
    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/auth/setChefAndPartner`,
      setData,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    cookies().set('isChef', response.data.isChef, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      maxAge: 60 * 60 * 24 * 14,
    });
    revalidatePath('/');
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

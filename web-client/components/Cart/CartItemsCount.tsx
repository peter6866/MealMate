import axios from 'axios';
import { cookies } from 'next/headers';

export const revalidate = 0;

export default async function CartItemsCount() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  if (!token) {
    return 0;
  }

  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/cart`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    const cartItems = response.data;

    return cartItems ? cartItems.length : 0;
  } catch (e) {
    return 0;
  }
}

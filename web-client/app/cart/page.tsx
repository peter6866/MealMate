import { Card } from '@nextui-org/card';
import Image from 'next/image';
import axios from 'axios';
import { cookies } from 'next/headers';
import CartList from '@/components/Cart/CartList';

export default async function Cart() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')!.value;

  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/cart`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const cartItems = response.data;

  return (
    <div className="min-h-screen bg-content2 p-4 dark:bg-content1">
      <p className="text-xl font-medium mb-4">Shopping Cart</p>
      {cartItems ? (
        <CartList cartItems={cartItems} />
      ) : (
        <div>
          <h1>You have no items in cart yet</h1>
        </div>
      )}
    </div>
  );
}

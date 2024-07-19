import { Card } from '@nextui-org/card';
import Image from 'next/image';
import axios from 'axios';
import { cookies } from 'next/headers';
import CartList from '@/components/Cart/CartList';
import cartLogo from '@/public/empty-cart.png';
import { Button } from '@nextui-org/button';

export default async function Cart() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

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
    <div className="min-h-screen bg-content2 dark:bg-content1 pb-16">
      <div className="p-4">
        <p className="text-xl font-medium mb-4">Shopping Cart</p>
        {cartItems && cartItems.length != 0 ? (
          <CartList cartItems={cartItems} />
        ) : (
          <Card
            className="p-6 flex flex-col items-center justify-center space-y-4 dark:bg-content2"
            shadow="none"
          >
            <Image
              src={cartLogo}
              alt="Empty Cart"
              width={120}
              height={120}
              priority
              className="mt-2"
            />
            <h2 className="text-lg font-semibold text-center">
              Your cart is empty
            </h2>
            <p className="text-sm text-default-500 text-center">
              Looks like you haven&apos;t added any items to your cart yet.
            </p>
          </Card>
        )}
      </div>
      {cartItems && cartItems.length > 0 && (
        <div className="fixed bottom-16 left-0 right-0 p-4 bg-content1 dark:bg-content1">
          <Button
            color="primary"
            fullWidth
            className="bg-mainLight dark:bg-mainDark text-lg"
          >
            Create Order
          </Button>
        </div>
      )}
    </div>
  );
}

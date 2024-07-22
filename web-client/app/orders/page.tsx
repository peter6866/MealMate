import { cookies } from 'next/headers';
import axios from 'axios';
import OrderList from '@/components/Orders/OrderList';

export default async function Orders() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;
  const isChef = cookieStore.get('isChef')?.value;

  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/orders`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const orders = response.data;

  if (!orders) {
    return (
      <div className="min-h-screen bg-content2 dark:bg-content1 pb-16">
        <div className="p-4">
          <p className="text-xl font-medium mb-4">
            {isChef === 'true' ? 'Incoming Orders' : 'Orders'}
          </p>
          <p className="text-lg text-center">No orders found</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-content2 dark:bg-content1 pb-16">
      <div className="p-4">
        <p className="text-xl font-medium mb-4 text-default-800">
          {isChef === 'true' ? 'Incoming Orders' : 'Orders'}
        </p>
        <OrderList orders={orders} isChef={isChef} />
      </div>
    </div>
  );
}

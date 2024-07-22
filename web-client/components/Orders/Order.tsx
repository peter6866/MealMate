import Image from 'next/image';
import { Card } from '@nextui-org/card';
import moment from 'moment';
import OrderCompleteButton from './OrderCompleteButton';

interface Item {
  id: string;
  name: string;
  imageUrl: string;
}

interface Order {
  id: string;
  createdBy: string;
  sendTo: string;
  status: string;
  items: Item[];
  orderDate: string;
}

export default function Order({
  order,
  isChef,
}: {
  order: Order;
  isChef: string | undefined;
}) {
  const orderFrom = order.createdBy === order.sendTo ? 'You' : 'Your Partner';
  const header =
    isChef === 'true' ? `Order from ${orderFrom}` : 'Order to the Chef';

  const formattedDate = moment(order.orderDate).format('MMMM D, YYYY HH:mm');

  return (
    <Card shadow="none" className="px-4 py-3 dark:bg-content2">
      <div className="flex justify-between items-center text-center mb-2">
        <p className="text-md font-medium text-default-700">{header}</p>
        <p
          className={`text-sm ${
            order.status === 'Started' ? 'text-danger-400' : 'text-mainLight'
          }`}
        >
          {order.status}
        </p>
      </div>
      <div className="grid grid-cols-1 gap-2">
        {order.items.map((item: Item) => (
          <div key={item.id} className="flex items-center justify-between p-0">
            <div className="flex items-center">
              <Image
                src={item.imageUrl}
                alt="Item"
                className="aspect-square rounded-xl"
                width={80}
                height={80}
              />
              <div className="ml-4">
                <h2 className="text-sm font-medium">{item.name}</h2>
                <p className="text-xs text-default-500">Quantity: 1</p>
              </div>
            </div>
          </div>
        ))}
      </div>
      <p className="text-sm text-default-500 mt-3 flex justify-between items-center">
        <span>Order Date:</span>
        <span>{formattedDate}</span>
      </p>
      {isChef === 'true' && order.status === 'Started' && (
        <OrderCompleteButton orderID={order.id} />
      )}
    </Card>
  );
}

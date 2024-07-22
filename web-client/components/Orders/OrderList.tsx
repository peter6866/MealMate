import Order from './Order';

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

export default function OrderList({
  orders,
  isChef,
}: {
  orders: Order[];
  isChef: string | undefined;
}) {
  return (
    <div className="grid grid-cols-1 gap-4">
      {orders.map((order: Order) => (
        <Order key={order.id} order={order} isChef={isChef} />
      ))}
    </div>
  );
}

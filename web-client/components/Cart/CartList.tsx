import Image from 'next/image';
import { Card } from '@nextui-org/card';
import RemoveFromCart from './RemoveFromCartForm';

interface MenuItem {
  id: string;
  name: string;
  imageUrl: string;
}

export default function CartList({ cartItems }: { cartItems: MenuItem[] }) {
  return (
    <Card shadow="none" className="grid grid-cols-1 gap-6 p-4 dark:bg-content2">
      {cartItems.map((item: MenuItem) => (
        <div key={item.id} className="flex items-center justify-between p-0">
          <div className="flex items-center">
            <Image
              src={item.imageUrl}
              alt="Item"
              className="aspect-square rounded-xl"
              width={100}
              height={100}
            />
            <div className="ml-4">
              <h2 className="text-md font-medium">{item.name}</h2>
              <p className="text-sm text-default-500">Quantity: 1</p>
            </div>
          </div>
          <div>
            <RemoveFromCart menuItemId={item.id} />
          </div>
        </div>
      ))}
    </Card>
  );
}

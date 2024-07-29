import CreateMealForm from '@/components/Meals/CreateMealForm';
import axios from 'axios';
import { cookies } from 'next/headers';

export default async function CreateMealsPage() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  const menuItemsResponse = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/menuItems`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const menuItems = menuItemsResponse.data;

  return (
    <div className="fixed inset-0 bg-content2 p-4 flex justify-center">
      <div className="w-full mt-8">
        <CreateMealForm menuItems={menuItems} />
      </div>
    </div>
  );
}

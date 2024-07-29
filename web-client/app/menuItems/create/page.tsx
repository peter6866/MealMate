import CreateMenuItemForm from '@/components/MenuItems/CreateMenuItemForm';
import axios from 'axios';

export default async function CreateMenuItemPage() {
  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/categories`
  );

  const categories = response.data;

  return (
    <div className="fixed inset-0 bg-content2 p-4 flex justify-center">
      <div className="w-full mt-8">
        <CreateMenuItemForm categories={categories} />
      </div>
    </div>
  );
}

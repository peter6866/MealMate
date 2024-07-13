import { Avatar } from '@nextui-org/avatar';
import { Button } from '@nextui-org/button';
import axios from 'axios';
import { cookies } from 'next/headers';
import { ExclamationCircleIcon } from '@heroicons/react/24/outline';
import { QuestionMarkCircleIcon } from '@heroicons/react/24/outline';
import { UserIcon } from '@heroicons/react/24/outline';

import SettingsModal from './settingsModal';

interface User {
  id: string;
  name: string;
  email: string;
  picture: string;
}

export default async function ProfilePage() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')!.value;
  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/auth/getUser`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const user = response.data as User;

  // Placeholder data (replace with actual data from your API)
  const stats = {
    mealsLogged: 42,
    totalOrders: 78,
    dishesEaten: 56,
  };

  return (
    <div className="flex flex-col min-h-[93vh] bg-[#60BEEB] dark:bg-[#115E83]">
      <div className="text-white dark:text-default-700 p-6 mt-4">
        <div className="flex items-center">
          <Avatar
            src={user.picture}
            className="w-20 h-20 text-large border-2 border-background"
          />
          <div className="ml-4">
            <h2 className="text-xl font-semibold">{user.name}</h2>
            <p className="text-sm opacity-85">{user.email}</p>
          </div>
        </div>
      </div>

      <div className="flex justify-around bg-background p-4 opacity-85 mx-4 rounded-xl shadow-lg">
        <div className="text-center">
          <p className="text-2xl font-bold text-default-800">
            {stats.mealsLogged}
          </p>
          <p className="text-xs text-default-500">Meals logged</p>
        </div>
        <div className="text-center">
          <p className="text-2xl font-bold text-default-800">
            {stats.totalOrders}
          </p>
          <p className="text-xs text-default-500">Total orders</p>
        </div>
        <div className="text-center">
          <p className="text-2xl font-bold text-default-800">
            {stats.dishesEaten}
          </p>
          <p className="text-xs text-default-500">Dishes eaten</p>
        </div>
      </div>

      {/* Settings Options */}
      <div className="flex flex-col flex-1 mt-8 bg-background rounded-t-2xl px-1 py-4">
        {[
          {
            label: 'Edit Profile',
            icon: UserIcon,
          },
          {
            label: 'Help & Support',
            icon: QuestionMarkCircleIcon,
          },
          {
            label: 'About',
            icon: ExclamationCircleIcon,
          },
        ].map((item, index) => (
          <Button
            key={index}
            variant="light"
            className="justify-start py-2 mb-2 text-md text-default-800"
            startContent={<item.icon className="w-6 h-6 mr-2" />}
          >
            {item.label}
          </Button>
        ))}
        <SettingsModal />
      </div>
    </div>
  );
}

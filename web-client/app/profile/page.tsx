import { Avatar } from '@nextui-org/avatar';
import { Spacer } from '@nextui-org/spacer';
import axios from 'axios';
import { cookies } from 'next/headers';
import LogoutButton from '@/components/LogoutButton';

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

  return (
    <div className="flex flex-col items-center p-4">
      <div className="w-full">
        <div className="flex flex-col items-center mt-4">
          <Avatar src={user.picture} className="w-20 h-20" />
          <h2 className="text-2xl font-bold mt-2">{user.name}</h2>
        </div>

        <Spacer y={2} />

        <div className="space-y-3">
          <LogoutButton />
        </div>
      </div>
    </div>
  );
}

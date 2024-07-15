'use client';

import { Button } from '@nextui-org/button';
import axios from 'axios';
import { ArrowRightStartOnRectangleIcon } from '@heroicons/react/24/outline';

export default function LogoutButton() {
  const handleLogout = async () => {
    try {
      await axios.get('/api/auth/logout');
      window.location.href = '/';
    } catch (error) {
      console.error('Error during logout:', error);
    }
  };

  return (
    <Button
      className="w-full font-semibold"
      color="danger"
      variant="solid"
      startContent={<ArrowRightStartOnRectangleIcon className="w-5 h-5 mr-2" />}
      onClick={handleLogout}
    >
      Log Out
    </Button>
  );
}

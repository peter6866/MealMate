'use client';

import { Button } from '@nextui-org/button';
import axios from 'axios';

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
      className="w-full justify-start"
      variant="light"
      onClick={handleLogout}
    >
      Log Out
    </Button>
  );
}

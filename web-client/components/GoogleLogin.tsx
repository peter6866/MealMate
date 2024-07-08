'use client';

import { useState } from 'react';
import axios from 'axios';
import { useAuth } from '@/context/AuthContext';

export default function GoogleLogin() {
  const [loading, setLoading] = useState(false);
  const { isLoggedIn, isLoading, logout } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  const handleGoogleLogin = async () => {
    setLoading(true);
    try {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/google_login`
      );
      window.location.href = response.data.url;
    } catch (error) {
      console.error('Error initiating Google login:', error);
      setLoading(false);
    }
  };

  return isLoggedIn ? (
    <div>
      <button onClick={logout} disabled={loading}>
        Logout
      </button>
    </div>
  ) : (
    <button onClick={handleGoogleLogin} disabled={loading}>
      {loading ? 'Loading...' : 'Login with Google'}
    </button>
  );
}

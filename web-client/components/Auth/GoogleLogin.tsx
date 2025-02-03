'use client';

import { useState } from 'react';
import axios from 'axios';
import GoogleButton from 'react-google-button';
import { useTheme } from 'next-themes';

export default function GoogleLogin() {
  const [loading, setLoading] = useState(false);
  const { theme } = useTheme();

  const handleGoogleLogin = async () => {
    setLoading(true);
    try {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/auth/google_login`
      );
      window.location.href = response.data.url;
    } catch (error) {
      console.error('Error initiating Google login:', error);
      setLoading(false);
    }
  };

  return (
    <GoogleButton
      type={theme === 'dark' ? 'dark' : 'light'}
      onClick={handleGoogleLogin}
      disabled={loading}
    />
  );
}

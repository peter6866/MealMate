'use client';

import { useState } from 'react';
import axios from 'axios';

export default function GoogleLogin() {
  const [loading, setLoading] = useState(false);

  const handleGoogleLogin = async () => {
    setLoading(true);
    try {
      const response = await axios.get('http://localhost:8080/google_login');
      window.location.href = response.data.url;
    } catch (error) {
      console.error('Error initiating Google login:', error);
      setLoading(false);
    }
  };

  return (
    <button onClick={handleGoogleLogin} disabled={loading}>
      {loading ? 'Loading...' : 'Login with Google'}
    </button>
  );
}

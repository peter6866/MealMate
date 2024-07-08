'use client';

import React, { createContext, useState, useContext, useEffect } from 'react';
import axios from 'axios';

type AuthContextType = {
  isLoggedIn: boolean;
  isLoading: boolean;
  setIsLoggedIn: (isLoggedIn: boolean) => void;
  logout: () => Promise<void>;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

async function validateToken(): Promise<boolean> {
  try {
    const response = await axios.get('/api/auth/validate-token');
    return response.data.isValid;
  } catch (error) {
    console.error('Error validating token:', error);
    return false;
  }
}

async function getLogout(): Promise<boolean> {
  try {
    const response = await axios.get('/api/auth/logout');
    return response.data.success;
  } catch (error) {
    console.error('Error during logout:', error);
    return false;
  }
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const checkLoginStatus = async () => {
      setIsLoading(true);
      try {
        const isValid = await validateToken();
        setIsLoggedIn(isValid);
      } catch (error) {
        console.error('Error checking login status:', error);
        setIsLoggedIn(false);
      } finally {
        setIsLoading(false);
      }
    };
    checkLoginStatus();
  }, []);

  const logout = async () => {
    setIsLoading(true);
    try {
      // await api.post('/logout');

      // Clear the token cookie
      const success = await getLogout();
      setIsLoggedIn(!success);
    } catch (error) {
      console.error('Error during logout:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <AuthContext.Provider
      value={{ isLoggedIn, isLoading, setIsLoggedIn, logout }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

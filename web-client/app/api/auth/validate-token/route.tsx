import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import axios from 'axios';

async function validateTokenWithBackend(token: string): Promise<boolean> {
  try {
    const response = await axios.post(
      `${process.env.BACKEND_URL}/api/auth/validate-token`,
      null,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    return response.status === 200;
  } catch (error) {
    console.error('Error validating token:', error);
    return false;
  }
}

export async function GET(request: NextRequest) {
  const cookieStore = cookies();
  const token = cookieStore.get('token');

  if (!token) {
    return NextResponse.json({ isValid: false, token: '' });
  }

  if (!token.value) {
    return NextResponse.json({ isValid: false, token: '' });
  }

  // const isValid = await validateTokenWithBackend(token.value);
  const isValid = true; // Temporarily return true to avoid making requests to the backend

  return NextResponse.json({ isValid, token: token.value });
}

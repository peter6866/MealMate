import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import axios from 'axios';

// TODO: Invalidate the token on the backend

// clears the token cookie
export async function GET(request: NextRequest) {
  const cookieStore = cookies();

  cookieStore.delete('token');
  cookieStore.delete('isChef');

  return NextResponse.json({ message: 'Logged out' });
}

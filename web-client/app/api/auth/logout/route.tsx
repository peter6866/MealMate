import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import axios from 'axios';

// clears the token cookie
export async function GET(request: NextRequest) {
  const cookieStore = cookies();
  cookieStore.set('token', '', {
    httpOnly: true,
    secure: false,
    sameSite: 'lax',
  });

  return NextResponse.json({ success: true });
}

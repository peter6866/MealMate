import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import axios from 'axios';

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const code = searchParams.get('code');
  const state = searchParams.get('state');

  if (!code || !state) {
    return NextResponse.redirect(new URL('/', request.url));
  }

  try {
    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/auth/loginOrRegister`,
      {
        code,
        state,
      }
    );

    const { token } = response.data;

    cookies().set('token', token, {
      httpOnly: true,
      secure: false,
      sameSite: 'lax',
    });

    const redirectUrl = new URL('/', request.url);

    return NextResponse.redirect(redirectUrl);
  } catch (error) {
    console.error(error);
    return NextResponse.redirect(new URL('/', request.url));
  }
}
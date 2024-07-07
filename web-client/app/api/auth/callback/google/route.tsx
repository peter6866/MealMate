import { NextRequest, NextResponse } from 'next/server';
import axios from 'axios';

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const code = searchParams.get('code');
  const state = searchParams.get('state');

  if (!code || !state) {
    return NextResponse.redirect('/');
  }

  try {
    const response = await axios.post(
      'http://localhost:8080/api/auth/loginOrRegister',
      {
        code,
        state,
      }
    );

    const { token } = response.data;
    console.log(token);

    const redirectUrl = new URL('/', request.url);

    return NextResponse.redirect(redirectUrl);
  } catch (error) {
    console.error(error);
    return NextResponse.redirect('/');
  }
}

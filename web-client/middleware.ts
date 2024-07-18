import { NextRequest, NextResponse } from 'next/server';
import axios from 'axios';

export async function middleware(req: NextRequest) {
  const token = req.cookies.get('token');

  async function verifyUser() {
    if (!token || token.value === '') {
      return { isLoggedIn: false, data: null };
    }

    try {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/auth/getUser`,
        {
          headers: {
            Authorization: `Bearer ${token.value}`,
          },
        }
      );

      return { isLoggedIn: response.status === 200, data: response.data };
    } catch (error) {
      console.error(error);
      return { isLoggedIn: false, data: null };
    }
  }

  const { isLoggedIn, data } = await verifyUser();

  // check if the user is on the landing page
  if (req.nextUrl.pathname === '/') {
    if (isLoggedIn) {
      return NextResponse.redirect(new URL('/menuItems', req.nextUrl));
    }
    return NextResponse.next();
  }

  if (!isLoggedIn) {
    return NextResponse.redirect(new URL('/login', req.nextUrl));
  }

  if (
    req.nextUrl.pathname !== '/profile' &&
    req.nextUrl.pathname !== '/login/setChefAndPartner' &&
    !data.partnerEmail
  ) {
    return NextResponse.redirect(
      new URL('/login/setChefAndPartner', req.nextUrl)
    );
  }

  if (req.nextUrl.pathname === '/menuItems/create' && !data.isChef) {
    return NextResponse.redirect(new URL('/menuItems', req.nextUrl));
  }

  if (
    req.nextUrl.pathname === '/login/setChefAndPartner' &&
    data.partnerEmail
  ) {
    return NextResponse.redirect(new URL('/menuItems', req.nextUrl));
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    '/profile/:path*',
    '/menuItems/:path*',
    '/',
    '/cart/:path*',
    '/login/setChefAndPartner',
  ],
};

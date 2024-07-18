import '@/styles/globals.css';
import { Metadata, Viewport } from 'next';

import { siteConfig } from '@/config/site';
import { Navbar } from '@/components/navbar';
import { Providers } from './providers';
import { Suspense } from 'react';
import CartItemsCount from '@/components/Cart/CartItemsCount';

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s - ${siteConfig.name}`,
  },
  description: siteConfig.description,
  icons: {
    icon: '/favicon.ico',
  },
};

export const viewport: Viewport = {
  themeColor: [
    { media: '(prefers-color-scheme: light)', color: 'white' },
    { media: '(prefers-color-scheme: dark)', color: 'black' },
  ],
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html suppressHydrationWarning lang="en">
      <head />
      <body>
        <Providers>
          <main className="container mx-auto max-w-7xl pb-16">{children}</main>
          <Suspense fallback={<Navbar cartItemsCount={0} />}>
            <CartItemsCountWrapper />
          </Suspense>
        </Providers>
      </body>
    </html>
  );
}

async function CartItemsCountWrapper() {
  const count = await CartItemsCount();
  return <Navbar cartItemsCount={count} />;
}

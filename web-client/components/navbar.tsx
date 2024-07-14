'use client';

import { MagnifyingGlassIcon } from '@heroicons/react/24/outline';
import { UserCircleIcon } from '@heroicons/react/24/outline';
import { CalendarDaysIcon } from '@heroicons/react/24/outline';
import { ShoppingCartIcon } from '@heroicons/react/24/outline';
import { ClipboardDocumentListIcon } from '@heroicons/react/24/outline';

import { MagnifyingGlassIcon as SearchSolid } from '@heroicons/react/24/solid';
import { UserCircleIcon as UserSolid } from '@heroicons/react/24/solid';
import { CalendarDaysIcon as CalendarSolid } from '@heroicons/react/24/solid';
import { ShoppingCartIcon as ShoppingCartSolid } from '@heroicons/react/24/solid';
import { ClipboardDocumentListIcon as ClipboardSolid } from '@heroicons/react/24/solid';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

interface NavItemProps {
  href: string;
  icon: any;
  solidIcon: any;
  label: string;
}

const NavItem: React.FC<NavItemProps> = ({
  href,
  icon: Icon,
  solidIcon: SolidIcon,
  label,
}) => {
  const pathname = usePathname();
  const isActive = pathname === href;

  return (
    <Link
      href={href}
      className="inline-flex flex-col items-center justify-center px-5"
    >
      {isActive ? (
        <SolidIcon className="h-6 w-6 text-[#60BEEB]" />
      ) : (
        <Icon className="h-6 w-6 text-default-600" />
      )}
      <span
        className={`text-sm ${isActive ? 'text-[#60BEEB]' : 'text-default-600'}`}
      >
        {label}
      </span>
    </Link>
  );
};

export const Navbar = () => {
  return (
    <div className="fixed bottom-0 left-0 z-50 w-full h-16 border-t border-gray-200 bg-white dark:bg-default-100 dark:border-0">
      <div className="grid h-full max-w-lg grid-cols-5 mx-auto font-medium">
        <NavItem
          href="/menuItems"
          icon={MagnifyingGlassIcon}
          solidIcon={SearchSolid}
          label="Discover"
        />
        <NavItem
          href="/orders"
          icon={ClipboardDocumentListIcon}
          solidIcon={ClipboardSolid}
          label="Orders"
        />
        <NavItem
          href="/meals"
          icon={CalendarDaysIcon}
          solidIcon={CalendarSolid}
          label="Meals"
        />
        <NavItem
          href="/cart"
          icon={ShoppingCartIcon}
          solidIcon={ShoppingCartSolid}
          label="Cart"
        />
        <NavItem
          href="/profile"
          icon={UserCircleIcon}
          solidIcon={UserSolid}
          label="Profile"
        />
      </div>
    </div>
  );
};

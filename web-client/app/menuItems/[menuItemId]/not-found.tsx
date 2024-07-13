import { Button } from '@nextui-org/button';
import Link from 'next/link';

export default function NotFound() {
  return (
    <div className="fixed inset-0 flex flex-col items-center justify-center min-h-screen bg-content1 px-4">
      <div className="w-full max-w-md p-8 space-y-4">
        <div className="w-24 h-24 mx-auto">
          <svg
            className="w-full h-full text-[#60BEEB]"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fillRule="evenodd"
              d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z"
              clipRule="evenodd"
            />
          </svg>
        </div>
        <p className="text-2xl font-bold text-center text-gray-800">
          Oops! Menu Item Not Found
        </p>
        <p className="text-center text-gray-600">
          We couldn&apos;t find the menu item you&apos;re looking for. It might
          have been removed or doesn&apos;t exist.
        </p>
        <Link href="/" passHref>
          <Button as="a" color="primary" className="w-full bg-[#60BEEB] mt-6">
            Back to All Items
          </Button>
        </Link>
      </div>
    </div>
  );
}

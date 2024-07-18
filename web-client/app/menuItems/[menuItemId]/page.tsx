import axios, { AxiosError } from 'axios';
import { cookies } from 'next/headers';
import Image from 'next/image';
import { Button } from '@nextui-org/button';
import { Chip } from '@nextui-org/chip';
import { ArrowLeftCircleIcon, TrashIcon } from '@heroicons/react/24/outline';
import Link from 'next/link';
import { notFound } from 'next/navigation';
import DeleteMenuItemForm from '@/components/MenuItems/DeleteMenuItemForm';
import ItemAddToCartForm from '@/components/Cart/ItemAddToCartForm';

interface Dish {
  id: string;
  name: string;
  categoryName: string;
  imageUrl: string;
  referenceLink?: string;
  spiceLevel?: string;
  alcoholContent?: string;
}

interface User {
  isChef: boolean;
}

export async function generateMetadata({ params }: { params: any }) {
  try {
    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/menuItems/${params.menuItemId}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    const { name } = response.data;
    return {
      title: name,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      const axiosError = error as AxiosError;
      if (axiosError.response?.status === 404) {
        notFound();
      }
    }
    throw error;
  }
}

export default async function MenuItem({ params }: { params: any }) {
  try {
    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;

    const [userResponse, menuItemResponse] = await Promise.all([
      await axios.get(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/auth/getUser`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      ),
      await axios.get(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/menuItems/${params.menuItemId}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      ),
    ]);

    const menuItem: Dish = menuItemResponse.data;
    const user: User = userResponse.data;

    const spiceMap: { [key: string]: string }[] = [
      { Spicy: 'Spicy' },
      { Mild: 'Mild' },
      { NotSpicy: 'Not Spicy' },
      { VerySpicy: 'Very Spicy' },
    ];

    const alcoholMap: { [key: string]: string }[] = [
      { NonAlcoholic: 'Non-Alcoholic' },
      { Alcoholic: 'Alcoholic' },
    ];

    const {
      name,
      categoryName,
      imageUrl,
      spiceLevel,
      alcoholContent,
      referenceLink,
    } = menuItem;

    return (
      <div className="fixed inset-0 flex flex-col bg-content1 pb-16">
        <div className="relative w-full aspect-square">
          <Image
            src={imageUrl}
            alt={name}
            fill
            className="object-cover"
            priority
          />
          <Link href="/" className="absolute top-4 left-4 z-10">
            <div className="relative p-1 rounded-full bg-black bg-opacity-30 backdrop-blur-sm">
              <ArrowLeftCircleIcon className="h-10 w-10 text-white drop-shadow-lg" />
            </div>
          </Link>
        </div>
        <div className="flex-1 p-6 flex flex-col justify-between">
          <div className="flex justify-between">
            <span className="text-3xl font-bold text-default-800 mb-2">
              {name}
            </span>
            {user.isChef && (
              <DeleteMenuItemForm menuItemId={params.menuItemId} />
            )}
          </div>
          <div className="flex gap-4 mt-3">
            <Chip variant="flat">{categoryName}</Chip>
            {spiceLevel && (
              <Chip color="danger" variant="bordered">
                {spiceMap.find((map) => map[spiceLevel])![spiceLevel]}
              </Chip>
            )}
            {alcoholContent && (
              <Chip color="warning">
                {alcoholMap.find((map) => map[alcoholContent])![alcoholContent]}
              </Chip>
            )}
          </div>
          {referenceLink && (
            <a
              href={referenceLink}
              target="_blank"
              rel="noopener noreferrer"
              className="text-primary-500 mt-4"
            >
              Click here to view the cooking instructions
            </a>
          )}
          <div className="mt-auto">
            <ItemAddToCartForm menuItemId={params.menuItemId} />
          </div>
        </div>
      </div>
    );
  } catch (error) {
    if (axios.isAxiosError(error)) {
      const axiosError = error as AxiosError;
      if (axiosError.response?.status === 404) {
        notFound();
      }
    }
    throw error;
  }
}

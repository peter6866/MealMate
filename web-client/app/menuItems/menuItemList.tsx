import React from 'react';
import MenuItem from './menuItem';
import axios from 'axios';
import { cookies } from 'next/headers';

interface Dish {
  id: string;
  name: string;
  categoryName: string;
  imageUrl: string;
  spiceLevel: string;
  alcoholContent: string;
}

interface searchParamsProps {
  category?: string;
  spiceLevel?: string;
  alcoholContent?: string;
}

async function fetchMenuItems() {
  const cookieStore = cookies();
  const token = cookieStore.get('token')?.value;

  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/menuItems`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  return response.data;
}

export default async function MenuItemList({
  filter,
}: {
  filter: searchParamsProps;
}) {
  const dishes = await fetchMenuItems();

  if (!dishes.length) {
    return null;
  }

  const { category, spiceLevel, alcoholContent } = filter;

  const filteredDishes = dishes.filter((dish: Dish) => {
    const categoryMatch = !category || dish.categoryName === category;
    const spiceLevelMatch = !spiceLevel || dish.spiceLevel === spiceLevel;
    const alcoholContentMatch =
      !alcoholContent || dish.alcoholContent === alcoholContent;

    return categoryMatch && spiceLevelMatch && alcoholContentMatch;
  });

  return (
    <div className="grid grid-cols-2 gap-4">
      {filteredDishes.map((dish: Dish) => (
        <MenuItem key={dish.id} dish={dish} />
      ))}
    </div>
  );
}

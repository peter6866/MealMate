import React from 'react';
import MenuItem from './menuItem';
import axios from 'axios';
import { cookies } from 'next/headers';

interface Dish {
  id: string;
  name: string;
  categoryName: string;
  imageUrl: string;
}

async function fetchMenuItems() {
  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/menuItems`
  );

  return response.data;
}

export default async function MenuItemList() {
  const dishes = await fetchMenuItems();

  if (!dishes.length) {
    return null;
  }

  return (
    <div className="grid grid-cols-2 gap-4">
      {dishes.map((dish: Dish) => (
        <MenuItem key={dish.id} dish={dish} />
      ))}
    </div>
  );
}

import axios from 'axios';

import FilterButton from './FilterButton';

interface Category {
  id: string;
  name: string;
}

interface SpiceLevel {
  id: string;
  value: string;
  label: string;
}

const spiceLevels: SpiceLevel[] = [
  {
    id: '1',
    value: 'NotSpicy',
    label: 'Not Spicy',
  },
  {
    id: '2',
    value: 'Mild',
    label: 'Mild',
  },
  {
    id: '3',
    value: 'Spicy',
    label: 'Spicy',
  },
  {
    id: '4',
    value: 'VerySpicy',
    label: 'Very Spicy',
  },
];

export default async function Filter() {
  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/categories`
  );

  const categories = response.data as Category[];

  return <FilterButton categories={categories} spiceLevels={spiceLevels} />;
}

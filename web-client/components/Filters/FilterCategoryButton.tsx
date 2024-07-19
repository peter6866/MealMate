'use client';

import { Button } from '@nextui-org/button';

interface Category {
  id: string;
  name: string;
}

export default function FilterCategoryButton({
  category,
  handleFilter,
  selectedFilter,
}: {
  category: Category;
  handleFilter: (filter: string) => void;
  selectedFilter: string | null;
}) {
  return (
    <Button
      size="sm"
      className={`text-sm rounded-2xl ${
        category.name === selectedFilter
          ? 'bg-mainLight dark:bg-mainDark text-white dark:text-default-800 font-medium'
          : 'bg-content3 text-default-800 '
      }`}
      onClick={() => handleFilter(category.name)}
    >
      {category.name}
    </Button>
  );
}

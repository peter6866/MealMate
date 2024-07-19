'use client';

import { Button } from '@nextui-org/button';

interface SpiceLevel {
  id: string;
  value: string;
  label: string;
}

export default function FilterSpiceLevelButton({
  spiceLevel,
  handleFilter,
  selectedFilter,
}: {
  spiceLevel: SpiceLevel;
  handleFilter: (filter: string) => void;
  selectedFilter: string | null;
}) {
  return (
    <Button
      size="sm"
      className={`text-sm rounded-2xl ${
        spiceLevel.value === selectedFilter
          ? 'bg-red-500 dark:bg-red-700 text-white dark:text-default-800 font-medium'
          : 'bg-content3 text-default-800 '
      }`}
      onClick={() => handleFilter(spiceLevel.value)}
    >
      {spiceLevel.label}
    </Button>
  );
}

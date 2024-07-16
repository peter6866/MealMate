'use client';

import { Button } from '@nextui-org/button';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';

interface SpiceLevel {
  id: string;
  value: string;
  label: string;
}

export default function FilterSpiceLevelButton({
  spiceLevel,
}: {
  spiceLevel: SpiceLevel;
}) {
  const searchParams = useSearchParams();
  const router = useRouter();
  const pathname = usePathname();

  const activeFilter = searchParams.get('spiceLevel');

  function handleCategoryFilter(spiceLevel: string) {
    const params = new URLSearchParams(searchParams);
    if (activeFilter === spiceLevel) {
      params.delete('spiceLevel');
      router.replace(`${pathname}?${params.toString()}`, { scroll: false });
    } else {
      params.set('spiceLevel', spiceLevel);
      router.replace(`${pathname}?${params.toString()}`, { scroll: false });
    }
  }

  return (
    <Button
      size="sm"
      className={`text-sm mr-2 ${spiceLevel.value === activeFilter ? 'bg-red-500 dark:bg-red-700 text-white dark:text-default-800 font-medium' : 'bg-content3 text-default-800 '}`}
      onClick={() => handleCategoryFilter(spiceLevel.value)}
    >
      {spiceLevel.label}
    </Button>
  );
}

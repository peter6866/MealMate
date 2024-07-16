'use client';

import { Button } from '@nextui-org/button';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';

interface Category {
  id: string;
  name: string;
}

export default function FilterCategoryButton({
  category,
}: {
  category: Category;
}) {
  const searchParams = useSearchParams();
  const router = useRouter();
  const pathname = usePathname();

  const activeFilter = searchParams.get('category');

  function handleCategoryFilter(cat: string) {
    const params = new URLSearchParams(searchParams);
    if (activeFilter === cat) {
      params.delete('category');
      router.replace(`${pathname}?${params.toString()}`, { scroll: false });
    } else {
      params.set('category', cat);
      router.replace(`${pathname}?${params.toString()}`, { scroll: false });
    }
  }

  return (
    <Button
      size="sm"
      className={`text-sm mr-2 ${category.name === activeFilter ? 'bg-mainLight dark:bg-mainDark text-white dark:text-default-800 font-medium' : 'bg-content3 text-default-800 '}`}
      onClick={() => handleCategoryFilter(category.name)}
    >
      {category.name}
    </Button>
  );
}

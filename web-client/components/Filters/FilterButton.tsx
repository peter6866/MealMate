'use client';

import {
  Modal,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
} from '@nextui-org/modal';
import { Button } from '@nextui-org/button';
import { useDisclosure } from '@nextui-org/react';
import { FunnelIcon } from '@heroicons/react/24/outline';
import FilterCategoryButton from './FilterCategoryButton';
import FilterSpiceLevelButton from './FilterSpiceLevelButton';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import { useState } from 'react';

interface Category {
  id: string;
  name: string;
}

interface SpiceLevel {
  id: string;
  value: string;
  label: string;
}

export default function FilterButton({
  categories,
  spiceLevels,
}: {
  categories: Category[];
  spiceLevels: SpiceLevel[];
}) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const searchParams = useSearchParams();
  const router = useRouter();
  const pathname = usePathname();

  const activeCategoryFilter = searchParams.get('category');
  const activeSpiceFilter = searchParams.get('spiceLevel');

  const [selectedCategory, setSelectedCategory] = useState<string | null>(
    activeCategoryFilter
  );
  const [selectedSpiceLevel, setSelectedSpiceLevel] = useState<string | null>(
    activeSpiceFilter
  );

  function handleCategoryFilter(category: string) {
    setSelectedCategory((prev) => (prev === category ? null : category));
  }

  function handleSpiceLevelFilter(spiceLevel: string) {
    setSelectedSpiceLevel((prev) => (prev === spiceLevel ? null : spiceLevel));
  }

  function handleApplyFilter(onClose: () => void) {
    const params = new URLSearchParams(searchParams);
    if (!selectedSpiceLevel) {
      params.delete('spiceLevel');
    } else {
      params.set('spiceLevel', selectedSpiceLevel);
    }

    if (!selectedCategory) {
      params.delete('category');
    } else {
      params.set('category', selectedCategory);
    }

    router.replace(`${pathname}?${params.toString()}`, { scroll: false });

    onClose();
  }

  return (
    <>
      <Button onClick={onOpen} isIconOnly variant="light">
        <FunnelIcon className="h-[30px] w-[30px] text-default-700" />
      </Button>
      <Modal
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        placement="center"
        hideCloseButton
        className="bg-content1"
      >
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className="text-xl">Filters</ModalHeader>
              <ModalBody className="pt-0">
                <p className="text-md">Categories</p>
                <div className="mb-2 flex flex-wrap gap-2">
                  {categories.map((category: Category) => (
                    <FilterCategoryButton
                      key={category.id}
                      category={category}
                      handleFilter={handleCategoryFilter}
                      selectedFilter={selectedCategory}
                    />
                  ))}
                </div>
                <p className="text-md">Spice Level</p>
                <div className="flex flex-wrap gap-2">
                  {spiceLevels.map((spiceLevel: SpiceLevel) => (
                    <FilterSpiceLevelButton
                      key={spiceLevel.id}
                      spiceLevel={spiceLevel}
                      handleFilter={handleSpiceLevelFilter}
                      selectedFilter={selectedSpiceLevel}
                    />
                  ))}
                </div>
              </ModalBody>
              <ModalFooter>
                <Button
                  color="primary"
                  onPress={() => handleApplyFilter(onClose)}
                  className="bg-mainLight dark:bg-mainDark rounded-2xl"
                >
                  Apply
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}

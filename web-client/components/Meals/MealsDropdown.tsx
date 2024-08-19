'use client';

import {
  Dropdown,
  DropdownTrigger,
  DropdownMenu,
  DropdownItem,
} from '@nextui-org/dropdown';
import { EllipsisHorizontalIcon } from '@heroicons/react/24/outline';
import { deleteMeal } from './DeleteMealAction';

export default function MealsDropdown({ mealId }: { mealId: string }) {
  return (
    <Dropdown>
      <DropdownTrigger>
        <EllipsisHorizontalIcon className="h-6 w-6 text-default-700" />
      </DropdownTrigger>
      <DropdownMenu
        aria-label="Actions"
        disabledKeys={['edit']}
        onAction={(key) => {
          if (key === 'delete') {
            deleteMeal(mealId);
          }
        }}
      >
        <DropdownItem key="edit">Edit meal</DropdownItem>
        <DropdownItem key="delete" className="text-danger" color="danger">
          Delete meal
        </DropdownItem>
      </DropdownMenu>
    </Dropdown>
  );
}

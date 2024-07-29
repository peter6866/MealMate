'use client';

import React, { useState } from 'react';
import { useFormStatus } from 'react-dom';
import { useFormState } from 'react-dom';
import { DatePicker } from '@nextui-org/date-picker';
import { Select, SelectItem } from '@nextui-org/select';
import { Button } from '@nextui-org/button';
import { Card, CardHeader, CardBody } from '@nextui-org/card';
import { Switch } from '@nextui-org/switch';
import { DocumentArrowUpIcon } from '@heroicons/react/24/outline';
import { CheckCircleIcon } from '@heroicons/react/24/outline';
import { createMeal } from './CreateMealAction';
import { useRouter } from 'next/navigation';
import { MealItem } from '@/types/meals';
import { MenuItem } from '@/types/menuItems';
import { Avatar } from '@nextui-org/avatar';
import { parseDate, getLocalTimeZone, today } from '@internationalized/date';

const initialFormState = {
  success: false,
  message: '',
};

const mealTypes = ['Breakfast', 'Lunch', 'Dinner', 'Snacks'];

export default function CreateMealForm({
  menuItems,
}: {
  menuItems: MenuItem[];
}) {
  const [state, formAction] = useFormState(createMeal, initialFormState);
  const [initialState, setInitialState] = useState({
    mealType: '',
    date: parseDate(new Date().toISOString().split('T')[0]),
    withPartner: false,
  });

  const [items, setItems] = useState<MealItem[]>([]);

  return (
    <Card className="w-full shadow-md">
      <CardHeader className="flex flex-col items-start px-6 pt-6">
        <p className="text-2xl font-bold text-mainLight dark:text-default-700">
          Log your meal
        </p>
      </CardHeader>
      <CardBody className="px-6">
        <form action={formAction}>
          <FileUpload />

          <Select
            label="Dishes"
            name="menuItems"
            required
            fullWidth
            size="lg"
            variant="underlined"
            selectionMode="multiple"
            selectedKeys={items.map((item) => item.id)}
            onChange={(e) => {
              const selectedIds = new Set(e.target.value.split(','));
              const selectedItems = menuItems
                .filter((item) => selectedIds.has(item.id))
                .map(({ id, name, imageUrl }) => ({ id, name, imageUrl }));
              setItems(selectedItems);
            }}
          >
            {menuItems.map((item) => (
              <SelectItem
                key={item.id}
                value={item.id}
                startContent={
                  <Avatar
                    alt={item.name}
                    src={item.imageUrl}
                    className="w-7 h-7"
                  />
                }
              >
                {item.name}
              </SelectItem>
            ))}
          </Select>

          <input type="hidden" name="items" value={JSON.stringify(items)} />

          <Select
            label="Meal Type"
            name="mealType"
            variant="underlined"
            fullWidth
            size="lg"
            className="mb-2"
            selectedKeys={[initialState.mealType]}
            onChange={(e) =>
              setInitialState({
                ...initialState,
                mealType: e.target.value,
              })
            }
          >
            {mealTypes.map((type) => (
              <SelectItem key={type} value={type}>
                {type}
              </SelectItem>
            ))}
          </Select>

          <DatePicker
            name="date"
            aria-label="Select date"
            variant="underlined"
            value={initialState.date}
            onChange={(date) => {
              setInitialState({
                ...initialState,
                date: date,
              });
            }}
            maxValue={today(getLocalTimeZone())}
            className="w-full mt-4"
          />

          <div className="flex justify-between mt-8">
            <p className="text-md text-default-700">
              Having this meal with your partner?
            </p>
            <Switch
              isSelected={initialState.withPartner}
              onChange={(e) =>
                setInitialState({
                  ...initialState,
                  withPartner: e.target.checked,
                })
              }
            />
          </div>

          <input
            type="hidden"
            name="withPartner"
            value={initialState.withPartner ? 'true' : 'false'}
          />

          {state.message && (
            <p className="text-danger-500 text-sm text-center mt-4">
              {state.message}
            </p>
          )}
          <SubmitButton />
        </form>
        <GoBackButton />
      </CardBody>
    </Card>
  );
}

function FileUpload() {
  const [fileName, setFileName] = useState<string | null>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    setFileName(file?.name || null);
  };
  return (
    <div className="w-full">
      <input
        type="file"
        name="image"
        accept="image/*"
        className="hidden"
        id="file-upload"
        onChange={handleFileChange}
      />
      <label htmlFor="file-upload" className="w-full">
        <Button
          variant={fileName ? 'solid' : 'bordered'}
          as="span"
          color={fileName ? 'success' : 'primary'}
          size="md"
          className="border-mainLight text-mainLight dark:border-default-700 dark:text-default-700"
          startContent={
            fileName ? (
              <CheckCircleIcon className="w-6 h-6 text-white" />
            ) : (
              <DocumentArrowUpIcon className="w-5 h-5" />
            )
          }
        >
          {fileName ? '' : 'Upload a photo for your meal'}
        </Button>
      </label>
    </div>
  );
}

function SubmitButton() {
  const { pending } = useFormStatus();
  return (
    <Button
      type="submit"
      color="primary"
      className="bg-mainLight dark:bg-mainDark mt-6"
      size="lg"
      fullWidth
      isLoading={pending}
    >
      Add
    </Button>
  );
}

function GoBackButton() {
  const router = useRouter();
  return (
    <Button
      color="danger"
      size="lg"
      variant="light"
      className="w-full mt-2"
      onClick={() => router.push('/meals')}
    >
      Go Back
    </Button>
  );
}

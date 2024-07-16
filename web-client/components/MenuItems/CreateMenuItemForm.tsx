'use client';

import React, { useState, useEffect } from 'react';
import { useFormStatus } from 'react-dom';
import { useFormState } from 'react-dom';
import { Input } from '@nextui-org/input';
import { Select, SelectItem } from '@nextui-org/select';
import { Button } from '@nextui-org/button';
import { Card, CardHeader, CardBody } from '@nextui-org/card';
import { DocumentArrowUpIcon } from '@heroicons/react/24/outline';
import { CheckCircleIcon } from '@heroicons/react/24/outline';
import { createMenuItem } from './CreateMenuItemAction';
import { useRouter } from 'next/navigation';

interface Category {
  id: string;
  name: string;
}

const initialFormState = {
  success: false,
  message: '',
};

const spiceLevels = [
  { label: 'Not Spicy', value: 'NotSpicy' },
  { label: 'Mild', value: 'Mild' },
  { label: 'Spicy', value: 'Spicy' },
  { label: 'Very Spicy', value: 'VerySpicy' },
];

const alcoholContents = [
  { label: 'Non-Alcoholic', value: 'NonAlcoholic' },
  { label: 'Alcoholic', value: 'Alcoholic' },
];

export default function CreateMenuItemForm({
  categories,
}: {
  categories: Category[];
}) {
  const [state, formAction] = useFormState(createMenuItem, initialFormState);
  const [selectedCategory, setSelectedCategory] = useState('');
  const [initialState, setInitialState] = useState({
    name: '',
    categoryId: '',
    alcoholContent: '',
    spiceLevel: '',
    referenceLink: '',
  });

  useEffect(() => {
    function handleCategoryChange(catId: string) {
      const cat = categories.find((category) => category.id === catId);

      setSelectedCategory(cat?.name || '');

      if (cat?.name === 'Drinks') {
        setInitialState({ ...initialState, spiceLevel: '' });
      } else {
        setInitialState({ ...initialState, alcoholContent: '' });
      }
    }
    handleCategoryChange(initialState.categoryId);
  }, [initialState.categoryId]);

  return (
    <Card className="w-full shadow-md">
      <CardHeader className="flex flex-col items-start px-6 pt-8">
        <p className="text-2xl font-bold text-mainLight dark:text-default-700">
          Add your Dish or Drink
        </p>
      </CardHeader>
      <CardBody className="px-6">
        <form action={formAction}>
          <FileUpload />

          <Input
            label="Name"
            name="name"
            required
            fullWidth
            size="lg"
            className="mb-4"
            variant="underlined"
            value={initialState.name}
            onChange={(e) =>
              setInitialState({ ...initialState, name: e.target.value })
            }
          />

          <Select
            label="Category"
            name="categoryId"
            required
            fullWidth
            size="lg"
            variant="underlined"
            selectedKeys={[initialState.categoryId]}
            onChange={(e) => {
              setInitialState({ ...initialState, categoryId: e.target.value });
            }}
            className="mb-4"
          >
            {categories.map((cat) => (
              <SelectItem key={cat.id} value={cat.id}>
                {cat.name}
              </SelectItem>
            ))}
          </Select>

          {selectedCategory === 'Drinks' ? (
            <Select
              label="Alcohol Content"
              name="alcoholContent"
              variant="underlined"
              fullWidth
              size="lg"
              className="mb-4"
              selectedKeys={[initialState.alcoholContent]}
              onChange={(e) =>
                setInitialState({
                  ...initialState,
                  alcoholContent: e.target.value,
                })
              }
            >
              {alcoholContents.map((content) => (
                <SelectItem key={content.value} value={content.value}>
                  {content.label}
                </SelectItem>
              ))}
            </Select>
          ) : (
            <Select
              label="Spice Level"
              name="spiceLevel"
              fullWidth
              size="lg"
              variant="underlined"
              className="mb-4"
              selectedKeys={[initialState.spiceLevel]}
              onChange={(e) =>
                setInitialState({ ...initialState, spiceLevel: e.target.value })
              }
            >
              {spiceLevels.map((level) => (
                <SelectItem key={level.value} value={level.value}>
                  {level.label}
                </SelectItem>
              ))}
            </Select>
          )}

          <Input
            label="Reference Link (Optional)"
            name="referenceLink"
            fullWidth
            size="lg"
            variant="underlined"
            className="mb-6"
            value={initialState.referenceLink}
            onChange={(e) =>
              setInitialState({
                ...initialState,
                referenceLink: e.target.value,
              })
            }
          />
          {state.message && (
            <p className="text-danger-500 text-sm text-center mb-4">
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
          className="max-w-40 border-mainLight text-mainLight dark:border-default-700 dark:text-default-700"
          startContent={
            fileName ? (
              <CheckCircleIcon className="w-6 h-6 text-white" />
            ) : (
              <DocumentArrowUpIcon className="w-5 h-5" />
            )
          }
        >
          {fileName ? '' : 'Upload Image'}
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
      className="bg-mainLight dark:bg-mainDark"
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
      onClick={() => router.push('/menuItems')}
    >
      Go Back
    </Button>
  );
}

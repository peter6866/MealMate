import axios from 'axios';
import FilterCategoryButton from './FilterCategoryButton';
import FilterSpiceLevelButton from './FilterSpiceLevelButton';

interface Category {
  id: number;
  name: string;
}

interface SpiceLevel {
  id: number;
  value: string;
  label: string;
}

const spiceLevels = [
  {
    id: 1,
    value: 'NotSpicy',
    label: 'Not Spicy',
  },
  {
    id: 2,
    value: 'Mild',
    label: 'Mild',
  },
  {
    id: 3,
    value: 'Spicy',
    label: 'Spicy',
  },
  {
    id: 4,
    value: 'VerySpicy',
    label: 'Very Spicy',
  },
];

// TODO: Change filter to a modal
export default async function Filter() {
  const response = await axios.get(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/categories`
  );

  const categories = response.data;

  return (
    <div className="mb-4">
      <div className="mb-2 overflow-x-auto whitespace-nowrap">
        {categories.map((category: Category) => (
          <FilterCategoryButton key={category.id} category={category} />
        ))}
      </div>
      <div className="overflow-x-auto whitespace-nowrap">
        {spiceLevels.map((spiceLevel: SpiceLevel) => (
          <FilterSpiceLevelButton key={spiceLevel.id} spiceLevel={spiceLevel} />
        ))}
      </div>
    </div>
  );
}

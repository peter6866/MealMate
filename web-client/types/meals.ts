export interface MealItem {
  id: string;
  name: string;
  imageUrl: string;
}

export interface Meal {
  id: string;
  mealType: string;
  mealDate: string;
  photoURL: string;
  items: MealItem[];
  withPartner: boolean;
}

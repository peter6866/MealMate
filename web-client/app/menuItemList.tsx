import MenuItem from './menuItem';

const dishes = [
  {
    name: '回锅肉',
    image:
      'https://jh-foodie-bucket.s3.us-east-2.amazonaws.com/1720807291594970000.jpeg',
  },
  {
    name: '水煮牛肉',
    image:
      'https://jh-foodie-bucket.s3.us-east-2.amazonaws.com/1720807291594970000.jpeg',
  },
  {
    name: 'Caesar Salad',
    image:
      'https://jh-foodie-bucket.s3.us-east-2.amazonaws.com/1720807291594970000.jpeg',
  },
  {
    name: 'Margherita Pizza',
    image:
      'https://jh-foodie-bucket.s3.us-east-2.amazonaws.com/1720807291594970000.jpeg',
  },
];

export default function MenuItemList() {
  return (
    <div className="grid grid-cols-2 gap-4">
      {dishes.map((dish, index) => (
        <MenuItem key={index} dish={dish} index={index} />
      ))}
    </div>
  );
}

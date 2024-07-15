import SetChefForm from '@/components/Auth/SetChefForm';
import { Card } from '@nextui-org/card';

export default function SetChefAndPartner() {
  return (
    <div className="fixed inset-0 min-h-[95vh] bg-content2 flex flex-col justify-center items-center p-4">
      <Card className="w-full max-w-md p-6 bg-content1 backdrop-blur-sm">
        <h2 className="mb-6 text-2xl font-semibold text-mainLight text-center">
          Set Up Your Profile
        </h2>
        <SetChefForm />
        <p className="mt-6 text-default-500 text-xs text-center">
          This information helps us personalize your Foodie experience.
        </p>
      </Card>
    </div>
  );
}

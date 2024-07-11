import GoogleLogin from '@/components/GoogleLogin';

export default function LoginPage() {
  return (
    <div className="flex flex-col items-center justify-center min-h-[90vh]">
      <p className="text-4xl mb-1">😍</p>
      <p className="text-center mb-4 w-[80%] font-semibold text-lg">
        Log in or create an account to log and order your food
      </p>
      <GoogleLogin />
    </div>
  );
}

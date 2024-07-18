'use client';

import GoogleLogin from '@/components/Auth/GoogleLogin';
import { Button } from '@nextui-org/button';
import { Card } from '@nextui-org/card';
import Image from 'next/image';
import icon from '@/app/logo.png';

export default function LandingPage() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-mainLight to-[#3A95BD] flex flex-col justify-between">
      <div className="flex flex-col items-center justify-center flex-grow text-center px-4 pt-12 pb-20">
        <Image
          src={icon}
          alt="Foodie Logo"
          width={130}
          height={130}
          className="mb-6"
        />
        <h1 className="font-bold mb-4 text-4xl text-white">Foodie</h1>
        <p className="mb-8 text-lg text-white">
          Your Couple&#39;s Culinary Companion
        </p>
        <Card className="w-full max-w-md p-6 bg-white/90 backdrop-blur-sm justify-between items-center">
          <h2 className="mb-4 text-2xl font-semibold text-mainLight">
            Get Started
          </h2>
          <p className="mb-6 text-sm">
            Log in to connect with your partner and start your culinary journey.
          </p>
          <GoogleLogin />
          <p className="mt-6 text-gray-500 text-xs">
            After logging in, you&#39;ll enter your partner&#39;s email and
            choose your role as a chef or food enthusiast.
          </p>
        </Card>
      </div>

      <div className="bg-white/90 backdrop-blur-sm py-8 px-4">
        <h3 className="text-2xl font-semibold mb-6 text-center text-mainLight">
          Why Couples Love Foodie
        </h3>
        <div className="grid grid-cols-2 gap-6">
          <FeatureItem
            icon="🍽️"
            text="Order Together"
            description="Sync your cravings and order as one"
          />
          <FeatureItem
            icon="📝"
            text="Log Meals"
            description="Keep track of your culinary adventures"
          />
          <FeatureItem
            icon="👩‍🍳"
            text="Assign Chef"
            description="Take turns being the kitchen hero"
          />
          <FeatureItem
            icon="❤️"
            text="Couple-Focused"
            description="Tailored for two, perfect for you"
          />
        </div>
      </div>

      {/* How It Works */}
      <div className="bg-white/90 backdrop-blur-sm py-6 px-4">
        <h3 className="text-2xl font-semibold mb-6 text-center text-mainLight">
          How Foodie Works
        </h3>
        <div className="grid grid-cols-2 gap-4">
          <StepItem number={1} text="Sign in with Google" icon="🔐" />
          <StepItem number={2} text="Enter partner's email" icon="✉️" />
          <StepItem number={3} text="Choose your role" icon="👨‍🍳" />
          <StepItem number={4} text="Start your food journey!" icon="🚀" />
        </div>
      </div>

      {/* Call to Action */}
      <div className="bg-white/90 backdrop-blur-sm py-8 px-4 text-center">
        <h3 className="text-2xl font-semibold mb-4 text-mainLight">
          Ready to Spice Up Your Relationship?
        </h3>
        <Button
          color="primary"
          size="lg"
          className="bg-mainLight"
          onClick={() => window.scrollTo({ top: 0, behavior: 'smooth' })}
        >
          Get Started Now
        </Button>
      </div>
    </div>
  );
}

function FeatureItem({
  icon,
  text,
  description,
}: {
  icon: string;
  text: string;
  description: string;
}) {
  return (
    <div className="flex flex-col items-center">
      <p className="mb-2 text-3xl">{icon}</p>
      <p className="text-center font-semibold text-sm mb-1">{text}</p>
      <p className="text-center text-xs text-gray-600">{description}</p>
    </div>
  );
}

function StepItem({
  number,
  text,
  icon,
}: {
  number: number;
  text: string;
  icon: string;
}) {
  return (
    <Card className="p-4 flex flex-col items-center justify-center h-32">
      <div className="flex items-center justify-center w-8 h-8 rounded-full bg-mainLight text-white font-bold mb-2">
        {number}
      </div>
      <p className="text-2xl mb-1">{icon}</p>
      <p className="text-center text-sm">{text}</p>
    </Card>
  );
}

'use client';

import React, { useState } from 'react';
import { useFormStatus } from 'react-dom';
import { Switch } from '@nextui-org/switch';
import { Input } from '@nextui-org/input';
import { Button } from '@nextui-org/button';
import { EnvelopeIcon } from '@heroicons/react/24/solid';
import { setChef } from './SetChefAction';
import { useFormState } from 'react-dom';
import { useRouter } from 'next/navigation';
import { ExclamationCircleIcon } from '@heroicons/react/24/outline';
import { Card, CardBody } from '@nextui-org/card';
import { CheckCircleIcon } from '@heroicons/react/24/outline';

const initialState = {
  success: false,
  message: '',
};

function SuccessButton() {
  const router = useRouter();

  return (
    <Button
      size="lg"
      className="w-full bg-mainLight dark:bg-mainDark text-white font-medium text-lg"
      onPress={() => router.push('/menuItems')}
    >
      Get Started
    </Button>
  );
}

function SubmitButton() {
  const { pending } = useFormStatus();
  return (
    <Button
      type="submit"
      size="lg"
      className="w-full bg-mainLight dark:bg-mainDark text-white font-medium text-lg"
      isLoading={pending}
    >
      Submit
    </Button>
  );
}

export default function SetChefForm() {
  const [isChef, setIsChef] = useState(false);
  const [partnerEmail, setPartnerEmail] = useState('');

  const [state, formAction] = useFormState(setChef, initialState);

  return (
    <form action={formAction} className="space-y-6">
      <div className="flex items-center justify-between">
        <label
          htmlFor="chef-switch"
          className="text-default-700 font-medium text-lg"
        >
          Are you the chef?
        </label>
        <Switch
          size="lg"
          id="chef-switch"
          color="primary"
          className="scale-75"
          isSelected={isChef}
          onValueChange={setIsChef}
          isDisabled={state?.success}
        />
      </div>
      <input type="hidden" name="is-chef" value={isChef ? 'true' : 'false'} />

      <Input
        value={partnerEmail}
        onValueChange={setPartnerEmail}
        name="partner-email"
        isRequired
        label="Partner Email"
        placeholder="Enter your partner's email"
        startContent={
          <EnvelopeIcon className="w-5 h-5 text-default-400 pointer-events-none flex-shrink-0" />
        }
        endContent={
          <div className="pointer-events-none flex items-center">
            <span className="text-default-400 text-sm">@gmail.com</span>
          </div>
        }
        isDisabled={state?.success}
      />
      {state?.message && !state?.success && (
        <Card className="bg-danger-50" shadow="none">
          <CardBody className="p-3">
            <div className="flex items-center space-x-2">
              <ExclamationCircleIcon className="w-6 h-6 text-danger-600" />
              <p className="text-danger-600">{state.message}</p>
            </div>
          </CardBody>
        </Card>
      )}
      {state?.message && state?.success && (
        <Card className="bg-success-50" shadow="none">
          <CardBody className="p-3">
            <div className="flex items-center space-x-2">
              <CheckCircleIcon className="w-6 h-6 text-success-600" />
              <p className="text-success-600">{state.message}</p>
            </div>
          </CardBody>
        </Card>
      )}
      {state?.success ? <SuccessButton /> : <SubmitButton />}
    </form>
  );
}

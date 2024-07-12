'use client';

import React from 'react';
import { Modal, ModalContent, ModalHeader, ModalBody } from '@nextui-org/modal';
import { Button } from '@nextui-org/button';
import { useDisclosure } from '@nextui-org/react';
import { Cog6ToothIcon } from '@heroicons/react/24/outline';
import LogoutButton from '@/components/LogoutButton';
import { ThemeSwitcher } from '@/components/theme-switch';

export default function SettingsModal() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  return (
    <>
      <Button
        variant="light"
        className="justify-start py-2 mb-2 text-md"
        startContent={<Cog6ToothIcon className="w-6 h-6 mr-2" />}
        onPress={onOpen}
      >
        Settings
      </Button>
      <Modal
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        placement="center"
        size="sm"
      >
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className="flex flex-col">Settings</ModalHeader>
              <ModalBody className="mb-4">
                <ThemeSwitcher />
                <LogoutButton />
              </ModalBody>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}

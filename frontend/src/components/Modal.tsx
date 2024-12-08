import React from "react";
import {
  Dialog,
  DialogBackdrop,
  DialogPanel,
  DialogTitle,
} from "@headlessui/react";

type ModalProps = {
  isOpen: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  buttonText: string;
  children: React.ReactNode;
  title: string;
};

export default function Modal({
  isOpen,
  setIsOpen,
  buttonText,
  children,
  title,
}: ModalProps) {
  return (
    <>
      <button
        onClick={() => setIsOpen(true)}
        className="border px-3 py-1.5 rounded-lg bg-blue-300"
      >
        {buttonText}
      </button>
      <Dialog
        open={isOpen}
        onClose={() => setIsOpen(false)}
        className="relative z-50"
      >
        <DialogBackdrop className="fixed inset-0 bg-black/30" />
        <div className="fixed inset-0 md:ml-32 flex w-screen items-center justify-center p-4">
          <DialogPanel className="space-y-4 border bg-white p-12 w-1/2 rounded-md">
            <DialogTitle className="font-bold">{title}</DialogTitle>
            {children}
          </DialogPanel>
        </div>
      </Dialog>
    </>
  );
}

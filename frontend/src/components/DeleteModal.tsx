import React, { useState } from "react";
import {
  Dialog,
  DialogBackdrop,
  DialogPanel,
  DialogTitle,
} from "@headlessui/react";
import { FaTrash } from "react-icons/fa";

type ModalProps = {
  isOpen: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  children: React.ReactNode;
  title: string;
};

function Modal({ isOpen, setIsOpen, children, title }: ModalProps) {
  return (
    <>
      <button onClick={() => setIsOpen(true)}>
        <FaTrash className="text-red-500" />
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

export default function DeleteModal({
  title,
  onPositive,
}: { title: string; onPositive: () => void }) {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <Modal isOpen={isOpen} setIsOpen={setIsOpen} title={title}>
      <div className="flex flex-row items-center gap-5">
        <button
          onClick={() => setIsOpen(false)}
          className="border px-3 py-1.5 rounded-lg border-red-500 w-full"
        >
          Batal
        </button>
        <button
          onClick={() => {
            onPositive();
            setIsOpen(false);
          }}
          className="border px-3 py-1.5 rounded-lg bg-red-500 w-full"
        >
          Yakin
        </button>
      </div>
    </Modal>
  );
}

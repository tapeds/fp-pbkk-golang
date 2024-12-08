import React, { useState } from "react";
import {
  Dialog,
  DialogBackdrop,
  DialogPanel,
  DialogTitle,
} from "@headlessui/react";
import { FaEdit, FaTrash } from "react-icons/fa";
import { FormProvider, useForm } from "react-hook-form";

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
        <FaEdit />
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

export default function EditModal({
  title,
  children,
  onSubmit,
  data,
}: {
  title: string;
  children: React.ReactNode;
  onSubmit: (data: any) => void;
  data: any;
}) {
  const [isOpen, setIsOpen] = useState(false);

  const methods = useForm({
    defaultValues: {
      ...data,
    },
  });

  const { handleSubmit } = methods;

  return (
    <Modal isOpen={isOpen} setIsOpen={setIsOpen} title={title}>
      <FormProvider {...methods}>
        <form className="flex flex-col gap-3" onSubmit={handleSubmit(onSubmit)}>
          {children}

          <button
            className="border px-3 py-1.5 rounded-lg bg-blue-400 text-white"
            type="submit"
          >
            Submit
          </button>
        </form>
      </FormProvider>
    </Modal>
  );
}

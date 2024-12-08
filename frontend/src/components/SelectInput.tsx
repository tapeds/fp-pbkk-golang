import * as React from "react";
import { get, RegisterOptions, useFormContext } from "react-hook-form";
import { FiChevronDown } from "react-icons/fi";

import clsx from "clsx";

import ErrorMessage from "./ErrorMessage";
import LabelText from "./LabelText";

export type SelectInputProps = {
  id: string;
  label?: string;
  helperText?: string;
  hideError?: boolean;
  validation?: RegisterOptions;
  readOnly?: boolean;
  placeholder?: string;
} & React.ComponentPropsWithoutRef<"select">;

export default function SelectInput({
  id,
  label,
  helperText,
  hideError = false,
  validation,
  className,
  readOnly = false,
  defaultValue = "",
  placeholder = "",
  children,
  ...rest
}: SelectInputProps) {
  const {
    register,
    formState: { errors },
    watch,
  } = useFormContext();

  const error = get(errors, id);
  const value = watch(id);

  return (
    <div className="w-full space-y-1.5 rounded-md">
      {label && (
        <LabelText required={validation?.required ? true : false}>
          {label}
        </LabelText>
      )}

      <div className="relative">
        <select
          {...register(id, validation)}
          id={id}
          name={id}
          defaultValue={defaultValue}
          disabled={readOnly}
          className={clsx(
            "w-full appearance-none truncate rounded-md border-none py-2.5 pl-3 pr-8",
            "ring-1 ring-typo-outline-1 focus:ring-typo-outline-1",
            "bg-typo-white font-poppins text-sm text-typo-secondary",
            "hover:ring-2 hover:ring-typo-main",
            readOnly && "cursor-not-allowed",
            error
              ? "ring-1 ring-inset ring-danger-main focus:ring-danger-main"
              : "focus:ring-typo-outline-1",
            value && "ring-primary-info-active focus:ring-primary-info-active",
            className,
          )}
          aria-describedby={id}
          {...rest}
        >
          {placeholder && (
            <option value="" disabled hidden>
              {placeholder}
            </option>
          )}
          {children}
        </select>
        <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
          <FiChevronDown className="text-xl text-typo-outline-1" />
        </div>
      </div>

      {!hideError && error && <ErrorMessage>{error.message}</ErrorMessage>}
    </div>
  );
}

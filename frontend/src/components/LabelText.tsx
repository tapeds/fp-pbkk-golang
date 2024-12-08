import { ReactNode } from "react";

import clsx from "clsx";

export default function LabelText({
  children,
  labelTextClasname,
  required,
}: {
  children: ReactNode;
  labelTextClasname?: string;
  required?: boolean;
}) {
  return (
    <label>
      <p className={clsx("text-xs text-black", labelTextClasname)}>
        {children} {required && <span className="text-red-500">*</span>}
      </p>
    </label>
  );
}

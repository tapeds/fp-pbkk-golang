import clsx from "clsx";
import React from "react";

export default function AdminLayout({
  children,
  className,
}: { children: React.ReactNode; className?: string }) {
  return (
    <main
      className={clsx(
        "p-6 min-h-screen bg-gray-100 flex justify-center items-center overflow-hidden",
        className,
      )}
    >
      {children}
    </main>
  );
}

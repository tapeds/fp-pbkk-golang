import clsx from "clsx";
import React from "react";
import Navbar from "./Navbar";

export default function Layout({
  children,
  className,
}: { children: React.ReactNode; className?: string }) {
  return (
    <>
    <Navbar />
      <main
        className={clsx(
          "p-6 min-h-screen bg-gray-100 flex justify-center items-center overflow-hidden",
          className,
        )}
      >
        {children}
      </main>
    </>
  );
}

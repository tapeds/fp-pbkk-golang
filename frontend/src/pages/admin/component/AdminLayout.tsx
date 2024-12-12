import clsx from "clsx";
import React, { useEffect } from "react";
import Sidebar from "./Sidebar";
import { useCookies } from "react-cookie";

export default function AdminLayout({
  children,
  className,
}: { children: React.ReactNode; className?: string }) {
  const [cookie] = useCookies(["user"]);

  return (
    <>
      {cookie.user.role == "admin" ? (
        <>
          <Sidebar />
          <main
            className={clsx(
              "p-6 min-h-screen bg-gray-100 flex justify-center items-center overflow-hidden md:ml-64",
              className,
            )}
          > 
            {children}
          </main>
        </>
      ) : (
        <></>
      )}
    </>
  );
}

import { useState } from "react";
import { useCookies } from "react-cookie";

import { IoMdClose } from "react-icons/io";
import { RxHamburgerMenu } from "react-icons/rx";
import { useNavigate } from "react-router-dom";

export default function Sidebar() {
  const [isOpen, setIsOpen] = useState(false);

  const navigate = useNavigate();
  const [, , removeCookie] = useCookies(["user"]);

  const sidebarItems = [
    { label: "Home", href: "/admin" },
    { label: "Penerbangan", href: "/admin/penerbangan" },
    { label: "Maskapai", href: "/admin/maskapai" },
    { label: "Bandara", href: "/admin/bandara" },
  ];
  return (
    <div>
      {/* Mobile Menu Toggle */}
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="md:hidden fixed top-4 left-4 z-50 p-2 bg-gray-200 rounded-md"
      >
        {isOpen ? <IoMdClose /> : <RxHamburgerMenu />}
      </button>

      {/* Sidebar */}
      <div
        className={`
          fixed top-0 left-0 h-full w-64 
          bg-white shadow-lg transition-transform duration-300
          ${isOpen ? "translate-x-0" : "-translate-x-full"}
          md:translate-x-0
          z-40
        `}
      >
        <div className="flex flex-col justify-between p-4 h-full">
          <div>
            {/* App Title */}
            <div className="text-2xl font-bold mb-8 text-center">Admin</div>

            {/* Navigation */}
            <nav className="space-y-2">
              {sidebarItems.map((item, index) => (
                <a
                  key={index}
                  href={item.href}
                  className="
                  flex items-center p-3 
                  hover:bg-gray-100 
                  rounded-md 
                  transition-colors 
                  group
                "
                >
                  <span className="text-gray-700 group-hover:text-blue-600">
                    {item.label}
                  </span>
                </a>
              ))}
            </nav>
          </div>

          <button
            onClick={() => {
              removeCookie("user");
              navigate("/login");
            }}
            className="border px-3 py-1.5 rounded-lg bg-red-300"
          >
            Logout
          </button>
        </div>
      </div>

      {/* Mobile Overlay */}
      {isOpen && (
        <div
          className="fixed inset-0 bg-black/50 md:hidden z-30"
          onClick={() => setIsOpen(false)}
        />
      )}
    </div>
  );
}

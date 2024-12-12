import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useCookies } from "react-cookie";
// import { MenuIcon, XIcon } from "@heroicons/react/outline"; // Import icon Tailwind untuk menu
// import { MenuIcon, XIcon } from '@heroicons/react/24/outline';
import { FaHome, FaUserAlt, FaBars } from 'react-icons/fa';

// export default function Header() {
//   return (
//     <div className="header">
//       
//       <FaUserAlt size={24} />
//       <FaBars size={24} />
//     </div>
//   );
// }



export default function Navbar() {
  const [cookies] = useCookies(["authToken"]);
  const isLoggedIn = Boolean(cookies.authToken);
  const [menuOpen, setMenuOpen] = useState(false);

  const toggleMenu = () => {
    setMenuOpen(!menuOpen);
  };

  return (
    <header className="bg-gray-800 text-white p-4">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">
        </h1>
        <button
          onClick={toggleMenu}
          className="lg:hidden block text-white focus:outline-none"
        >
          {menuOpen ? (
            <FaBars size={24} />
          ) : (
            <FaBars size={24} />
          )}
        </button>
        <nav className="hidden lg:flex gap-4">
          <Link to="/" className="hover:underline">
            Home
          </Link>
          {isLoggedIn && (
            <Link to="/pesanan" className="hover:underline">
              Pesanan
            </Link>
          )}
          <Link to={isLoggedIn ? "/profil" : "/login"} className="hover:underline">
            {isLoggedIn ? "Profil" : "Login"}
          </Link>
        </nav>
      </div>

      {menuOpen && (
        <nav className="lg:hidden flex flex-col gap-2 mt-4 bg-gray-700 p-4 rounded-lg">
          <Link
            to="/"
            className="hover:underline"
            onClick={() => setMenuOpen(false)}
          >
            Home
          </Link>
          {isLoggedIn && (
            <Link
              to="/pesanan"
              className="hover:underline"
              onClick={() => setMenuOpen(false)}
            >
              Pesanan
            </Link>
          )}
          <Link
            to={isLoggedIn ? "/profil" : "/login"}
            className="hover:underline"
            onClick={() => setMenuOpen(false)}
          >
            {isLoggedIn ? "Profil" : "Login"}
          </Link>
        </nav>
      )}
    </header>
  );
}

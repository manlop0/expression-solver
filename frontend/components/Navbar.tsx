"use client"

import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

export default function Navbar() {
  const links = [
    { label: "Expressions List", href: "/" },
    { label: "Operations Settings", href: "/operations" },
    { label: "Copmuters List", href: "/workers" },
  ];

  const currentPath = usePathname()

  return (
    <nav className="w-full h-16  items-center flex justify-around shadow-xl">
      {links.map((link) => (
        <Link
          href={link.href}
          key={link.href}
          className={` border-x-black border-x flex flex-auto h-full items-center justify-center ${currentPath === link.href ? "bg-white text-black" : " bg-secondary  hover:bg-blue-400"}`}
        >
          {link.label}
        </Link>
      ))}
    </nav>
  );
}

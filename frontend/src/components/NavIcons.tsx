"use client";
import Link from "next/link";

import Image from "next/image";
import { useEffect, useState } from "react";

const NavIcons = () => {
  const [hasToken, setHasToken] = useState(false);

  useEffect(() => {
    const token = document.cookie
      .split("; ")
      .find((row) => row.startsWith("jwt="));

    if (token) {
      setHasToken(true);
    }
  }, []);

  return (
    <div className="flex items-center gap-4 xl:gap-6">
      {hasToken && (
        <>
          <Image src="/basket.png" alt="Basket" width={28} height={28} className="cursor-pointer" />
          <Image src="/account.png" alt="Account" width={28} height={28} className="cursor-pointer" />
        </>
      )}

      {!hasToken && (
        <>
          <button className="w-24 font-medium border-2 border-accent text-white bg-transparent px-4 py-2 rounded-md hover:bg-accent hover:text-white hover:opacity-80 transition">
            <Link href="/login">Masuk</Link>
          </button>
          <button className="w-24 font-medium bg-accent border-2 border-accent text-white px-4 py-2 rounded-md hover:opacity-80 transition">
            <Link href="/register">Daftar</Link>
          </button>
        </>
      )}
    </div>
  );
};

export default NavIcons;
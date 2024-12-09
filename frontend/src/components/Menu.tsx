"use client";

import Image from "next/image";
import Link from "next/link";
import { useState } from "react";

const Menu = () => {
    const [open, setOpen] = useState(false);

    return (
        <div className="">
            <Image
                src="/menu.png"
                alt=""
                width={24}
                height={24}
                className="cursor-pointer"
                onClick={() => setOpen((prev) => !prev)}
            />
            {open && (
                <div className="absolute bg-secondary-dark text-white right-0 top-20 w-max h-auto flex flex-col items-start gap-8 text-xl z-10 p-4">
                    <Link href="/">
                        <div className="flex items-center gap-4 font-medium">
                            <Image src="/basket.png" alt="" width={28} height={28} className="cursor-pointer"/>
                            Keranjang
                        </div>
                    </Link>
                    <Link href="/">
                    <div className="flex items-center gap-4 font-medium">
                        <Image src="/account.png" alt="" width={28} height={28} className="cursor-pointer"/>
                        Akun
                    </div>
                    </Link>
                    <Link href="/">
                        <button className="w-24 font-medium border-2 border-accent text-white bg-transparent px-4 py-2 rounded-md hover:bg-accent hover:text-white hover:opacity-80 transition">
                            Masuk
                        </button>
                    </Link>
                    <Link href="/">
                        <button className="w-24 font-medium bg-accent text-white px-4 py-2 rounded-md hover:opacity-80 transition">
                            Daftar
                        </button>
                    </Link>
                </div>
            )}
        </div>
    );
};

export default Menu;
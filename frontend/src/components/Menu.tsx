"use client";

import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";

const Menu = () => {
    const [open, setOpen] = useState(false);

    const [hasToken, setHasToken] = useState(false);

        useEffect(() => {
            const token = document.cookie
            .split("; ")
            .find((row) => row.startsWith("token"));

            if (token) {
            setHasToken(true);
            }
        }, []);

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
                <div className="fixed rounded:md bg-secondary-dark text-white right-0 top-16 w-400 h-100 flex flex-col items-start gap-4 text-xl p-4">
                    {hasToken && (
                    <>
                        <Link href="/">
                        <div className="text-sm flex items-center gap-4 font-medium">
                            <Image src="/basket.png" alt="" width={24} height={24} className="cursor-pointer"/>
                            Keranjang
                        </div>
                        </Link>
                        <Link href="/">
                        <div className="text-sm flex items-center gap-4 font-medium">
                            <Image src="/account.png" alt="" width={24} height={24} className="cursor-pointer"/>
                            Akun
                        </div>
                        </Link>
                    </>
                    )}

                    {!hasToken && (
                        <>
                        <button className="w-20 font-medium border-2 border-accent text-white bg-transparent px-2 py-1 text-sm rounded-md hover:bg-accent hover:text-white hover:opacity-80 transition">
                        Masuk
                        </button>
                        <button className="w-20 font-medium bg-accent border-2 border-accent text-white px-2 py-1 text-sm rounded-md hover:opacity-80 transition">
                        Daftar
                        </button>
                        </>
                    )}                                  
                </div>
            )}
        </div>
    );
};

export default Menu;
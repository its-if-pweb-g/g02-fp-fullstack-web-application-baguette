"use client";

import Image from "next/image";
import { useRouter } from "next/navigation";

const SearchBar = () => {
    const router = useRouter();

    const handleSearch = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const name = formData.get("name") as string;
    
        if(name){
          router.push(`/list?name=${name}`)
        }
    };

    return (
        <form className="flex items-center justify-between gap-2 p-1 md:gap-4 md:p-2 border-2 border-white rounded-full flex-1" onSubmit={handleSearch}>
            <button className="cursor-pointer hover:opacity-80 transition">
                <Image src="/search.png" alt="Search" width={20} height={20} className="md:w-7 md:h-7" />
            </button>
            <input className="flex-1 bg-transparent outline-none text-white placeholder-slate-300 text-sm md:text-base" type="text" placeholder="Cari produk tertentu..."/>
        </form>
    );
  };
  
export default SearchBar;
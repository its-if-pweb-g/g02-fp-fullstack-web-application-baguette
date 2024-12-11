"use client";

import { useState, useRef, useEffect } from "react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import Filter from "./Filter";
import { API_URL } from "@/utils/config";

const SearchBar = () => {
  const router = useRouter();
  const [showFilter, setShowFilter] = useState(false);
  const [filters, setFilters] = useState<{ type: string[]; flavor: string[]; price: string[] }>({
    type: [],
    flavor: [],
    price: [],
  });
  const [searchTerm, setSearchTerm] = useState("");

  const handleSearch = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const params = new URLSearchParams();
    if (searchTerm) params.append("q", searchTerm);
    if (filters.type.length) params.append("type", filters.type.join(","));
    if (filters.flavor.length) params.append("flavor", filters.flavor.join(","));
    if (filters.price.length) params.append("price", filters.price.join(","));

    router.push(`/products?${params.toString()}`);
  };

  const filterRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (filterRef.current && !filterRef.current.contains(event.target as Node)) {
        setShowFilter(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  return (
    <div className="relative w-full" ref={filterRef}>
      <form
        className="flex items-center justify-between gap-2 p-1 md:gap-4 md:p-2 border-2 border-white rounded-md flex-1"
        onSubmit={handleSearch}
      >
        <input
          className="flex-1 bg-transparent outline-none text-white placeholder-slate-300 text-sm md:text-base"
          type="text"
          placeholder="Cari produk tertentu..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          onFocus={() => setShowFilter(true)}
        />
        <button type="submit" className="cursor-pointer hover:opacity-80 transition">
          <Image src="/search.png" alt="Search" width={20} height={20} className="md:w-7 md:h-7" />
        </button>
      </form>

      {showFilter && (
        <div className="absolute top-full mt-2 left-0 w-full bg-secondary-dark rounded-md z-10">
          <Filter onFilterChange={setFilters} />
        </div>
      )}
    </div>
  );
};

export default SearchBar;
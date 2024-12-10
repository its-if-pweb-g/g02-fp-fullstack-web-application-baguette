"use client";

import { useState, useEffect } from "react";

const Filter = ({ onFilterChange }: { onFilterChange: (filters: any) => void }) => {
  const [selectedFilters, setSelectedFilters] = useState<{
    type: string[];
    flavor: string[];
    price: string[];
  }>({
    type: [],
    flavor: [],
    price: [],
  });

  const handleCheckboxChange = (category: string, value: string) => {
    setSelectedFilters((prev) => {
      const currentCategory = prev[category as keyof typeof selectedFilters];
      const newCategory = currentCategory.includes(value)
        ? currentCategory.filter((item) => item !== value)
        : [...currentCategory, value];

      return { ...prev, [category]: newCategory };
    });
  };

  useEffect(() => {
    onFilterChange(selectedFilters);
  }, [selectedFilters, onFilterChange]);

  return (
    <div className="bg-secondary-dark text-white p-4 rounded-md shadow-md">
      <div className="mb-4">
        <h2 className="font-bold mb-2">Tipe Roti</h2>
        <div className="flex flex-wrap gap-4">
          {["baguettes", "bagel", "cupcakes", "cakes", "rolls", "loaf", "cookies", "muffins", "danish"].map((type) => (
            <label key={type} className="flex items-center gap-2 cursor-pointer">
              <input
                type="checkbox"
                onChange={() => handleCheckboxChange("type", type)}
                className="accent-accent"
              />
              {type.charAt(0).toUpperCase() + type.slice(1)}
            </label>
          ))}
        </div>
      </div>

      <div className="mb-4">
        <h2 className="font-bold mb-2">Rasa</h2>
        <div className="flex flex-wrap gap-4">
          {["original", "coklat","keju", "stroberi", "mint", "red velvet", "kacang", "kayu manis"].map((flavor) => (
            <label key={flavor} className="flex items-center gap-2 cursor-pointer">
              <input
                type="checkbox"
                onChange={() => handleCheckboxChange("flavor", flavor)}
                className="accent-accent"
              />
              {flavor.charAt(0).toUpperCase() + flavor.slice(1)}
            </label>
          ))}
        </div>
      </div>

      <div className="mb-4">
        <h2 className="font-bold mb-2">Harga</h2>
        <div className="flex flex-wrap gap-4">
          {[
            { label: "< Rp 30.000", value: "1" },
            { label: "Rp 30.000 - Rp 60.000", value: "2" },
            { label: "> Rp 60.000", value: "3" },
          ].map((price) => (
            <label key={price.value} className="flex items-center gap-2 cursor-pointer">
              <input
                type="checkbox"
                onChange={() => handleCheckboxChange("price", price.value)}
                className="accent-accent"
              />
              {price.label}
            </label>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Filter;
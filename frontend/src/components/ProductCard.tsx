import React from 'react';
import { API_URL } from '@/utils/config';

interface ProductCardProps {
  name: string;
  header_description: string;
  price: number;
  id: string;
  sold: number;
}

const ProductCard: React.FC<ProductCardProps> = ({ name, header_description, price, id, sold }) => {
  return (
    <div className="w-80 h-[400px] bg-secondary-light rounded-lg shadow-lg flex flex-col">
  
      <img
        src={`${API_URL}/api/image/${id}`}
        alt={name}
        className="rounded-t-lg w-full h-44 object-cover"
      />

      <div className="p-4 flex flex-col justify-between flex-grow">
        <h2 className="text-customBlack text-xl font-bold text-center">{name}</h2>
        <p className="text-md text-center text-customBlack font-medium mt-2">{header_description}</p>
        <p className="text-customBlack mt-4 text-2xl font-bold text-center">{`Rp ${price.toLocaleString()}`}</p>

        <div className="flex items-center justify-between mt-4 text-sm">
          <span className="text-customBlack text-md font-medium">{sold} terjual</span>
          <button className="px-4 py-2 bg-[#FFC857] text-white font-bold rounded-full">
            Tambah
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { API_URL } from '@/utils/config';
import Cookies from 'js-cookie';
import {jwtDecode} from 'jwt-decode';

interface ProductCardProps {
  name: string;
  header_description: string;
  price: number;
  id: string;
  sold: number;
}

const ProductCard: React.FC<ProductCardProps> = ({ name, header_description, price, id, sold }) => {
  const router = useRouter();
  const [isAdmin, setIsAdmin] = useState(false);

  useEffect(() => {
    const token = Cookies.get('token');
    if (token) {
      try {
        const decoded: any = jwtDecode(token);
        setIsAdmin(decoded.role === 'admin');
      } catch (error) {
        console.error('Failed to decode token:', error);
      }
    }
  }, []);

  const handleCardClick = () => {
    if (isAdmin) {
      router.push(`/admin/product/${id}`);
    } else {
      router.push(`/product/${id}`);
    }
  };

  const handleAddToCart = async (q: number = 1) => {
    const token = Cookies.get('token');

    if (!token) {
      router.push('/register');
      return;
    }

    try {
      const productData = {
        id,
        name,
        price,
        quantity: q,
      };

      const response = await fetch(`${API_URL}/api/user/cart/products`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${Cookies.get('token')}`, 
        },
        body: JSON.stringify(productData),
      });

      if (!response.ok) {
        throw new Error('Failed to add product to cart');
      }

      alert('Produk berhasil ditambahkan ke keranjang');
    } catch (error) {
      alert('Terjadi kesalahan pada server');
    }
  };

  return (
    <div
      className="w-80 h-[400px] bg-secondary-light rounded-lg shadow-lg flex flex-col cursor-pointer"
      onClick={handleCardClick}
    >
      <img
        src={`${API_URL}/api/image/${id}`}
        alt={name}
        className="rounded-t-lg w-full h-44 object-cover"
      />

      <div className="p-4 flex flex-col justify-between flex-grow">
        <h2 className="text-customBlack text-xl font-bold text-center">{name}</h2>
        <p className="text-md text-center text-customBlack font-medium mt-2">{header_description}</p>
        <p className="text-customBlack mt-4 text-2xl font-bold text-center">{`Rp ${price.toLocaleString()}`}</p>

        <div className="flex items-center justify-between mt-4 text-sm relative">
          <span className="text-customBlack text-md font-medium">{sold} terjual</span>
          <button
            className="px-4 py-2 shadow-lg bg-[#FFC857] text-white font-bold rounded-full z-10"
            onClick={() => handleAddToCart(1)}
          >
            Tambah
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;
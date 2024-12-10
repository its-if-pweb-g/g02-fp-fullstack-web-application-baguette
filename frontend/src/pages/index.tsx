import React, { useEffect, useState } from 'react';
import ProductCard from '@/components/ProductCard';
import { API_URL } from '@/utils/config';
import Navbar from '@/components/Navbar';
import { Poppins } from "next/font/google";

const poppins = Poppins({
  subsets: ['latin'],
  weight: ['400', '700'], 
  variable: '--font-poppins', 
});

interface Product {
  name: string;
  header_description: string;
  price: number;
  id: string;
  sold: number;
}

const HomePage = () => {
  const [topSellerProducts, setTopSellerProducts] = useState<Product[]>([]);
  const [newestProducts, setNewestProducts] = useState<Product[]>([]);
  const [allProducts, setAllProducts] = useState<Product[]>([]);
  const [visibleCount, setVisibleCount] = useState(8);
  const [loading, setLoading] = useState<boolean>(true); 
  const [error, setError] = useState<string | null>(null); 

  useEffect(() => {
    const fetchTopSellerProducts = async () => {
      try {
        const response = await fetch(`${API_URL}/api/products/sort?top=4`, {
          headers: {
            'ngrok-skip-browser-warning': 'true',
          },
        });
    
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
    
        const data = await response.json();
        setTopSellerProducts(data);
      } catch (error: any) {
        setError('Error fetching top seller products');
        console.error(error);
      } finally {
        setLoading(false);
      }
    };

    fetchTopSellerProducts();
  }, []);

  useEffect(() => {
    const fetchNewestProducts = async () => {
      try {
        const response = await fetch(`${API_URL}/api/products/sort?new=4`, {
          headers: {
            'ngrok-skip-browser-warning': 'true',
          },
        });
        const data = await response.json();
        setNewestProducts(data);
      } catch (error: any) {
        setError('Error fetching newest products');
        console.error(error);
      } finally {
        setLoading(false);
      }
    };

    fetchNewestProducts();
  }, []);

  useEffect(() => {
    const fetchAllProducts = async () => {
      try {
        const response = await fetch(`${API_URL}/api/products`, {
          headers: {
            'ngrok-skip-browser-warning': 'true',
          },
        });
        const data = await response.json();
        setAllProducts(data);
      } catch (error: any) {
        setError('Error fetching all products');
        console.error(error);
      } finally {
        setLoading(false);
      }
    };

    fetchAllProducts();
  }, []);

  const handleShowMore = () => {
    setVisibleCount(visibleCount + 8);
  };

  if (loading) {
    return <div>Loading...</div>; 
  }

  if (error) {
    return <div>{error}</div>; 
  }

  return (
    <div>
      <Navbar /> 
      <div className="container mx-auto text-customBlack">
        <section>
          <h2 className="text-2xl font-bold mb-4">Top Seller Products</h2>
          <div className="flex gap-4 flex-wrap">
            {topSellerProducts.map((product) => (
              <ProductCard key={product.id} {...product} />
            ))}
          </div>
        </section>

        <section className="mt-8">
          <h2 className="text-2xl font-bold mb-4">Newest Products</h2>
          <div className="flex gap-4 flex-wrap">
            {newestProducts.map((product) => (
              <ProductCard key={product.id} {...product} />
            ))}
          </div>
        </section>

        <section className="mt-8">
          <h2 className="text-2xl font-bold mb-4">All Products</h2>
          <div className="flex gap-4 flex-wrap">
            {allProducts.slice(0, visibleCount).map((product) => (
              <ProductCard key={product.id} {...product} />
            ))}
          </div>
          {allProducts.length > visibleCount && (
            <button
              className="mt-4 px-4 py-2 bg-accent text-white font-bold rounded-md hover:opacity-80 transition"
              onClick={handleShowMore}
            >
              Show More
            </button>
          )}
        </section>
      </div>
    </div>
  );
};

export default HomePage;

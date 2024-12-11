import React, { useEffect, useState } from 'react';
import ProductCard from '@/components/ProductCard';
import { API_URL } from '@/utils/config';
import Navbar from '@/components/Navbar';
import { Poppins } from "next/font/google";
import HeroSection from '@/components/HeroSection';
import Footer from '@/components/Footer';

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
      <HeroSection />
      <div className="mx-auto text-customBlack px-4 mt-8 mb-8">
        <section>
          <div className="max-w-7xl mx-auto">
            <h1 className="text-3xl font-bold mb-4 text-customBlack">Produk Top Seller</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
              {topSellerProducts.map((product) => (
                <ProductCard key={product.id} {...product} />
              ))}
            </div>
          </div>
        </section>

        <section className="mt-8">
          <div className="max-w-7xl mx-auto">
            <h1 className="text-3xl font-bold mb-4 text-customBlack">Produk Terbaru</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
              {newestProducts.map((product) => (
                <ProductCard key={product.id} {...product} />
              ))}
            </div>
          </div>
        </section>

        <section className="mt-8">
          <div className="max-w-7xl mx-auto">
            <h1 className="text-3xl font-bold mb-4 text-customBlack">Semua Produk</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
              {allProducts.slice(0, visibleCount).map((product) => (
                <ProductCard key={product.id} {...product} />
              ))}
            </div>
            {allProducts.length > visibleCount && (
              <div className="flex justify-center mt-4 mb-4">
                <button
                  className="px-4 py-2 bg-accent text-white font-bold rounded-md hover:opacity-80 transition shadow-lg"
                  onClick={handleShowMore}
                >
                  Tampilkan Lebih
                </button>
              </div>
            )}
          </div>
        </section>
      </div>
      <Footer/>
    </div>
  );
};

export default HomePage;

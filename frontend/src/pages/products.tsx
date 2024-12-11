"use client";

import { useEffect, useState } from "react";
import { useSearchParams } from "next/navigation";
import { API_URL } from "@/utils/config";
import ProductCard from "@/components/ProductCard";
import Navbar from "@/components/Navbar";
import Footer from "@/components/Footer";

interface Product {
  name: string;
  header_description: string;
  price: number;
  id: string;
  sold: number;
};

const ProductPage = () => {
  const searchParams = useSearchParams();
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchProducts = async () => {
      const params = new URLSearchParams(searchParams as any);

      try {
        const response = await fetch(`${API_URL}/api/products/search?${params.toString()}`, {
          headers: {
            'ngrok-skip-browser-warning': 'true',
          },
        });
        const data = await response.json();
        setProducts(data);
      } catch (error) {
        console.error("Error fetching products:", error);
      } finally {
        setLoading(false);
      }
    };

    if (searchParams) {
      fetchProducts();
    }
  }, [searchParams]);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <Navbar />
      <div className="max-w-7xl mx-auto mt-8 mb-8">
        <h1 className="text-3xl font-bold text-customBlack">Hasil Pencarian</h1>
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8 mt-6">
          {products && products.length > 0 ? (
            products.map((product) => (
              <ProductCard key={product.id} {...product} />
            ))
          ) : (
            <p className="font-medium text-customBlack">Tidak ada produk yang sesuai.</p>
          )}
        </div>
      </div>
      <Footer/>
    </div>
  );
};

export default ProductPage;
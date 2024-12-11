import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { API_URL } from '@/utils/config';
import Cookies from 'js-cookie';
import { jwtDecode } from 'jwt-decode';
import Navbar from '@/components/Navbar';

interface Product {
  name: string;
  header_description: string;
  description: string;
  price: number;
  id: string;
  sold: number;
  stock: number;
  image: string;
}

const AdminProductEdit = () => {
  const router = useRouter();
  const { id } = router.query;
  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [isAdmin, setIsAdmin] = useState<boolean>(false);
  const [formData, setFormData] = useState<Product>({
    name: '',
    header_description: '',
    description: '',
    price: 0,
    id: '',
    sold: 0,
    stock: 0,
    image: '',
  });

  const checkAdminRole = () => {
    const token = Cookies.get('token');
    if (token) {
      try {
        const decoded: any = jwtDecode(token);
        if (decoded?.role === 'admin') {
          setIsAdmin(true);
        } else {
          router.push('/');
        }
      } catch (error) {
        console.error('Error decoding JWT token', error);
        router.push('/');
      }
    } else {
      router.push('/');
    }
  };

  const fetchProductData = async () => {
    if (id) {
      try {
        const response = await fetch(`${API_URL}/api/products/${id}`, {
          headers: {
            'ngrok-skip-browser-warning': 'true',
          },
        });
        const data = await response.json();
        setProduct(data);
        setFormData({ ...data, image: '' });
      } catch (error) {
        console.error('Error fetching product data:', error);
      } finally {
        setLoading(false);
      }
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;

    setFormData({
      ...formData,
      [name]: name === 'price' || name === 'stock' ? parseFloat(value) || 0 : value,
    });
  };

  const handleImageChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      if (file.type === 'image/png') {
        const reader = new FileReader();
        reader.onloadend = () => {
          setFormData({ ...formData, image: reader.result as string });
        };
        reader.readAsDataURL(file);
      } else {
        alert('Only .png files are allowed');
        e.target.value = '';
      }
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const { id, ...updatedFormData } = {
      ...formData,
      price: parseFloat(formData.price as unknown as string) || 0,
      stock: parseFloat(formData.stock as unknown as string) || 0,
    };

    try {
      const response = await fetch(`${API_URL}/api/products/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${Cookies.get('token')}`,
        },
        body: JSON.stringify(updatedFormData),
      });
      if (response.ok) {
        alert('Product updated successfully');
        router.push(`/product/${id}`);
      } else {
        alert('Failed to update product');
      }
    } catch (error) {
      console.error('Error updating product:', error);
      alert('An error occurred while updating the product');
    }
  };

  useEffect(() => {
    checkAdminRole();
    fetchProductData();
  }, [id]);

  if (loading) return <div>Loading...</div>;
  if (!product) return <div>Product not found</div>;

  return (
    <div>
      <Navbar />
      <div className="max-w-7xl mx-auto p-8 text-customBlack">
        <h1 className="text-3xl font-bold mb-4">Edit Produk</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-lg font-semibold">Nama Produk</label>
            <input
              type="text"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              className="border p-2 w-full"
            />
          </div>
          <div>
            <label className="block text-lg font-semibold">Deskripsi Header</label>
            <input
              type="text"
              name="header_description"
              value={formData.header_description}
              onChange={handleInputChange}
              className="border p-2 w-full"
            />
          </div>
          <div>
            <label className="block text-lg font-semibold">Deskripsi</label>
            <textarea
              name="description"
              value={formData.description}
              onChange={handleInputChange}
              className="border p-2 w-full"
            />
          </div>
          <div>
            <label className="block text-lg font-semibold">Harga</label>
            <input
              type="number"
              name="price"
              value={formData.price}
              onChange={handleInputChange}
              className="border p-2 w-full"
            />
          </div>
          <div>
            <label className="block text-lg font-semibold">Stok</label>
            <input
              type="number"
              name="stock"
              value={formData.stock}
              onChange={handleInputChange}
              className="border p-2 w-full"
            />
          </div>
          <div>
            <label className="block text-lg font-semibold">Ubah Gambar Produk (hanya .png)</label>
            <input
              type="file"
              accept=".png"
              onChange={handleImageChange}
              className="border p-2 w-full"
            />
            {formData.image && (
              <div className="mt-4">
                <p className="font-semibold">Preview:</p>
                <img
                  src={formData.image}
                  alt="Product Preview"
                  className="w-64 h-64 object-cover border mt-2"
                />
              </div>
            )}
          </div>
          <div className="flex justify-center space-x-4">
            <button type="submit" className="bg-accent w-24 text-white font-semibold py-2 px-4 rounded">
              Simpan
            </button>
            <button
              type="button"
              onClick={() => router.push('/')}
              className="bg-customRed w-24 text-white font-semibold py-2 px-4 rounded"
            >
              Batal
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AdminProductEdit;
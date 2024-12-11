import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { API_URL } from '@/utils/config';
import Navbar from '@/components/Navbar';
import Cookies from 'js-cookie';
import BuyProductButton from '@/components/BuyProductButton';
import Footer from '@/components/Footer';

interface Product {
    name: string;
    header_description: string;
    description: string;
    price: number;
    id: string;
    sold: number;
    stock: number;
}

const ProductDetail = () => {
    const router = useRouter();
    const { id } = router.query;
    const [product, setProduct] = useState<Product>({
        name: '',
        header_description: '',
        description: '',
        price: 0,
        id: '',
        sold: 0,
        stock: 0,
    });
    const [loading, setLoading] = useState<boolean>(true);
    const [quantity, setQuantity] = useState(1);

    const handleAddToCart = async (q: number) => {
        const token = Cookies.get('token');

        if (!token) {
            router.push('/register');
            return;
        }

        if (!product) return;

        try {
            const productData = {
                id: product.id,
                name: product.name,
                price: product.price,
                quantity: q,
            };

            const response = await fetch(`${API_URL}/api/user/cart/products`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
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

    const handleIncrement = () => {
        setQuantity((prev) => (prev < product.stock ? prev + 1 : prev));
    };

    const handleDecrement = () => {
        setQuantity((prev) => (prev > 1 ? prev - 1 : 1));
    };

    useEffect(() => {
        if (id) {
            const fetchProduct = async () => {
                try {
                    const response = await fetch(`${API_URL}/api/products/${id}`, {
                        headers: {
                            'ngrok-skip-browser-warning': 'true',
                        },
                    });
                    const data = await response.json();
                    setProduct(data);
                } catch (error) {
                    console.error('Error fetching product data:', error);
                } finally {
                    setLoading(false);
                }
            };

            fetchProduct();
        }
    }, [id]);

    if (loading) return <div>Loading...</div>;
    if (!product) return <div>Product not found.</div>;

    return (
        <div>
            <Navbar />
            <div className="max-w-7xl mx-auto p-8">
                <div className="flex flex-col md:flex-row gap-8">
                    <div className="md:w-1/2">
                        <img
                            src={`${API_URL}/api/image/${product.id}`}
                            alt={product.name}
                            className="rounded-lg w-full h-96 md:h-[400px] object-cover"
                        />
                    </div>

                    <div className="md:w-1/2 flex flex-col justify-between text-customBlack">
                        <h1 className="text-4xl font-bold text-customBlack mb-4">{product.name}</h1>
                        <p className="text-md text-customBlack font-medium">{product.sold} terjual</p>
                        <p className="mt-2">{product.description}</p>
                        <h2 className="text-3xl text-customBlack mt-4 font-bold">{`Rp ${product.price.toLocaleString()}`}</h2>
                        <p className="font-medium">Stok: {product.stock}</p>
                        <div className="flex items-center gap-8 mt-2">
                            <button
                                onClick={handleDecrement}
                                className="bg-secondary-light shadow-md text-customBlack font-bold py-2 px-4 rounded"
                            >
                                -
                            </button>
                            <span className="text-xl font-bold">{quantity}</span>
                            <button
                                onClick={handleIncrement}
                                className="bg-secondary-light shadow-md text-customBlack font-bold py-2 px-4 rounded"
                            >
                                +
                            </button>
                        </div>
                        <div className="flex flex-wrap gap-4 items-center mt-4">
                            <button
                                className="px-4 py-2 bg-accent shadow-lg text-white text-md font-semibold rounded-md z-10"
                                onClick={() => handleAddToCart(quantity)}
                            >
                                Tambah ke Keranjang
                            </button>

                            <BuyProductButton
                                id={product.id}
                                name={product.name}
                                price={product.price}
                                quantity={quantity}
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ProductDetail;

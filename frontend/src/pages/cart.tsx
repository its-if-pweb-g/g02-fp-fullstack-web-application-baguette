import Navbar from "@/components/Navbar";
import { API_URL } from '@/utils/config';
import Image from "next/image";
import Cookies from "js-cookie";
import { useEffect, useState } from "react";

interface cartItemData {
    id: string;
    name: string;
    price: number;
    quantity: number;
}

const Cart = () => {
    const [cartItems, setCartItems] = useState<cartItemData[]>([]);
    const [quantity, setQuantity] = useState(1);

    const fetchProduct = async () => {
        try {
            const response = await fetch(`${API_URL}/api/user/cart`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${Cookies.get('token')}`,
                    'ngrok-skip-browser-warning': 'true',
                },
            });

            if (!response.ok) {
                throw new Error('Failed to add product to cart');
            }

            const data = await response.json()
            setCartItems(data)

        } catch (error) {
            alert('Terjadi kesalahan pada server');
        }
    };

    const handleIncrement = () => {
        setQuantity((prev) => (prev + 1));
    };

    const handleDecrement = () => {
        setQuantity((prev) => (prev > 1 ? prev - 1 : 1));
    };

    useEffect(() => {
        fetchProduct();
    }, []);

    return (
        <div>
            <Navbar></Navbar>

            <div className="mx-24 mt-8 ">
                <h1 className="text-4xl font-bold text-customBlack"> Isi Keranjang </h1>

                <div className="flex gap-8">
                    <div className="w-[760px] mr-20">
                        {cartItems && cartItems.length > 0 ? (
                            cartItems.map((item) => (
                                <div className="w-full h-44 bg-white justify-between rounded-lg flex mt-8">
                                    <div className="flex gap-8">
                                        <img src={`${API_URL}/api/image/${item.id}`} alt="" className="rounded-l-md" />
                                        <div className="text-customBlack py-3 flex flex-col justify-between">
                                            <div>
                                                <h1 className="text-2xl font-semibold">{item.name}</h1>
                                                <h1 className="text-2xl font-bold mt-1">Rp {item.price}</h1>
                                            </div>
                                            <div>
                                                <div className="flex items-center gap-8 mt-2">
                                                    <button
                                                        onClick={handleDecrement}
                                                        className="bg-secondary-light text-customBlack font-bold py-2 px-4 rounded"
                                                    >
                                                        -
                                                    </button>
                                                    <span className="text-xl font-bold">{quantity}</span>
                                                    <button
                                                        onClick={handleIncrement}
                                                        className="bg-secondary-light text-customBlack font-bold py-2 px-4 rounded"
                                                    >
                                                        +
                                                    </button>
                                                </div>
                                            </div>
                                        </div>
                                    </div>


                                    <Image src="/bin.png" alt="bin" width={32} height={32} className="object-contain" />
                                </div>

                            ))
                        ) : (
                            <div className="flex items-center justify-center border-l-4 border-red-300 bg-red-50 text-gray-700 px-4 py-2">
                                <p>Isi keranjangmu kosong, mulai belanja!</p>
                            </div>
                        )}
                    </div>

                    <div className="w-[428px] h-fit bg-white px-6 py-6">
                        <div className="mb-10">
                            <h1 className="text-customBlack text-2xl font-bold mb-10">Total Belanja</h1>
                            <h1 className="text-2xl font-bold mt-1 text-customBlack mb-8">Rp 1234567</h1>
                            <h1 className="text-customBlack text-xl font-semibold mb-6">Powered by Midtrasn</h1>
                            <button
                                type="submit"
                                className="w-24 font-medium border-2 border-accent bg-accent px-4 py-2 rounded-md text-primary  hover:opacity-80 transition"
                                >
                                Beli
                            </button>
                        </div>
                        <p className="text-customBlack font-semibold py-2">Alamat: Jalan Shiganshina No. 3 Surabaya</p>
                        <button
                            className="border-customBlack text-customBlack font-bold  px-4 rounded-md border-2 bg-white py-2 mt-6"
                        >
                            Ganti Alamat
                        </button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Cart;
import Navbar from "@/components/Navbar";
import { API_URL } from "@/utils/config";
import Image from "next/image";
import Cookies from "js-cookie";
import { useEffect, useState } from "react";
import BuyProductsCartButton from "@/components/BuyProductsCartButton";
import CartItem from '@/components/CartItem';
import Link from "next/link";


interface cartItemData {
    id: string;
    name: string;
    price: number;
    quantity: number;
}

const Cart = () => {
    const [cartItems, setCartItems] = useState<cartItemData[]>([]);
    
    const calculateTotalPrice = () => {
        return cartItems.reduce((total, item) => total + item.price * item.quantity, 0);
    };

    const updateQuantity = (id: string, quantity: number) => {
        setCartItems((prevItems) =>
            prevItems.map((item) =>
                item.id === id ? { ...item, quantity: quantity } : item
            )
        );
    };

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
            console.log(data)
            setCartItems(data)

        } catch (error) {
        }
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
                                <CartItem key={item.id} id={item.id} name={item.name} price={item.price} quantity={item.quantity} updateQuantity={updateQuantity}/>
                            ))
                        ) : (
                            <div className="flex items-center mt-4 justify-center rounded-xl border-l-4 border-red-300 bg-red-50 text-gray-700 px-4 py-2">
                                <p>Isi keranjangmu kosong, mulai belanja!</p>
                            </div>
                        )}
                    </div>

                    <div className="w-[428px] h-fit bg-white px-6 py-6">
                        <div className="mb-10">
                            <h1 className="text-customBlack text-2xl font-bold mb-10">Total Belanja</h1>
                            <h1 className="text-2xl font-bold mt-1 text-customBlack mb-8">Rp 
                                {calculateTotalPrice().toLocaleString('id-ID')}
                            </h1>
                            <h1 className="text-customBlack text-xl font-semibold mb-6">Powered by Midtrasn</h1>
                            <BuyProductsCartButton /> 
                        </div>
                        <p className="text-customBlack font-semibold py-2"></p>
                        <button
                            className="border-customBlack text-customBlack font-bold  px-4 rounded-md border-2 bg-white py-2 mt-6"
                        >
                            <Link href="/ubahData">Ganti Alamat</Link>
                        </button>
                    </div>
                </div>
            </div>
        </div>


    )
}

export default Cart;
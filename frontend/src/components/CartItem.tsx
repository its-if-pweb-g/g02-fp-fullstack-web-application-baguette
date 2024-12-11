import Image from "next/image";
import { API_URL } from "@/utils/config";
import { useState } from "react";
import Cookies from "js-cookie";


interface items {
    id: string;
    name: string;
    price: number;
    quantity: number;
    updateQuantity: (id: string, quantity: number) => void;
}


const BuyProductsCartButton: React.FC<items> = ({ id, name, price, quantity, updateQuantity}) => {
    const [q, setQuantity] = useState(quantity);

    const handlerIncProduct = async () => {
        try {
            setQuantity(q + 1);
            updateQuantity(id, q + 1);

            const response = await fetch(`${API_URL}/api/user/cart/products/inc/${id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${Cookies.get("token")}`,
                    "ngrok-skip-browser-warning": "true",
                },
                body: JSON.stringify({ quantity: 1 }),
            });

            if (!response.ok) {
                throw new Error("Failed to increment product quantity");
            }
        } catch (error) {
            alert("Terjadi kesalahan pada server");
        }

    };

    const handlerDecProduct = async () => {
        if (q <= 1) return;
        setQuantity(q - 1);
        updateQuantity(id, q - 1);
        try {

            const response = await fetch(`${API_URL}/api/user/cart/products/dec/${id}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${Cookies.get("token")}`,
                    "ngrok-skip-browser-warning": "true",
                },
            });

            if (!response.ok) {
                throw new Error("Failed to decrement product quantity");
            }
        } catch (error) {
            alert("Terjadi kesalahan pada server");
        }

    };

    const deleteProductFromCart = async () => {
        try {
            const response = await fetch(`${API_URL}/api/user/cart/products/${id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${Cookies.get("token")}`,
                },
            });

            if (!response.ok) {
                throw new Error(`Error: ${response.status} - ${response.statusText}`);
            }

            alert("Berhasil menghapus produk");
        } catch (error) {
            alert("Gagal menghapus produk");
        }
        window.location.reload()
    };

    return (
        <div className="w-full h-44 bg-white justify-between rounded-lg flex mt-8">
            <div className="flex gap-8">
                <img
                    src={`${API_URL}/api/image/${id}`}
                    alt=""
                    className="rounded-l-md"
                />
                <div className="text-customBlack py-3 flex flex-col justify-between">
                    <div>
                        <h1 className="text-2xl font-semibold">{name}</h1>
                        <h1 className="text-2xl font-bold mt-1">Rp {price}</h1>
                    </div>
                    <div>
                        <div className="flex items-center gap-8 mt-2">
                            <button
                                onClick={handlerDecProduct}
                                className="bg-secondary-light text-customBlack font-bold py-2 px-4 rounded"
                            >
                                -
                            </button>
                            <span className="text-xl font-bold">{q}</span>
                            <button
                                onClick={handlerIncProduct}
                                className="bg-secondary-light text-customBlack font-bold py-2 px-4 rounded"
                            >
                                +
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <Image
                src="/bin.png"
                alt="bin"
                width={32}
                height={32}
                className="object-contain"
                onClick={deleteProductFromCart}
            />
        </div>
    );
};

export default BuyProductsCartButton;

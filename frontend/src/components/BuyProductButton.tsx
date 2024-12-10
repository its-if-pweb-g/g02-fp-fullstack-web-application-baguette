import getCookie from "@/utils/utility";
import { API_URL } from "@/utils/config";

interface ProductPayment {
    id: string;
    name: string;
    quantity: number;
    price: number;
}

const BuyProductButton: React.FC<ProductPayment> = ({ name, id, quantity, price }) => {
    const handleBuyNow = async () => {
        try {
            const response = await fetch(`${API_URL}/api/user/pay`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + getCookie("token"),
                },
                body: JSON.stringify({
                    product_id: id,
                    name: name,
                    price: price,
                    quantity: quantity,
                }),
            });

            if (!response.ok) {
                throw new Error("Failed to make a purchase.");
            }

            const data = await response.json();
            window.location.href = data.redirect_url
        } catch (error) {
            console.error(error);
            alert("Purchase failed. Please try again.");
        }
    };

   
    return (
        <button className="mt-4 px-6 py-2 w-52 font-semibold bg-secondary-dark text-white rounded-lg"
            onClick={handleBuyNow}
        >
            Beli Langsung
        </button>
    );
};

export default BuyProductButton;
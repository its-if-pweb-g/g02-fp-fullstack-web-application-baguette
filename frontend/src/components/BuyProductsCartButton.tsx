import { API_URL } from "@/utils/config";
import getCookie from "@/utils/utility";


const BuyProductsCartButton = () => {
    const handleBuyNow = async () => {
        try {
            const response = await fetch(`${API_URL}/api/user/cart/pay`, {
                headers: {
                    Authorization: "Bearer " + getCookie("token"),
                },
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
        <button className="mt-4 px-6 py-2 w-[103px] font-semibold bg-secondary-dark text-white rounded-lg"
            onClick={handleBuyNow}
        >
            Beli Langsung
        </button>
    );
};

export default BuyProductsCartButton;
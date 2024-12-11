import { useEffect, useState } from "react";
import Navbar from "@/components/Navbar";
import { API_URL } from "@/utils/config";
import Cookies from "js-cookie";
import { useRouter } from "next/router";

interface RegisterData {
  name: string;
  email: string;
  phone: string;
  address: string;
}

export default function Home() {
  const [userData, setUserData] = useState<RegisterData | null>(null);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`${API_URL}/api/user`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + Cookies.get('token'),
            'ngrok-skip-browser-warning': 'true',
          },
        });

        if (!response.ok) {
          throw new Error('Failed to fetch user data');
        }

        const data: RegisterData = await response.json();
        const filteredData: RegisterData = {
          name: data.name,
          email: data.email,
          phone: data.phone,
          address: data.address
        };

        setUserData(filteredData);
      } catch (error) {
        console.error(error);
        setError('Failed to load user data');
      }
    };

    fetchData();
  }, []);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (userData) {
      const { id, value } = e.target;
      setUserData({ ...userData, [id]: value });
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    console.log(userData)
    if (userData) {
      try {
        const response = await fetch(`${API_URL}/api/user`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + Cookies.get('token'),
            'ngrok-skip-browser-warning': 'true',
          },
          body: JSON.stringify(userData),
        });
    
        alert('Data berhasil diperbarui');
      } catch (error) {
        alert('Terjadi kesalahan saat memperbarui data');
      }
    }
  };

  const handleLogout = () => {
    Cookies.remove('token'); 
    router.push('/'); 
  };

  return (
    <div className="bg-primary h-lvh mb-3">
      <Navbar />
      
      <div className="items-center flex flex-col">
        <div className="w-[764px] h-[340px] bg-primary px-6 py-6">
          <h1 className="text-2xl font-bold text-customBlack mb-6">Ubah Data Diri</h1>
          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label htmlFor="nama" className="block text-customBlack font-bold mb-2">Nama Lengkap</label>
              <input
                type="text"
                id="name"
                value={userData ? userData.name : ''}
                onChange={handleInputChange}
                className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
              />
            </div>
            <div className="mb-4">
              <label htmlFor="alamat" className="block text-customBlack font-bold mb-2">Alamat</label>
              <input
                type="text"
                id="address"
                value={userData ? userData.address : ''}
                onChange={handleInputChange}
                className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
              />
            </div>
            <div className="mb-4">
              <label htmlFor="email" className="block text-customBlack font-bold mb-2">Email</label>
              <input
                type="email"
                id="email"
                value={userData ? userData.email : ''}
                onChange={handleInputChange}
                className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
              />
            </div>
            <div className="mb-4">
              <label htmlFor="telepon" className="block text-customBlack font-bold mb-2">Nomor Telepon</label>
              <input
                type="number"
                id="phone"
                value={userData ? userData.phone : ''}
                onChange={handleInputChange}
                className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
              />
            </div>
            <div className="flex items-center justify-center">
              <button type="submit" className="w-24 font-medium border-2 border-accent bg-accent px-4 py-2 rounded-md text-primary transition mx-4">Kirim</button>
              <button type="button" onClick={() => router.push("/")} className="w-24 font-medium border-2 border-customRed bg-customRed px-4 py-2 rounded-md text-primary transition mx-4">Batal</button>
            </div>
          </form>

          <h1 className="mt-8 text-2xl font-bold text-customBlack mb-4">Keluar Akun</h1>
          <button onClick={handleLogout} className="w-24 font-medium border-2 border-customRed bg-customRed px-4 py-2 rounded-md text-primary transition">Log out</button>
        </div>
      </div>
    </div>
  );
}

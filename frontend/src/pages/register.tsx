import { useState, useEffect } from "react";
import Image from "next/image";
import Link from "next/link";
import Cookies from "js-cookie";
import { useRouter } from "next/router"; 
import { API_URL } from "@/utils/config";

const API_BASE_URL = API_URL;

if (!API_BASE_URL) {
  throw new Error("API_BASE_URL is not defined in the environment variables");
}

export default function Home() {
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [phone, setPhone] = useState("");
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const userData = {
      name,
      email,
      password,
      phone
    };

    try {
      const response = await fetch(`${API_BASE_URL}/api/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      });

      if (!response.ok) {
        throw new Error("Failed to register user");
      }
      const data = await response.json();
      Cookies.set("token", data.token, { expires: 7 }); 
      router.push("/")

    } catch (error) {
      console.error("Registration failed", error);
    }
  };
  return (
    <div className="bg-primary h-lvh">
      {/* Navbar */}
      <div className="h-16 md:h-20 px-4 md:px-8 lg:px-16 xl:px-32 2xl:px-64 bg-secondary-dark">
        <div className="hidden md:flex items-center justify-center gap-16 h-full">
          <div>
            <Link href="/" className="flex items-center gap-4">
              <Image src="/logo.png" alt="" width={28} height={28} />
              <div className="text-2xl text-primary">Baguette</div>
            </Link>
          </div>
        </div>
      </div>

      {/* Registration Form */}
      <div className="items-center flex flex-col">
        <div className="w-[764px] h-[340px] bg-primary px-6 py-6">
          <h1 className="text-2xl font-bold text-customBlack mb-6">Daftar</h1>
          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label
                htmlFor="email"
                className="block text-customBlack font-bold mb-2"
              >
                Email
              </label>
              <input
                type="email"
                id="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="Eren@gmail.com"
                className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
              />
            </div>
            <div className="mb-4">
              <label
                htmlFor="name"
                className="block text-customBlack font-bold mb-2"
              >
                Nama Lengkap
              </label>
              <input
                type="text"
                id="name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Eren"
                className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
              />
            </div>
            <div className="mb-6">
              <label
                htmlFor="password"
                className="block text-customBlack font-bold mb-2"
              >
                Kata Sandi
              </label>
              <input
                type="password"
                id="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="**********"
                className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
              />
            </div>
            <div className="mb-4">
                <label
                  htmlFor="phone"
                  className="block text-customBlack font-bold mb-2"
                >
                  Nomor Telepon
                </label>
                <input
                  type="string"
                  id="phone"
                  value={phone}
                  onChange={(e) => setPhone(e.target.value)}
                  placeholder="081234567890"
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
                />
              </div>
            <div className="flex items-center justify-between">
              <button
                type="submit"
                className="w-24 font-medium border-2 border-accent bg-accent px-4 py-2 rounded-md text-primary hover:opacity-80 transition"
              >
                Kirim
              </button>
              <p className="text-sm text-gray-600 ml-4">
                  <Link href="/login" className="text-secondary-dark underline">
                  Sudah punya akun?{" "}Masuk di sini.
                  </Link>
              </p>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}

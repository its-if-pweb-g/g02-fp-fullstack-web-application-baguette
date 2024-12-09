import Image from "next/image";
import localFont from "next/font/local";
import Navbar from "../components/Navbar";
import NavIcons from "../components/NavIcons";
import Link from "next/link";
import Menu from "../components/Menu";

export default function Home() {
    return (
      <div className="bg-primary h-lvh">
        {/* {Navbar} */}
        <div className="h-16 md:h-20 px-4 md:px-8 lg:px-16 xl:px-32 2xl:px-64 bg-secondary-dark">
            <div className="hidden md:flex items-center justify-center gap-16 h-full">
                <div>
                    <Link href="/" className="flex items-center gap-4">
                    <Image src="/logo.png" alt="" width={28} height={28} />
                    <div className="text-2xl">Baguette</div>
                    </Link>
                </div>
            </div>
        </div>

        {/*  */}
        <div className="items-center flex flex-col">
          <div className="w-[764px] h-[340px] bg-primary px-6 py-6">
            <h1 className="text-2xl font-bold text-customBlack mb-6">Masuk</h1>
            <form>
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
                  placeholder="Eren@gmail.com"
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166]"
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
                  placeholder="**********"
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166]"
                />
              </div>
              <div className="flex items-center justify-between">
                <button
                  type="submit"
                  className="w-24 font-medium border-2 border-accent bg-accent px-4 py-2 rounded-md text-primary  transition"
                >
                  Kirim
                </button>
                <p className="text-sm text-gray-600 ml-4">
                  <Link href="/register" className="text-secondary-dark underline">
                  Tidak punya akun?{" "}Daftar sekarang.
                  </Link>
                </p>
              </div>
            </form>
          </div>
        </div>
      </div>
    );
  }
  
import Footer from "@/components/Footer";
import Navbar from "@/components/Navbar";
import Image from "next/image";
import Link from "next/link";


export default function Home() {
    return (
      <div className="bg-primary h-lvh">
        {/* {Navbar} */}
        <Navbar></Navbar>

        {/*  */}
        <div className="items-center flex flex-col">
          <div className="w-[764px] h-[340px] bg-primary px-6 py-6">
            <h1 className="text-2xl font-bold text-customBlack mb-6">Ubah Data Diri</h1>
            <form>
              <div className="mb-4">
                <label
                  htmlFor="email"
                  className="block text-customBlack font-bold mb-2"
                >
                  Nama Lengkap
                </label>
                <input
                  type="text"
                  id="Nama"
                  placeholder="Eren"
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="tanggalLahir"
                  className="block text-customBlack font-bold mb-2"
                >
                  Tanggal Lahir
                </label>
                <input
                  type="date"
                  id="Tanggal Lahir"
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="jenisKelamin"
                  className="block text-customBlack font-bold mb-2"
                >
                  Jenis Kelamin
                </label>
                <input
                  type="text"
                  id="jenisKelamin"
                  placeholder=""
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="email"
                  className="block text-customBlack font-bold mb-2"
                >
                  Alamat
                </label>
                <input
                  type="text"
                  id="Nama"
                  placeholder="Perumdos"
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
                />
              </div>
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
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="telepon"
                  className="block text-customBlack font-bold mb-2"
                >
                  Nomor Telepon
                </label>
                <input
                  type="number"
                  id="telepon"
                  placeholder="081234567890"
                  className="w-full px-4 py-2 border border-customBlack rounded focus:outline-none focus:ring-2 focus:ring-[#FFD166] text-customBlack"
                />
              </div>
              <div className="flex items-center justify-center">
                <button
                  type="submit"
                  className="w-24 font-medium border-2 border-accent bg-accent px-4 py-2 rounded-md text-primary  transition mx-4"
                >
                  Kirim
                </button>
                <button
                  type="submit"
                  className="w-24 font-medium border-2 border-customRed bg-customRed px-4 py-2 rounded-md text-primary  transition mx-4"
                >
                  Batal
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    );
  }
  
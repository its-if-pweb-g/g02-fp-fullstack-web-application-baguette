import Link from "next/link";
import Menu from "./Menu";
import Image from "next/image";
import SearchBar from "./SearchBar";
import NavIcons from "./NavIcons";

const Navbar = () => {
    return (
      <div className="h-16 md:h-20 px-4 md:px-8 lg:px-16 xl:px-32 2xl:px-64 bg-secondary-dark">
        <div className="h-full flex items-center justify-between gap-4 relative md:hidden">
            <Link href="/">
                <div className="text-lg md:text-xl ">Baguette</div>
            </Link>
            <SearchBar/>
            <Menu/>
        </div>

        <div className="hidden md:flex items-center justify-between gap-16 h-full">
            <div className="">
                <Link href="/" className="flex items-center gap-4">
                    <Image src="/logo.png" alt="" width={28} height={28}/>
                    <div className="text-2xl text-primary">Baguette</div>
                </Link>
            </div>
            <div className="flex-1 mx-8">
                <SearchBar />
            </div>
            <div className="">
                <NavIcons/>
            </div>
        </div>
      </div>
    );
  };
  
  export default Navbar;
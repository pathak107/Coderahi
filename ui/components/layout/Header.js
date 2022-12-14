import Link from "next/link";
import ThemeToggle from "../ThemeToggle";
import Image from "next/image";
import Logo from '../../public/coderahi_logo.png'

const Header = () => {
    return (
        <>
            <div className="navbar bg-base-100 border-b-2 border-sky-400 h-2">
                <div className="flex-1 gap-1">
                    <div className="avatar">
                        <div className="w-12 rounded-full">
                            <Image
                                src={Logo}
                                alt="Coderahi logo"
                            />
                        </div>
                    </div>
                    <a className="btn btn-ghost normal-case text-xl">Coderahi</a>
                </div>
                <div className="flex-none">
                    <ul className="menu menu-horizontal p-0">
                        <Link href="/"><li><a>Home</a></li></Link>
                        <Link href="/courses"><li><a>Courses</a></li></Link>
                        <li><a href="https://blog.coderahi.in">Blog</a></li>
                        <Link href="/about"><li><a>About</a></li></Link>
                        <ThemeToggle />
                    </ul>
                </div>
            </div>
        </>
    );
}

export default Header;
import Link from "next/link";
import { FaGithub, FaLinkedin, FaTwitter } from "react-icons/fa";

const Footer = () => {
    return (
        <footer className="footer footer-center p-10 bg-base-200 text-base-content rounded">
            <div className="grid grid-flow-col gap-4">
                <Link href="/"><a className="link link-hover">Home</a></Link>
                <Link href="/courses"><a className="link link-hover">Courses</a></Link>
                <Link href="https://blog.coderahi.in"><a className="link link-hover">Blog</a></Link>
                <Link href="/about"><a className="link link-hover">About</a></Link>

            </div>
            <div>
                <div className="grid grid-flow-col gap-4">
                    <a href="https://www.linkedin.com/in/shubham-pathak107/"><FaLinkedin/></a>
                    <a href="https://github.com/pathak107"><FaGithub /></a>
                    <a href="https://twitter.com/ShubhamPathk107"><FaTwitter/></a> 
                </div>
            </div>
            <div>
                <p>Copyright © 2022 - All right reserved by ACME Industries Ltd</p>
            </div>
        </footer>
    );
}

export default Footer;
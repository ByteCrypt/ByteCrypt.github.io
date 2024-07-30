import User from "./User";
import ByteCrypt_Logo from "../images/ByteCrypt_Logo.svg";
import Image from "next/image";
import Link from "next/link";

export default function Header() {
    const buttonClick = () => {
        alert("You did a thing");
    };

    return (
        <>
            <Link className="link flex" href="/">
                <Image className="h-[30px] w-auto" src={ByteCrypt_Logo} alt="Bytecrypt Logo" />

                <p className="text-2xl italic font-extrabold">
                    <em>BYTECRYPT</em>
                </p>
            </Link>
            <div className="flex text-sm gap-2 items-center">
                <Link className="link" href="/about">About Us</Link>
                <Link className="link" href="/info">Info</Link>
                <button className="h-7 w-24 rounded-3xl bg-violet-700 hover:bg-blue-500 align-top" onClick={buttonClick}>Get Started</button>
                <User />
            </div>
        </>
    );
}

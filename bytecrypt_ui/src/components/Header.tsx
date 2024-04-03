
import ByteCrypt_Logo from "../images/ByteCrypt_Logo.svg";

export default function Header() {
    const buttonClick = () => {
        alert("You did a thing");
    };

    return (
        <>
            <a className="flex" href="#">
                <img className="h-[30px] w-auto" src={ByteCrypt_Logo} alt="Bytecrypt Logo" />

                <p className="text-2xl italic font-extrabold">
                    <em>BYTECRYPT</em>
                </p>
            </a>
            <div className="flex text-sm gap-2 items-center">
                <a href="#">About Us</a>
                <a href="#">Info</a>
                <button className="h-7 w-24 rounded-3xl bg-blue-500 align-top" onClick={buttonClick}>Get Started</button>
            </div>
        </>
    );
}

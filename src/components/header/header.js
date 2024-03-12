import "./header.css";

import ByteCrypt_Logo from "../../images/ByteCrypt_Logo.svg";

export default function Header() {

    const buttonClick = () => {
        alert("You did a thing");
    };

    return (
        <div class="header">
            <div class="header-logo">
                <a class="header-logo-main" href="#">
                    <img class="header-logo-image" src={ByteCrypt_Logo} alt="Bytecrypt Logo" />

                    <p class="header-logo-text">
                        <strong><em>BYTECRYPT</em></strong>
                    </p>
                </a>
            </div>
            <div class="nav">
                <a class="nav-link" href="#">About Us</a>
                <a class="nav-link" href="#">Info</a>
                <button class="nav-button" onClick={buttonClick}>Get Started</button>
            </div>
        </div>
    );
}
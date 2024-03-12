import "./footer.css";

import Twitter_Icon from "../../images/logo-white.png";
import YouTube_Icon from "../../images/yt_logo_rgb_dark.png";

export default function Footer() {
    return (
        <div class="footer">
            <div class="footer-top">
                <div class="footer-about">
                    <p class="footer-about-main">About Us:</p>
                    <a class="footer-link" href="#">Help</a>
                    <a class="footer-link" href="#">Contact Us</a>
                    <a class="footer-link" href="#">User Agreement</a>
                    <a class="footer-link" href="#">Privacy Policy</a>
                </div>

                <div class="footer-social">
                    <a class="twitter-link" href="https://twitter.com/thebytecrypt">
                        <img class="twitter-icon" src={Twitter_Icon}></img>
                    </a>
                    <a class="youtube-link" href="https://thebytecrypt.com">
                        <img class="youtube-icon" src={YouTube_Icon}></img>
                    </a>
                </div>
            </div>

            <div class="footer-bottom">
                <div class="footer-quote">
                    <p class="footer-quote-main">ByteCrypt</p>
                    <p class="footer-quote-text">Take your data to the grave.</p>
                </div>
                <p class="footer-copyright">©2023-2024 Mason Software</p>
            </div>
        </div>
    );
}
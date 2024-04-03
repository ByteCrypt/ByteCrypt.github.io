import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXTwitter } from "@fortawesome/free-brands-svg-icons";
import { faYoutube } from "@fortawesome/free-brands-svg-icons";
import { faGithub } from "@fortawesome/free-brands-svg-icons";

export default function Footer() {

    return (
        <>
            <div className="flex flex-col">
                <p className="text-base font-bold">About Us:</p>
                <div className="text-sm ml-2 flex flex-col">
                    <a href="#">Help</a>
                    <a href="#">Contact Us</a>
                    <a href="#">User Agreement</a>
                    <a href="#">Privacy Policy</a>
                </div>
            </div>

            <div className="self-end">
                <p className="text-lg text-center font-bold">ByteCrypt</p>
                <p className="text-sm uppercase">Take your data to the grave</p>
            </div>

            <div className="flex gap-2">
                <a href="https://github.com/ExtremelyRyan/ByteCrypt">
                    <FontAwesomeIcon icon={faGithub} />
                </a>
                <a href="https://twitter.com/thebytecrypt">
                    <FontAwesomeIcon icon={faXTwitter} />
                </a>
                <a href="https://thebytecrypt.com">
                    <FontAwesomeIcon className="text-red-500/70" icon={faYoutube} />
                </a>
            </div>
        </>
    );
}

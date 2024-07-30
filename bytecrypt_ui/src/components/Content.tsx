import { useState } from "react";
import ByteCryptLogo from "../images/ByteCrypt_Logo.svg";
import Image from "next/image";
import { BackendAddress, Path } from "@/common/path";
import { HttpHeader, Method } from "@/common/utils";

export default function Content() {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");

    return (
        <>
            <div className="flex flex-col justify-center gap-1">
                <h1 className="max-w-96 text-5xl font-bold text-center sm:text-left">
                    Never worry about data leaks again.
                </h1>

                <p className="my-5 text-base text-center sm:text-left">
                    Launching soon.  Sign up for a sneak peek and more.
                </p>

                <form className="flex flex-col sm:self-start self-center">
                    <div className="flex gap-1 flex-col sm:flex-row">
                        <input
                            className="bg-black w-48 h-7 text-sm text-zinc-400 border-slate-800 border-2"
                            type="Text"
                            value={name}
                            placeholder=" First Name"
                            onChange={event => setName(event.target.value)}
                        ></input>

                        <input
                            className="bg-black w-48 h-7 text-sm text-zinc-400 border-slate-800 border-2"
                            type="Text"
                            value={email}
                            onChange={event => setEmail(event.target.value)}
                            placeholder=" you@mail.com"
                        ></input>
                    </div>

                    <input
                        className="self-center sm:self-start hover:bg-blue-500 bg-violet-700 rounded-3xl text-lg mt-3 h-8 w-24 cursor-pointer"
                        type="button"
                        value="Subscribe"
                        onClick={_ => subscribe(name, email)}
                    ></input>
                </form>
            </div>

            <Image className="h-96 w-auto" src={ByteCryptLogo} alt="ByteCrypt Logo" />
        </>
    );
}

function subscribe(name: string, email: string) {
    // https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html
    // Regex for sanitizing the email
    const sanitizeEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (sanitizeEmail.test(email)) {
        fetch(`${BackendAddress}${Path.Subscribe}`, {
            method: Method.Post,
            headers: {
                [HttpHeader.ContentType]: HttpHeader.ApplicationJson,
            },
            body: JSON.stringify({ name: name, email: email }),
        }).then(response => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            return response.json();
        }).then(data => console.log("Success:", data))
            .catch((error) => console.error("Error:", error));
    }
}

function unsubscribe(email: string) {
    // https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html
    // Regex for sanitizing the email
    const sanitizeEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (sanitizeEmail.test(email)) {
        fetch(`${BackendAddress}${Path.Unsubscribe}`, {
            method: Method.Post,
            headers: {
                [HttpHeader.ContentType]: HttpHeader.ApplicationJson,
            },
            body: JSON.stringify({ email: email }),
        }).then(response => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            return response.json();
        }).then(data => console.log("Success:", data))
            .catch((error) => console.error("Error:", error));
    }
}
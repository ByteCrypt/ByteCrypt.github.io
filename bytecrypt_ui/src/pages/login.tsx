import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { useRef, useState } from "react";
import { BackendAddress, Path } from "@/common/path";
import { HttpHeader, Method } from "@/common/utils";

export default function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [usernameBorder, setUsernameBorder] = useState("border-slate-800")
    const [passwordBorder, setPasswordBorder] = useState("border-slate-800")
    const [errorMessage, setErrorMessage] = useState([""]);
    const timer = useRef<NodeJS.Timeout | null>(null);

    const setError = () => {
        let errorMessage: string[] = [];

        // Username validation
        if (username === "") {
            errorMessage.push("Username cannot be blank.");
            setUsernameBorder("border-red-400");
        } else {
            setUsernameBorder("border-slate-800");
        }

        // Password validation
        if (password === "") {
            errorMessage.push("Password cannot be blank.");
            setPasswordBorder("border-red-400");
        } else {
            setPasswordBorder("border-slate-800");
        }

        // Display error message
        if (errorMessage.length !== 0) {
            setErrorMessage(errorMessage);
        } else {
            setErrorMessage([""]);
        }

        resetToDefaults();
    }

    const resetToDefaults = () => {
        if (timer.current) {
            clearTimeout(timer.current);
        }

        timer.current = setTimeout(() => {
            setErrorMessage([""]);
            setUsernameBorder("border-slate-800");
            setPasswordBorder("border-slate-800");
        }, 5000);
    }

    return (
        <>
            <div className="text-white fixed top-0 w-full flex flex-col">
                <div className="my-4 header flex sm:flex-row flex-col sm:justify-between justify-center mx-auto sm:mx-10">
                    <Header />
                </div>
                <hr />
            </div>

            <div className="App flex flex-col justify-between text-3xl text-white h-screen pt-24">
                <div className="mx-auto my-auto flex flex-col justify-around w-11/12 sm:w-1/2 h-1/3 border-2 rounded-3xl max-w-96">
                    <p className="text-center font-bold">Login</p>
                    <div className="Error text-red-400 text-base text-wrap self-center flex flex-col max-h-7 -mt-2">
                        {errorMessage.map((message, i) => (
                            <p>{message}</p>
                        ))}
                    </div>
                    <form className="flex flex-col self-center font-normal text-base">
                        <label>Username:</label>
                        <input
                            className={`bg-black w-48 h-7 text-sm text-zinc-400 ${usernameBorder} border-2`}
                            type="text"
                            value={username}
                            onChange={event => setUsername(event.target.value)}
                            placeholder=" Username">
                        </input>
                        <label>Password:</label>
                        <input
                            className={`bg-black w-48 h-7 text-sm text-zinc-400 ${passwordBorder} border-2`}
                            type="text"
                            value={password}
                            onChange={event => setPassword(event.target.value)}
                            placeholder=" Password">
                        </input>

                        <div className="flex gap-2 self-center">
                            <input
                                className="self-center  hover:bg-blue-500 bg-violet-700 rounded-3xl text-lg mt-3 h-8 w-20 cursor-pointer"
                                type="button"
                                value="Login"
                                onClick={_ => attemptLogin(setError(), username, password)}
                            ></input>
                            <input
                                className="self-center  hover:bg-blue-500 bg-violet-700 rounded-3xl text-lg mt-3 h-8 w-20 cursor-pointer"
                                type="button"
                                value="Register"
                                onClick={_ => attemptRegister(setError(), username, password)}
                            ></input>
                        </div>
                    </form>
                </div>

                <div className="mt-auto">
                    <hr />
                    <div className="footer my-4 mx-10 flex sm:flex-row flex-col sm:justify-between">
                        <Footer />
                    </div>
                </div>
            </div>
        </>
    );
}

function attemptLogin(setError: void, username: string, password: string) {
    setError;
    if (username === "" || password === "") {
        return;
    }

    fetch(`${BackendAddress}${Path.Login}`, {
        method: Method.Post,
        headers: {
            [HttpHeader.ContentType]: HttpHeader.ApplicationJson,
        },
        body: JSON.stringify({ username: username, password: password }),
    }).then(response => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.json();
    }).then(data => console.log("Success:", data))
        .catch((error) => console.error("Error:", error));
}

function attemptRegister(setError: void, username: string, password: string) {
    setError;
    if (username === "" || password === "") {
        return;
    }

    fetch(`${BackendAddress}${Path.Register}`, {
        method: Method.Post,
        headers: {
            [HttpHeader.ContentType]: HttpHeader.ApplicationJson,
        },
        body: JSON.stringify({ username: username, password: password }),
    }).then(response => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.json();
    }).then(data => console.log("Success:", data))
        .catch((error) => console.error("Error:", error));
}
import Link from "next/link";
import { useState } from "react";

export class UserInfo {
    Name: string | undefined;
    Token: string | undefined;
    constructor(name: string | null = null, token: string | null = null) {
        this.Name = name ?? undefined;
        this.Token = token ?? undefined;
    }

    Display = () => {
        return (
            <div className="flex">
                <p>{this.Name}</p>
                {/** Display picture goes here */}
            </div>
        );
    };
}

export default function User({ Name = null, Token = null }) {
    const [userInfo, setUser] = useState(new UserInfo(Name, Token));

    return (
        <div className="h-auto w-36 border-2 hover:bg-slate-600 border-gray-200 py-2 px-2 rounded-lg text-center font-bold link">
                { userInfo.Name === undefined ? (<Link href="/login">Login</Link>) : userInfo.Name }
        </div>
    );
}
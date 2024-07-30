import Link from "next/link";
import { useState } from "react";

export class UserInfo {
    Name: string | undefined;
    Token: string | undefined;
    constructor(name: string | null = null, token: string | null = null) {
        this.Name = name ?? undefined;
        this.Token = token ?? undefined;
    }
}

export default function User({ Name = null, Token = null }) {
    const [userInfo, setUser] = useState(new UserInfo(Name, Token));

    return (
        <div className="h-max w-auto border-3 py-2 px-2 rounded-lg">
                { userInfo.Name === undefined ? (<Link className="link" href="/login" >Login</Link>) : userInfo.Name }
        </div>
    );
}
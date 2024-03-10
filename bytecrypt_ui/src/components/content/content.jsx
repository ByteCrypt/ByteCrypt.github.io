import React, { useState } from "react"
import "./content.css";

export default function Content() {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    // Regex for sanitizing the email
    const sanitizeEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    const subscribe = () => {
        // Will need to ensure the input is validated and sanitized
        // https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html

        if (sanitizeEmail.test(email)) {
            fetch('http://localhost:5150/api/subscribe', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email: email }),
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => console.log('Success:', data))
            .catch((error) => console.error('Error:', error));
        }
    };

    return (
        <div class="subscribe" >
            <h1 class="subscribe-main" style={{ maxWidth: "8em"}}>
                Never worry about data leaks again.
            </h1>

            <p class="subscribe-text" style={{ fontSize: "0.7em"}}>
                Launching soon.  Sign up for a sneak peek and more.
            </p>

            <form class="subscribe-form">
                <div class="subscribe-inputs">
                    <input
                        class="subscribe-name-input"
                        type="Text"
                        value={name}
                        placeholder=" First Name"
                        onChange={event => setName(event.target.value)}
                    ></input>

                    <input
                        class="subscribe-email-input"
                        type="Text"
                        value={email}
                        onChange={event => setEmail(event.target.value)}
                        placeholder=" you@mail.com"
                    ></input>
                </div>

                <input
                    class="subscribe-button"
                    type="button"
                    value="Subscribe"
                    onClick={subscribe}
                ></input>
            </form>
        </div>
    );
}

function fetchData() {
    fetch('http://localhost:5150/api/data')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => console.log(data))
      .catch(error => console.error('Error fetching data:', error));
  }
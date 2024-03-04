import './App.css';
import React, { useState } from 'react';

function SubscribeForm() {
  const [email, setEmail] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();  // Prevent default form submission
    subscribe(email);    // Call the subscribe function
  }

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Email:
        <input
          type="email"
          value = {email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
      </label>
      <button type="submit">Subscribe</button>
    </form>
  )
}

export default function Bytecrypt() {
  return (
    <div>
      <h1>Bytecrypt</h1>
      <SubscribeForm />
    </div>
  )
}

function AboutPage() {

}

function InfoPage() {

}

function HomePage() {

}

function subscribe(email) {
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
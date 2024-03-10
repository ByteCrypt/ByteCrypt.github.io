import React, { useState } from 'react';
import './App.css';

import Header from './components/header/header';
import Content from './components/content/content';
import Footer from './components/footer/footer';

export default function App() {

  return (
    <div class="App"
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between"
      }}>

      <div class="app-header"
        style={{ 
          margin: "2em 4em" 
      }}>
        <Header></Header>
      </div>
      
      <div class="app-subscribe"
        style={{
          display: "flex",
          flex: 1,
          margin: "5em 4em 5em 4em",
      }}>
        <Content></Content>
        <img class="subscribe-logo" src="./images/crypt.png" alt="Bytecrypt Logo"
          style={{
            maxWidth: "35em",
            marginRight: "-15em",
        }}></img>
      </div>


      <div class="app-footer"
        style={{
          maxHeight: "2em",
      }}>
        <hr></hr>
        <Footer></Footer>
      </div>
    </div>
  );
}
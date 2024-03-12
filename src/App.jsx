import React, { useState } from 'react';
import './App.css';

import Header from './components/header/header';
import Content from './components/content/content';
import Footer from './components/footer/footer';

export default function App() {

  return (
    <div class="App">
      <div class="main">
        <div class="app-header">
          <Header></Header>
        </div>
        
        <div class="app-content">
          <Content></Content>
          <img class="app-content-logo" src="./images/Bytecrypt-logo_no_bg.svg" alt="Bytecrypt Logo"></img>
        </div>
      </div>
    
      <div class="app-footer">
        <hr></hr>
        <Footer></Footer>
      </div>

    </div>
  );
}
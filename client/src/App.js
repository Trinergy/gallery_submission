import React from 'react';
import ImageList from './components/imageList.js';
import './App.css';

function App() {
  return (
    <div className="App">
      <section className="section">
        <div className="container">
          <h1 className="title">
            Hello World
          </h1>
          <p className="subtitle">
            My first website with <strong>Bulma</strong>!
          </p>
        </div>
      </section>
      <ImageList />
    </div>
  );
}

export default App;

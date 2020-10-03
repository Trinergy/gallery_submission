import React from 'react';
import ImageList from './components/imageList.js';
import './App.css';

function App() {
  return (
    <div className="App has-background-primary-light">
      <section className="section">
        <div className="container">
          <h1 className="title">
            <strong className="has-text-warning-dark">Kenny's</strong> Useless Images
          </h1>
          <p className="subtitle">
            My first website with <strong className="has-text-primary-dark">Bulma</strong>!
          </p>
        </div>
      </section>
      <ImageList />
    </div>
  );
}

export default App;

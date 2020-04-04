import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <form>
          <input class="search" type="text" placeholder="Search by name"></input>
          <button type="submit"><i class="submit-search"></i></button>
        </form>
      </header>
    </div>
  );
}

export default App;

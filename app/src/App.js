import React from 'react';
import './App.css';
import './bootstrap.css'

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <form class="form-inline my-2 my-lg-0">
          <input class="form-control mr-sm-2" type="text" placeholder="Search by name"></input>
          <input class="btn btn-secondary" type="submit" value="search"></input>
        </form>
      </header>
    </div>
  );
}

export default App;

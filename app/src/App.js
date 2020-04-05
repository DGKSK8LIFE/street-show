import React from 'react';
import './App.css';
import './bootstrap.css';
import '@material-ui/core';

function App() {
  return (
    <div className="App" style={{ display: `flex`, justifyContent: `center`, alignItems: `center` }}>
      <header className="App-header">
        <form class="form-inline my-2 my-lg-0">
          <input class="form-control form-control-lg" type="text" placeholder="Search by name"></input>
          <input class="btn btn-secondary btn-lg" type="submit" value="search"></input>
        </form>
      </header>
    </div>
  );
}

export default App;

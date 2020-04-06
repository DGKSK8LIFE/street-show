import React from 'react';
import './App.css';
import LandingPage from './LandingPage';
import RestingReviewCard from './BuskerList';
import { BrowserRouter as Router, Switch, Route } from 'react-dom';

function App() {
  return (
    <Router>
      <div className="App">
        <Route path="/" component={LandingPage} />
        <Route path="/buskerlist" component={RestingReviewCard} />
      </div>
    </Router>
  )

}

export default App;

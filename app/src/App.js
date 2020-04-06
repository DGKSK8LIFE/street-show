import React from 'react';
import './App.css';
import LandingPage from './LandingPage';
import RestingReviewCard from './BuskerList';
import { BrowserRouter as Router, Route } from 'react-dom';

function App() {
  return (
    <Router>
      <Route path="/" exact component={LandingPage} />
      <Route path="/buskerlist" exact component={RestingReviewCard} />
    </Router>
  );
}

export default App;

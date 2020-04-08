import React from 'react';
import './App.css';
import LandingPage from './LandingPage';
import RecipeReviewCard from './BuskerList';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

function App() {
  return (
    <Router>
      <Route path="/" exact component={LandingPage} />
      <Route path="/buskerlist" exact component={RecipeReviewCard} />
    </Router>
  );
}

export default App;

import React from 'react';
import './App.css';
import LandingPage from './LandingPage';
import RestingReviewCard from './BuskerList';
import { BrowserRouter, Route } from 'react-dom';

function App() {
  return (
    <BrowserRouter>
      <Route path="/" exact component={LandingPage} />
      <Route path="/buskerlist" exact component={RestingReviewCard} />
    </BrowserRouter>
  );
}

export default App;

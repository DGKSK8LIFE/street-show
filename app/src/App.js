import React from 'react';
import './App.css';
import LandingPage from './LandingPage';
import  RecipeReviewCard  from './BuskerList';
import { BrowserRouter, Route } from 'react-dom';

function App() {
  return (
    <BrowserRouter>
      <Route path="/" exact component={LandingPage} />
      <Route path="/buskerlist" exact component={RecipeReviewCard} />
    </BrowserRouter>
  );
}

export default App;

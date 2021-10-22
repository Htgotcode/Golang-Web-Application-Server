// Google Analytics import, intialisation and setup of page view tracking
import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.css';

import HomePage from './views/HomePage';
import Error from './components/Error';
import Navigation from './components/Navigation';
import CardPage from './views/CardPage';
import CartPage from './views/CartPage';
import ProfilePage from './views/ProfilePage';
import AllCardsPage from './views/AllCardsPage';

// Component renders application
class App extends Component {
    render() {
      return(
        <BrowserRouter>
         <Navigation />
        <Switch>
          <Route exact path="/" component={HomePage}/>
          <Route path="/card-add" component={CardPage}/>
          <Route path="/cart" component={CartPage}/>
          <Route path="/card" component={CardPage}/>
          <Route path="/all-cards" component={AllCardsPage}/>
          <Route path="/profile" component={ProfilePage}/>
          <Route component={Error}/>
        </Switch>
      </BrowserRouter>
      );
    }
  }

export default App;

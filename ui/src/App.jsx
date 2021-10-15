import React, { Component } from 'react';
import { BrowserRouter, Link, NavLink, Route, Switch } from 'react-router-dom';

import HomePage from './views/HomePage'
import Error from './components/Error'
import CardPage from './views/CardPage';
import ProfilePage from './views/ProfilePage';
import UploadsPage from './views/SellerUploadsPage';
import MarketPage from './views/MarketPage';

class App extends Component {
    render() {
      return(
        <BrowserRouter>
         <nav>
          <ul>
            <li><a href="/">Home</a></li>
            <li><a href="/card">Card</a></li>
            <li><a href="/profile">Profile</a></li>
            <li><a href="/uploads">Your uploads</a></li>
            <li><a href="/market">Market Place</a></li>
          </ul>
        </nav>
        <Switch>
          <Route exact path="/" component={HomePage}/>
          <Route path="/card" component={CardPage}/>
          <Route path="/profile" component={ProfilePage}/>
          <Route path="/uploads" component={UploadsPage}/>
          <Route path="/market" component={MarketPage}/>
          <Route component={Error}/>
        </Switch>
      </BrowserRouter>
      );
    }
  }

export default App;

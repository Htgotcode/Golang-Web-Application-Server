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
          <Route path="/cart-create" component={CartPage}/>
          <Route path="/cart-response" component={CartPage}/>
          <Route component={Error}/>
        </Switch>
      </BrowserRouter>
      );
    }
  }

export default App;

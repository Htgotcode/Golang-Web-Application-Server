import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.css';

import HomePage from './views/HomePage';
import Error from './components/Error';
import Navigation from './components/Navigation';
import CardPage from './views/CardPage';
import ProfilePage from './views/ProfilePage';
import UploadsPage from './views/SellerUploadsPage';

class App extends Component {
    render() {
      return(
        <BrowserRouter>
         <Navigation />
        <Switch>
          <Route exact path="/" component={HomePage}/>
          <Route path="/card-add" component={CardPage}/>
          <Route path="/profile" component={ProfilePage}/>
          <Route path="/uploads" component={UploadsPage}/>
          <Route component={Error}/>
        </Switch>
      </BrowserRouter>
      );
    }
  }

export default App;

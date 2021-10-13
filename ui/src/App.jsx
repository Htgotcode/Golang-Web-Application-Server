import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Home from './home/Home'
import CardDetail from './cards/CardDetail'
import Navigation from './nav/Nav'

class App extends Component {
    render() {
      return (      
         <BrowserRouter>
          <div>
            <Navigation />
              <Switch>
               <Route path="/" component={Home} exact/>
               <Route path="/card-sell" component={CardDetail}/>
              <Route component={Error}/>
             </Switch>
          </div> 
        </BrowserRouter>
      );
    }
  }

export default App;

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
import LoggedInPage from './views/LoggedInPage';
import { Container } from 'react-bootstrap';
import GetCards from './components/Market';
const AUTH0_CLIENT_ID="xJpx6qZVNhbzuwsNse6gpcXt1oOjF1m6"
const AUTH0_DOMAIN="dev-oapcbgti.us.auth0.com"
const AUTH0_CALLBACK_URL=location.href
const AUTH0_API_AUDIENCE = "http://localhost:8080/"

class App extends React.Component {
  parseHash() {
    this.auth0 = new auth0.WebAuth({
      domain: AUTH0_DOMAIN,
      clientID: AUTH0_CLIENT_ID
    });
    this.auth0.parseHash(window.location.hash, (err, authResult) => {
      if (err) {
        return console.log(err);
      }
      if (
        authResult !== null &&
        authResult.accessToken !== null &&
        authResult.idToken !== null
      ) {
        localStorage.setItem("access_token", authResult.accessToken);
        localStorage.setItem("id_token", authResult.idToken);
        localStorage.setItem(
          "profile",
          JSON.stringify(authResult.idTokenPayload)
        );
        window.location = window.location.href.substr(
          0,
          window.location.href.indexOf("#")
        );
      }
    });
  }

  setup() {
    $.ajaxSetup({
      beforeSend: (r) => {
        if (localStorage.getItem("access_token")) {
          r.setRequestHeader(
            "Authorization",
            "Bearer " + localStorage.getItem("access_token")
          );
        }
      }
    });
  }

  setState() {
    let idToken = localStorage.getItem("id_token");
    if (idToken) {
      this.loggedIn = true;
    } else {
      this.loggedIn = false;
    }
  }

  componentWillMount() {
    this.setup();
    this.parseHash();
    this.setState();
  }

  render() {
    if (this.loggedIn) {
      return (
      <Container>
        <Navigation />
        <LoggedInPage />
        <GetCards />
      </Container>
      )
    }
    return <HomePage />;
  }
}
export default App
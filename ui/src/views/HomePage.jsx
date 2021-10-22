import React from 'react'

import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

const AUTH0_CLIENT_ID="xJpx6qZVNhbzuwsNse6gpcXt1oOjF1m6"
const AUTH0_DOMAIN="dev-oapcbgti.us.auth0.com"
const AUTH0_CALLBACK_URL="http://localhost:8080/market"
const AUTH0_API_AUDIENCE = "http://localhost:8080/"



//Home page for navigation and all cards on sale
class HomePage extends React.Component {
  constructor(props) {
    super(props);
    this.authenticate = this.authenticate.bind(this);
  }
  authenticate() {
    this.WebAuth = new auth0.WebAuth({
      domain: AUTH0_DOMAIN,
      clientID: AUTH0_CLIENT_ID,
      scope: "openid profile",
      audience: AUTH0_API_AUDIENCE,
      responseType: "token id_token",
      redirectUri: AUTH0_CALLBACK_URL
    });
    this.WebAuth.authorize();
  }
    render() {
        return (
          <div className="container">
            <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
              <h1>Card Trading World</h1>
              <div className="container align-contents-center">
                  <img href="./images/logo255.png"/>
              </div>
              <p>Sign in to get access </p>
              <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
            </div>
          </div>
        )
      }
 }

export default HomePage
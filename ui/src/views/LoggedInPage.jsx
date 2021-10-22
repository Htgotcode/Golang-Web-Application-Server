import React from 'react'
import ReactGA from 'react-ga';
import GetCards from '../components/Market';
import CardPage from './CardPage';
import HomePage from './HomePage';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

class LoggedIn extends React.Component {
    constructor(props) {
      super(props);
      this.serverRequest = this.serverRequest.bind(this);
      this.logout = this.logout.bind(this);
    }
  
    logout() {
      localStorage.removeItem("id_token");
      localStorage.removeItem("access_token");
      localStorage.removeItem("profile");
      location.reload();
    }
  
    serverRequest() {
      $.get("http://localhost:8080/card", res => {
        this.setState({
            card: res
        });
      });
    }
  
    componentDidMount() {
      this.serverRequest();
    }
  
    render(){
        return(
            <div className="container" >
              <a onClick={this.logout}>Log out</a>
                <CardPage />
            </div>
        );
    }
  }

  export default LoggedIn
import React from 'react'
import  GetCards from '../components/Market';

// Google Analytics import, intialisation and setup of page view tracking
import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

//Home page for navigation and all cards on sale
// Component renders market place cards
class HomePage extends React.Component {
    render(){
        return(
            <div className="container" >
                <GetCards />
            </div>
        );
    }
 }

export default HomePage
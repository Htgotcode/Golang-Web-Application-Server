import React from 'react'
import  AllCards from '../components/AllCards';

import ReactGA from 'react-ga';

// Google Analytics import, intialisation and setup of page view tracking
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

//Home page for navigation and all cards on sale
class AllCardsPage extends React.Component {
    render(){
        return(
            <div className="container" >
                <AllCards />
            </div>
        );
    }
 }

export default AllCardsPage
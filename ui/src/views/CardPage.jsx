//Card details, payment details, card listing
import React from 'react'
import  AllCards from '../components/AllCards';

// Google Analytics import, intialisation and setup of page view tracking
import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);
//Home page for navigation and all cards on sale
//Component calls AllCards function
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
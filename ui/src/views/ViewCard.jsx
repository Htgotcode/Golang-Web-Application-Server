import React from 'react'
import Cardlisting from '../components/CardListing';
import ReactGA from 'react-ga';

// Google Analytics import, intialisation and setup of page view tracking
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

//Component calls the GetAccount function and renders the profile page
class ViewCard extends React.Component {
    render(){
        return(
            <div>
                <CardListing/>
            </div>
        );
    }
}

export default ViewCard
import React from 'react'
import RenderCardListing from '../components/CardListing';

// Google Analytics import, intialisation and setup of page view tracking
import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

// Component calls RenderCardListing function and renders card listing page
class CardListingPage extends React.Component {
    render(){
        return(
            <div>
                <RenderCardListing/>
            </div>
        );
    }
}

export default CardListingPage
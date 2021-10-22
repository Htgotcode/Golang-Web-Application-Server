//Card details, payment details, card listing
import React from 'react'
import RenderCardListing from '../components/CardListing';

import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

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
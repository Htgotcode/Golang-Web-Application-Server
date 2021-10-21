//Card details, payment details, card listing
import React from 'react'
import AddCard from '../components/AddCard';

import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);
class CardPage extends React.Component {
    render(){
        return(
            <div>
            <AddCard/>
            </div>
        );
    }
}

export default CardPage
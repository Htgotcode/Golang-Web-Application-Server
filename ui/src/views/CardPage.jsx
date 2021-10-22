import React from 'react'
import AddCard from '../components/AddCard';

// Google Analytics import, intialisation and setup of page view tracking
import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

//Home page for navigation and all cards on sale
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
import React from 'react'
import  GetCards from '../components/Market';

import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

//Home page for navigation and all cards on sale
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
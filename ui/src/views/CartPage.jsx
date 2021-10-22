import React from 'react'
import RenderCart from '../components/Cart';

import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

class CartPage extends React.Component {
    render(){
        return(
            <div className="container">
                <RenderCart/>
            </div>
        );
    }
}

export default CartPage
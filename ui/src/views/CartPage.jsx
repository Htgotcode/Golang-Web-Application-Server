import React from 'react'
import RenderCart from '../components/Cart';

// Google Analytics import, intialisation and setup of page view tracking
import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);

// Component renders shopping cart
class CartPage extends React.Component {
    render(){
        return(
            <div className="container">
                <div className="row">
                    <div className="col-9">
                        <div id="shopping-cart">
                            <RenderCart/>
                        </div>
                    </div>
                    <div className="col-3">
                        <div id="shopping-cart-details">

                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default CartPage
import React from 'react'
import RenderCart from '../components/Cart';

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
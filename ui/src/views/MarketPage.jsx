//Market & Filter
import React from 'react'
import Market from '../components/Market';

class MarketPage extends React.Component {
    render(){
            return(
                <div className="container">
                    <div className="row">
                        <div className="col-2 bg-primary">
                            <p>Filter Form</p>
                        </div>
                        <div className="col-10">
                            <div className="bg-light">
                                <Market/>
                            </div>
                        </div>
                    </div>
                </div>
            );
        }
}

export default MarketPage
import React from 'react'
import Market from '../components/Market';

//Home page for navigation and all cards on sale
class HomePage extends React.Component {
    render(){
        return(
            <div className="container" style={{ backgroundImage: "url(./images/logo255.png)" }}>
                <div className="row">
                    <div className="col-2">
                        <p>Filter Form</p>
                    </div>
                    <div className="col-10">
                        <div className="">
                            <Market/>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
 }

export default HomePage
import React from 'react'
import  GetCards from '../components/Market';

//Home page for navigation and all cards on sale
class HomePage extends React.Component {
    render(){
        return(
                <div className="container p-0" >
                    <div className="row">
                        <div className="col-2">
                            <p>Filter Form</p>
                        </div>
                        <div className="col-10">
                            <div className="">
                                <GetCards />
                            </div>
                        </div>
                    </div>
                </div>
        );
    }
 }

export default HomePage
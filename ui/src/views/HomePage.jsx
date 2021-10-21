import React from 'react'
import  GetCards from '../components/Market';

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
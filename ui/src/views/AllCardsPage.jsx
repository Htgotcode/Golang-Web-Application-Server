import React from 'react'
import  AllCards from '../components/AllCards';

//Home page for navigation and all cards on sale
class AllCardsPage extends React.Component {
    render(){
        return(
            <div className="container" >
                <AllCards />
            </div>
        );
    }
 }

export default AllCardsPage
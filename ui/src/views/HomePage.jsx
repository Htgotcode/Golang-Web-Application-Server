import React from 'react'

//Home page for navigation and all cards on sale
class HomePage extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            name: "Kyle" 
        }
    }
    render(){
        return (
            <div>
                <h1>This is the Home Page {this.state.name}</h1>
            </div>
            );
    }
 }

// function Welcome(props){
//     return <h1>Hello, {props.name}</h1>
// }

export default HomePage
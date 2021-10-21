//Payment History & Payment details
import React from 'react'
import GetAccount from '../components/Account';
import ReactGA from 'react-ga';
ReactGA.initialize('UA-210763128-1');
ReactGA.pageview(window.location.pathname + window.location.search);
class ProfilePage extends React.Component {
    render(){
        return(
            <div>
                <GetAccount/>
            </div>
        );
    }
}

export default ProfilePage
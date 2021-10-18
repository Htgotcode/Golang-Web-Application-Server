//Payment History & Payment details, Sales History & Account details
import React from 'react'

import PurchaseHistory from './components/PurchaseHistory';
import PaymentDetails from './components/PaymentDetails';

class ProfilePage extends React.Component {
    constructor(props) {
        super(props);
    
        this.state = {
          humor: 'happy'
        }
      }
      render() {
        return(
          <div>
            <h1>{this.props.name}</h1>
            <p>
              {this.props.color}
            </p>
          </div>
        );
      }
}

export default ProfilePage
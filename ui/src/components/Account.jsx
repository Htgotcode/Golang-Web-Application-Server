
import React, { Children, useState, useEffect } from 'react';
import MaterialIcon from 'material-icons-react';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';
import axios from "axios";
import { Container } from 'react-bootstrap';
// Function to get account data and render it on the profile page 
function ViewAccount() {
    const [ACCOUNTS, setACCOUNTS] = useState([])
    const [isLoading, setLoading] = useState(true);

    useEffect(() => {
      getAccounts();
    }, []);

    const getAccounts = () => {
      axios.get('/account')
        .then((response) => {
          setACCOUNTS(response.data);
          setLoading(false);
      });
    };

    if (isLoading){
      return <div className="App">Loading...</div>;
    } else {
      return (
        <div className="container p-0">
              {ACCOUNTS.map((account) => {
                     return(
                       <div key={account._id}>
                                {/* Account details for current user */}
                                <h1>Account Details</h1>
                                <p>
                                    Username: {account.username}
                                </p>
                                <p>
                                    Password: {account.password}
                                </p>
                                <p>
                                    Email Address: {account.email}
                                </p>
                                
                                {/* Purchase History table*/}
                                <h2>Purchase History</h2>
                                <table title="Purchase History" class="table table-striped">
                                  <thead>
                                      <tr>
                                      <th>Name</th>
                                      <th>Description</th>
                                      <th>Brand</th>
                                      <th>SetName</th>
                                      <th>Rarity</th>
                                      <th>SellingPrice</th>
                                      <th>UploadedAt</th>
                                      <th>CardID</th>
                                      </tr>
                                  </thead>
                                  <tbody>
                                      <tr>
                                      <td>{account.purchaseHistory.Name}</td>
                                      <td>{account.purchaseHistory.Description}</td>
                                      <td>{account.purchaseHistory.Brand}</td>
                                      <td>{account.purchaseHistory.SetName}</td>
                                      <td>{account.purchaseHistory.Rarity}</td>
                                      <td>${account.purchaseHistory.SellingPrice}</td>
                                      <td>{account.purchaseHistory.UploadedAt}</td>
                                      <td>{account.purchaseHistory.CardID}</td>
                                      </tr>
                                  </tbody>
                                </table>

                                {/* Sale History table */}
                                <h2>Sale History</h2>
                                <table title="Sale History" class="table table-striped">
                                <title>Sale History</title>
                                  <thead>
                                      <tr>
                                      <th>Name</th>
                                      <th>Description</th>
                                      <th>Brand</th>
                                      <th>SetName</th>
                                      <th>Rarity</th>
                                      <th>SellingPrice</th>
                                      <th>UploadedAt</th>
                                      <th>CardID</th>
                                      </tr>
                                  </thead>
                                  <tbody>
                                      <tr>
                                      <td>{account.saleHistory.Name}</td>
                                      <td>{account.saleHistory.Description}</td>
                                      <td>{account.saleHistory.Brand}</td>
                                      <td>{account.saleHistory.SetName}</td>
                                      <td>{account.saleHistory.Rarity}</td>
                                      <td>${account.saleHistory.SellingPrice}</td>
                                      <td>{account.saleHistory.UploadedAt}</td>
                                      <td>{account.saleHistory.CardID}</td>
                                      </tr>
                                  </tbody>
                                </table>
                      </div>
                    );     
                  }
              ) 
            }
        </div>
      );
    } 
}
// Component that calls the ViewAccount function
class GetAccount extends React.Component {
  render() {
      return(
        <div>
          <ViewAccount/>
      </div>
      )
  }
}
 
export default GetAccount
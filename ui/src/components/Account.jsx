
import React, { Children, useState, useEffect } from 'react';
import MaterialIcon from 'material-icons-react';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';
import axios from "axios";
import { Container } from 'react-bootstrap';
 
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
                      <p key={account._id}>
                          {account.purchaseHistory.Name}
                      </p>
                      <p key={account._id}>
                          {account.purchaseHistory.Description}
                      </p>
                      <p key={account._id}>
                          {account.purchaseHistory.Brand}
                      </p>
                      <p key={account._id}>
                          {account.purchaseHistory.SetName}
                      </p>
                      <p key={account._id}>
                          {account.purchaseHistory.Rarity}
                      </p>
                      <p key={account._id}>
                          ${account.purchaseHistory.SellingPrice}
                      </p>
                      <p key={account._id}>
                          {account.purchaseHistory.UploadedAt}
                      </p>
                      <p key={account._id}>
                          {account.purchaseHistory.CardID}
                      </p>
                      </div>
                    );     
                  }
              ) 
            }
        </div>
      );
    } 
}

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

// type Account struct {
// 	ID        primitive.ObjectID `bson:"_id"`
// 	Username  string             `bson:"username"`
// 	Password  string             `bson:"password"`
// 	Email     string             `bson:"email"`
// 	CreatedAt time.Time          `bson:"created_at"`

// 	PurchaseHistory []struct {
// 		Name         string    `bson:"name"`
// 		Description  string    `bson:"description"`
// 		Brand        string    `bson:"brand" validate:"required"`
// 		SetName      string    `bson:"set_name" validate:"required"`
// 		Rarity       string    `bson:"rarity" validate:"required"`
// 		SellingPrice uint      `bson:"selling_price" validate:"required"`
// 		UploadedAt   time.Time `bson:"uploaded_at"`
// 		CardID       string    `bson:"card_id"`
// 	} `json:"purchaseHistory"`

// 	SaleHistory []struct {
// 		Name         string    `bson:"name"`
// 		Description  string    `bson:"description"`
// 		Brand        string    `bson:"brand" validate:"required"`
// 		SetName      string    `bson:"set_name" validate:"required"`
// 		Rarity       string    `bson:"rarity" validate:"required"`
// 		SellingPrice uint      `bson:"selling_price" validate:"required"`
// 		UploadedAt   time.Time `bson:"uploaded_at"`
// 		CardID       string    `bson:"card_id"`
// 	} `json:"saleHistory"`

{/* <form onSubmit={this.handleSubmit}>
          <label>
          name:
            <input type="text" value={this.state.name} name="name" onChange={this.handleName} />
          </label>
          <label>
          description:
            <input type="text" value={this.state.description} name="description" onChange={this.handleDescription} />
          </label>
          <label>
          brand:
            <input type="text" value={this.state.brand} name="brand" onChange={this.handleBrand} />
          </label>
          <label>
          setname:
            <input type="text" value={this.state.setname} name="setname" onChange={this.handleSetName} />
          </label>
          <label>
          rarity:
            <input type="text" value={this.state.rarity} name="rarity" onChange={this.handleRarity} />
          </label>
          <label>
          sellingprice:
            <input type="text" value={this.state.sellingprice} name="sellingprice" onChange={this.handleSellingPrice} />
          </label>
          <label>
          imageurl:
            <input type="text" value={this.state.imageurl} name="imageurl" onChange={this.handleImageURL} />
          </label>
          <label>
          ownerid:
          <input type="text" value={this.state.ownerid} name="ownerid" onChange={this.handleOwnerID} />
          </label>
          <button type="submit">Add</button>
        </form> */}

// import React, {useState, useEffect} from 'react';

// import axios from "axios";

// import {Button, Form, Container, Modal } from 'react-bootstrap'

// import Account from './SingleAccountComponent';

// class Accounts extends React.Component {

//     handleSubmit = event => {

//             event.preventDefault();
        
//             const accounts = {
        
//               username: this.state.username,
        
//               password: this.state.password,
        
//               email: this.state.email,
        
//               createdat: this.state.createdat,
        
//               purchaseHistory: this.state.purchaseHistory,
        
//               saleHistory: this.state.saleHistory,                  
//             };
        
//             getAllAccounts()
        
//             axios.post(`/account-view`, accounts, {headers: {'Content-Type': 'application/json'}})
        
//               .then(res => {
        
//                 console.log(res);
        
//                 console.log(res.data);
        
//               })
        
//           }

//     render() {
//         const(accounts) = this.state;
//         console.log(accounts)
//         return (
//             <div>
//                 <table class="table table-striped">
//                     <thead>
//                         <tr>
//                         <th>ID</th>
//                         <th>Username</th>
//                         <th>Password</th>
//                         <th>Email</th>
//                         <th>CreatedAt</th>
//                         <th>PurchaseHistory</th>
//                         <th>SaleHistory</th>
//                         </tr>
//                     </thead>
//                     <tbody>
//                         {accounts}
//                     </tbody>
//                 </table>
//             </div>
            
//         );
//     }
// }

// function getAllAccounts(){
//     var url = "http://localhost:8080/account"
//     axios.get(url, {
//         responseType: 'json'
//     }).then(response => {
//         if(response.status == 200){
//             setAccounts(response.data)
//         }
//     })
// }

// export default Accounts

// axios.post(`/account-view`, account, {headers: {'Content-Type': 'application/json'}})
//           .then(res => {
//             console.log(res);
//             console.log(res.data);
//           })
//       }

//       return (
//         <div>
//             {/* account profile details */}
//             <Container>
//             <table class="table table-striped">
//                   <thead>
//                       <tr>
//                       <th>ID</th>
//                       <th>Username</th>
//                       <th>Password</th>
//                       <th>Email</th>
//                       <th>CreatedAt</th>
//                       </tr>
//                   </thead>
//                   <tbody>
//                   {account != null && accounts.map((order, i) => (
//                       <Account accountData={account} deleteSingleOrder={deleteSingleOrder} setChangeWaiter={setChangeWaiter} setChangeOrder={setChangeOrder}/>
//                   ))}
//                   </tbody>
//           </table>
//             </Container>
  
//             <Container>
//                 {/* purchase history */}
//             </Container>
  
//             <Container>
//                 {/* sale history */}
//             </Container>
//         </div>
//       )
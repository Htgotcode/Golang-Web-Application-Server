// import React, {useState, useEffect} from 'react';

// import axios from "axios";

// import {Button, Form, Container, Modal } from 'react-bootstrap'

// import Order from './single-order.component';

// const PurchaseHistory = () => {

//     const [orders, setOrders] = useState([])
//     const [refreshData, setRefreshData] = useState(false)

//     //gets run at initial loadup
//     useEffect(() => {
//         getPurchaseHistory();
//     }, [])

//     //refreshes the page
//     if(refreshData){
//         setRefreshData(false);
//         getPurchaseHistory();
//     }

//     return (
//         <div>
//             <getPurchaseHistory/>
//         </div>
        
//     );

//     function getPurchaseHistory(){
//         var url = "http://localhost:8080/profile"
//         axios.get(url, {
//             responseType: 'json'
//         }).then(response => {
//             if(response.status == 200){
//                 setOrders(response.data)
//             }
//         })
//     }
// }

// export default PurchaseHistory
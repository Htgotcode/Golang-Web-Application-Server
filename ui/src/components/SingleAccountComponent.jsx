import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';

import {Button, Card, Row, Col} from 'react-bootstrap'

const Account = ({accountData, setChangeWaiter, deleteSingleOrder, setChangeOrder}) => {

    return (
        <Card>
            <Row>
                <Col>ID:{ accountData !== undefined && accountData.id}</Col>
                <Col>Username:{ accountData !== undefined && accountData.username}</Col>
                <Col>Password:{ accountData !== undefined && accountData.password}</Col>
                <Col>Email: ${accountData !== undefined && accountData.email}</Col>
                <Col>CreatedAt: ${accountData !== undefined && accountData.createdAt}</Col>
                <Col>PurchaseHistory: ${accountData !== undefined && accountData.purchaseHistory}</Col>
                <Col>SaleHistory: ${accountData !== undefined && accountData.saleHistory}</Col>
            </Row>
        </Card>
    )

    // function changeWaiter(){
    //     setChangeWaiter(
    //         {
    //             "change": true,
    //             "id": orderData._id
    //         }
    //     )
    // }

    // function changeOrder(){
    //     setChangeOrder(
    //         {
    //             "change": true,
    //             "id": orderData._id
    //         }
    //     )
    // }

}

export default Account
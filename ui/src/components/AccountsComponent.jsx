import React, {useState, useEffect} from 'react';

import axios from "axios";

import {Button, Form, Container, Modal } from 'react-bootstrap'

import Account from './SingleAccountComponent';

const Accounts = () => {

    const [accounts, setAccounts] = useState([])
    const [refreshData, setRefreshData] = useState(false)

    useEffect(() => {
        getAllAccounts();
    }, [])

    if(refreshData){
        setRefreshData(false);
        getAllAccounts();
    }

    return (
        <div>
            {/* list all current accounts */}
            <Container>
                {accounts != null && accounts.map((account, i) => (
                    <Account accountData={account}/>
                ))}
            </Container>
        </div>
        
    );

    function getAllAccounts(){
        var url = "http://localhost:8080/profile"
        axios.get(url, {
            responseType: 'json'
        }).then(response => {
            if(response.status == 200){
                setOrders(response.data)
            }
        })
    }
}

export default Accounts
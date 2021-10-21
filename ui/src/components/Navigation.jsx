import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Container from 'react-bootstrap/Container';
import Image from 'react-bootstrap/Image'
import { Cart2 } from 'react-bootstrap-icons';

class Navigation extends React.Component {
    render(){
        return(
            <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand href="/"><Image className="d-inline-block align-top" src="./images/logo255.png" width="40" height="40" alt="Logo"/></Navbar.Brand>
                    <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                    <Navbar.Collapse id="responsive-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link className="h4" href="/card-add">Upload Card</Nav.Link>
                        <Nav.Link className="h4" href="/all-cards">All Cards</Nav.Link>
                    </Nav>
                    <Nav>
                        <Nav.Link className="" href="/cart"><Cart2 size={20}/></Nav.Link>
                        <Nav.Link className="text-white" href="/profile"><Image src="./images/logo255.png" width="30" height="30" alt="Profile Picture"/><br/>NAME GOES HERE</Nav.Link>
                    </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        );
    }
}

export default Navigation
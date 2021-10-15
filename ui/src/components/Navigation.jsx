import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Container from 'react-bootstrap/Container';
import Image from 'react-bootstrap/Image'
import 'bootstrap/dist/css/bootstrap.css';

class Navigation extends React.Component {
    render(){
        return(
            <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand href="/"><Image className="d-inline-block align-top" src="./logo.svg" width="40" height="40" alt="Logo"/></Navbar.Brand>
                    <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                    <Navbar.Collapse id="responsive-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link className="h4" href="/">Home</Nav.Link>
                        <Nav.Link className="h4" href="/card">Card</Nav.Link>
                        <Nav.Link className="h4" href="/uploads">Uploads</Nav.Link>
                        <Nav.Link className="h4" href="/market">Market</Nav.Link>
                    </Nav>
                    <Nav>
                        <Nav.Link className="text-white" href="/profile"><Image src="" width="30" height="30" alt="Profile Picture"/></Nav.Link>
                    </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        );
    }
}

export default Navigation
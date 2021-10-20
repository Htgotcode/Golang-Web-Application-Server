import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Container from 'react-bootstrap/Container';
import Image from 'react-bootstrap/Image'

class Navigation extends React.Component {
    render(){
        return(
            <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand href="/"><Image className="d-inline-block align-top" src="./images/logo255.png" width="50" height="50" alt="Logo"/></Navbar.Brand>
                    <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                    <Navbar.Collapse id="responsive-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link className="h4" href="/">Home</Nav.Link>
                        <Nav.Link className="h4" href="/card-add">Card</Nav.Link>
                        <Nav.Link className="h4" href="/uploads">Uploads</Nav.Link>
                    </Nav>
                    <Nav>
                        <Nav.Link className="text-white" href="/profile"><Image src="./images/logo255.png" width="30" height="30" alt="Profile Picture"/><br/>NAME GOES HERE</Nav.Link>
                    </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        );
    }
}

export default Navigation
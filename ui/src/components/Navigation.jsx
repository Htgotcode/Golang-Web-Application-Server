import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Button from 'react-bootstrap/Button';
import Nav from 'react-bootstrap/Nav';
import Container from 'react-bootstrap/Container';
import Image from 'react-bootstrap/Image'


// Component allows for navigating between pages via the navigation bar
class Navigation extends React.Component {
    constructor(props) {
        super(props);
        this.logout = this.logout.bind(this);
      }
    
    logout() {
        localStorage.removeItem("id_token");
        localStorage.removeItem("access_token");
        localStorage.removeItem("profile");
        location.reload();
    }

    render(){
            return(
                <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
                    <Container>
                        <Navbar.Brand href="/market"><Image className="" src="./images/logo255.png" width="40" height="40" alt="Logo"/></Navbar.Brand>
                        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                        <Navbar.Collapse id="responsive-navbar-nav">
                        <Nav className="me-auto">
                            <Nav.Link className="h4" href="/market">Market</Nav.Link>
                            <Nav.Link className="h4" href="/card-add">Upload Card</Nav.Link>
                            <Nav.Link className="h4" href="/all-cards">All Cards</Nav.Link>
                            <Nav.Link className="h4" href="/cart">Cart</Nav.Link>
                        </Nav>
                        <Nav>
                            <Button type ="submit" variant="secondary" value="Logout" onClick={this.logout} >Log-out</Button>
                            <Nav.Link className="text-white" href="/profile"><Image src="./images/logo255.png" width="30" height="30" alt="Profile Picture"/><br/>JohnDoe</Nav.Link>
                        </Nav>
                        </Navbar.Collapse>
                    </Container>
                </Navbar>
            );
        }
}

export default Navigation
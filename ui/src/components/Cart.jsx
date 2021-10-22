import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Card from 'react-bootstrap/Card';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import Row from 'react-bootstrap/Row';

function Cart() {
  const [isLoading, setLoading] = useState(true);
  const [Cart, setCart] = useState([]);


  useEffect(() => {
      getCart();

  }, []);

  const getCart = () => {
    axios.get(`/cart-response`)
      .then((response) => {
      setCart(response.data);
      setLoading(false);
    });
  };


  const buyCart = (cartItem) => {
    axios.delete(`/cart-response${cartItem._id}`, {headers: {'Content-Type': 'application/json'}, data: {_id: cartItem._id}})
          .then(res => {
      console.log(res);
      console.log(res.data);
    })
    
    alert("Email sent to relevant card owners.")
  };

  
    if (isLoading){
      return <div className="App">Loading...</div>;
    } else {
      if(Cart != null){
        return (
          <Container>
            
            {Cart.map((cartItem) => {
              if(cartItem.email=="johndoe@example.com"){
                    return (
                      <Row> 
                        {/* MAKE USER NAME */}
                        <h1>{cartItem.email}'s Cart</h1>
                          {cartItem.cards.map((card) => {
                                return (
                                    <Col xs={4} md={4} lg={2} xl={2} xxl={2} sm={4} className="shadow rounded m-3" key={card._id}>
                                      <Card>
                                        <Card.Header as="h5">{card.name} - {card.brand}</Card.Header>
                                        <Card.Body>
                                        <Card.Text>{card.description}</Card.Text>
                                        <Card.Text>{card.setname}</Card.Text>
                                        <Card.Text>{card.rarity}</Card.Text>
                                        <Card.Text>${card.sellingprice}</Card.Text>
                                        </Card.Body>
                                    </Card>  
                                  </Col> 
                                )
                              })
                            } 
                          <div className="h5">Total: ${cartItem.total}</div>
                          <a href="/cart"><Button type ="submit" value="Buy" onClick={() => buyCart(cartItem)}>Buy Cart</Button></a>
                          </Row>
                    )
                  } else {
                    return(
                      <div className="App">
                        Cart is empty.
                      </div>
                    )
                    }
                  })
                } 
            
            </Container>
    );
    } else {
        return(
          <div className="App">
            Cart is empty.
          </div>
        )
      } 
    }
  }
  
  class RenderCart extends React.Component {
      render() {
            return (
              <div className="App">
              <Cart />
              </div>
          );
      }
  }
    
  
export default RenderCart
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Card from 'react-bootstrap/Card';
import Col from 'react-bootstrap/Col';
import Button from 'react-bootstrap/Button';

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
  };
  
  
    if (isLoading){
      return <div className="App">Loading...</div>;
    } else {
      
        return (
          <div className="container">
            <div className="row">
            {Cart.map((cartItem) => {
              if(cartItem.userid=='test'){
                    return (
                      <div> 
                        {/* MAKE USER NAME */}
                    <h1>{cartItem.userid}'s Cart</h1>
                    <div>
                      {cartItem.cards.map((card) => {
                            return (
                            <div> 
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
                            </div>  
                            )
                          })
                        } 
                        <Button type ="submit" value="Buy" onClick={() => buyCart(cartItem)}>Buy Cart</Button>
                      </div>
                    </div>
                    )
                  } else {
                    return(
                      <div>
                        User not Found.
                      </div>
                    )
                    }
                  })
                } 
            </div>
          </div>
          
      );
      
     
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
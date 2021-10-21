import React, { useState, useEffect } from 'react';
import { Search } from 'react-bootstrap-icons';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
import axios from "axios";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { Cart2 } from 'react-bootstrap-icons';

function Market() {
  const [CARDS, setCARDS] = useState([]);
  const [isLoading, setLoading] = useState(true);
  const [Cart, setCart] = useState([]);
  const [Name, setName] = useState('');
  const [foundCard, setFoundCard] = useState(CARDS);
  var [cartCount, setCartCount] = useState(0);

  useEffect(() => {
    getCards();
  }, []);

  const getCards = () => {
    axios.get('/card')
      .then((response) => {
      setCARDS(response.data);
      setFoundCard(response.data);
      setLoading(false);
    });
  };

    const removeCard = (card) => {  
        axios.delete(`/card${card._id}`,{headers: {'Content-Type': 'application/json'}, data: {_id: card._id}})
            .then(response => {
              console.log(response);
      })
    };

    const createCart = () => {
      if(cartCount > 0){
        axios.post(`/cart-create`, { userid: "test", cards: Cart }, {headers: {'Content-Type': 'application/json'}})
          .then(res => {
            console.log(res);
            console.log(res.data);
      })
      }
      
    }

  const addToCart = (card) => {
    setCart((Cart) => [...Cart, card]);
    setCartCount(cartCount + 1);
  };

  const clearCart = () => {
    setCart([]);
    setCartCount(0);
    alert("Cart cleared.");
  };

  const filter = (e) => {
    const keyword = e.target.value;

    if (keyword !== '') {
      const RESULTS = CARDS.filter((card) => {
        return card.name.toLowerCase().includes(keyword.toLowerCase());
      });
      setFoundCard(RESULTS);
    } else {
      setFoundCard(CARDS);
      // If the text field is empty, show all cards
    }

    setName(keyword);
  };

  if (isLoading){
    return <div className="App">Loading...</div>;
  } else {
    return (
      <Container>
        <Container>
        <span className="align-middle"><Search size={15}/>  </span>
        <span className="align-middle">
        <input
          type="search"
          value={Name}
          onChange={filter}
          className="input"
          placeholder="Search name..."
          size="50"
        />    </span>
        <span className="align-right"><Button input type ="submit" value="Clear cart" onClick={() => clearCart()}>Clear Cart</Button>    </span>
        <span className="align-right"><a href="/cart"><Cart2 size={20} onClick={() => createCart()}/></a>    {cartCount}</span>
        </Container>
            <Container className="mt-3">
              <Row>
                  {foundCard && foundCard.length > 0 ? (
                        foundCard.map((card) => {
                        return (

                          <Col xs={4} md={4} lg={2} xl={2} xxl={2} sm={4} className="shadow rounded m-3" key={card._id}>
                              <Card>
                                <Card.Header as="h5">{card.name} - {card.brand}</Card.Header>
                                <Card.Body>
                                <Card.Img variant="top" src={card.imageurl} />
                                <Card.Text>{card.description}</Card.Text>
                                <Card.Text>{card.setname}</Card.Text>
                                <Card.Text>{card.rarity}</Card.Text>
                                <Card.Text>${card.sellingprice}</Card.Text>
                                <Button type ="submit" value="Add to cart" onClick={() => addToCart(card)}>Add to Cart</Button>
                                <span>      </span>
                                <Button type ="submit" value="Delete" onClick={() => removeCard(card)}>Delete Card</Button>
                                </Card.Body>
                            </Card>  
                          </Col> 
                          )
                      }
                    )
                  ) : (
                    <p>Card not found.</p>
                  )}
              
            </Row>
          </Container>
      </Container>
    );
  }
  
}

class GetCards extends React.Component {
    render() {
          return (
            <div>
            <Market />
            </div>
        );
    }
}
  



export default GetCards
import React, { useState, useEffect } from 'react';
import { Search } from 'react-bootstrap-icons';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';
import axios from "axios";

function Market() {
  const [CARDS, setCARDS] = useState([])
  const [isLoading, setLoading] = useState(true);
  const [Cart, setCart] = useState([]);
  const [Name, setName] = useState('');
  const [foundCard, setFoundCard] = useState(CARDS);
  
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

  const addToCart = (card) => {
    setCart([...Cart, card]);
    alert(card.name +" Added to cart.");
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
      <div className="container p-0">
        <span className="align-middle"><Search size={15}/>  </span>
        <span className="align-middle">
        <input
          type="search"
          value={Name}
          onChange={filter}
          className="input"
          placeholder="Search name..."
          size="80"
        /></span>
        <div className="container">
          <div className="row">
            {foundCard && foundCard.length > 0 ? (
                  foundCard.map((card) => {
                  return (
                  <div className="col-3" key={card._id}>
                        <Card className="m-1">
                        <div className="">
                          <Card.Img variant="top" src={card.imageurl} />
                        </div>
                        <Card.Body>
                          <Card.Title>{card.name}</Card.Title>
                          <Card.Text>
                          {card.description}
                          </Card.Text>
                            <ListGroup variant="flush">
                            <ListGroup.Item>{card.brand}</ListGroup.Item>
                            <ListGroup.Item>{card.setname}</ListGroup.Item>
                            <ListGroup.Item>{card.rarity}</ListGroup.Item>
                            <ListGroup.Item>{card.sellingprice}</ListGroup.Item>
                            <ListGroup.Item>{card.ownerid}</ListGroup.Item>
                          </ListGroup>
                          <input type ="submit" value="Add to cart" onClick={() => addToCart(card)}/>
                        </Card.Body>
                      </Card>  
                    </div> 
                  )
                }
              )
            ) : (
              <p>Card not found.</p>
            )}
          </div>
        </div>
      </div>
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
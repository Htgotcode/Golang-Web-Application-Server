import React, { Children, useState, useEffect } from 'react';
import MaterialIcon from 'material-icons-react';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';
import axios from "axios";


function Market() {
  const [CARDS, setCARDS] = useState([])
  const [isLoading, setLoading] = useState(true);
  const [Cart, setCart] = useState([]);
  console.log(Cart);
  const [Name, setName] = useState('');
  const [foundCard, setFoundCard] = useState(CARDS);

  useEffect(() => {
    getCards();
  }, []);

  const getCards = () => {
    axios.get('/card')
      .then((response) => {
      setCARDS(response.data);
      setFoundCard(response.data)
      setLoading(false);
    });
  };

  console.log(CARDS)
  const addToCart = (card) => {
    setCart([...Cart, card]);
  };

  const filter = (e) => {
    const keyword = e.target.value;

    if (keyword !== '') {
      const RESULTS = CARDS.filter((card) => {
        return card.name.toLowerCase().startsWith(keyword.toLowerCase());
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
        <span className="align-middle"><MaterialIcon icon="search" size="tiny"/></span>
        <span className="align-middle">
        <input
          type="search"
          value={Name}
          onChange={filter}
          className="input"
          placeholder="Search"
          size="40"
        /></span>
        <div className="container">
          <div className="row">
            {foundCard && foundCard.length > 0 ? (
                  foundCard.map((card) => {
                    return (
                  <div className="col-3" key={card._id}>
                        <Card className="m-1">
                        <Card.Img variant="top" src={card.imageurl} />
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
                          {/* <input type ="submit" value="Add to cart" onClick={() => addToCart(card)}/> */}
                        </Card.Body>
                        <input type ="submit" value="Add to cart" onClick={() => addToCart(card)}/>
                      </Card>  
                    </div> 
                  )
                }
              )
            ) : (
              <p>Pokemon not found.</p>
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
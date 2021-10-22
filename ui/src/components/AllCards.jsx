import React, { useState, useEffect } from 'react';
import { Search } from 'react-bootstrap-icons';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';
import pokemon from 'pokemontcgsdk'
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';

// Pokemon API configuration key
pokemon.configure({apiKey: '8c44560c-5a0a-4e9e-828a-898992ad6345'})

// Function to get all pokemon cards
function CardsBase() {
    const [CARDS, setCARDS] = useState([])
    const [isLoading, setLoading] = useState(true);
    const [Name, setName] = useState('');
    
    useEffect(() => {
      getAllCards();
    }, [])

    const getAllCards = () => {
      pokemon.card.where({ q:'' })
      .then(response => {
        setCARDS(response.data); 
        setLoading(false);
    })
  };

  const handleSearchInput = (e) => {
    setName(e.target.value)
  };  

    const handleSearch = (e) => {
      e.preventDefault();
      pokemon.card.where({ q:'name:'+Name, pageSize: 50 })
        .then(response => {
            setCARDS(response.data); 
            setLoading(false);
      })
    }
  
    if (isLoading){
      return <div className="App">Loading...</div>;
    } else {
      return (
        <div className="container p-0">
          <Form onSubmit={handleSearch}>
          <span className="align-middle"><Search size={15}/>  </span>
          <Form.Group className="m-1" controlId="formSearch">
            <Form.Label>Search</Form.Label>
            <Form.Control type="search" className="input" name="name" placeholder="Search full card database..." onChange={handleSearchInput} value={Name}/>
            <Form.Text className="text-muted">
              eg. Pikachu
            </Form.Text>
          </Form.Group>
          
           <span className="align-middle">
           <Button variant="primary" type="submit">
            Search
          </Button>
          </span>
          </Form>
          

          <div className="container">
            <div className="row">
                    {CARDS.map((card) => {
                    return (
                    <div className="col-3" key={card._id}>
                          <Card className="m-1">
                          <div className="">
                            <Card.Img variant="top" src={card.images.small} />
                          </div>
                          <Card.Body>
                            <Card.Title>{card.name}</Card.Title>
                            <Card.Text>
                            {card.flavorText}
                            </Card.Text>
                              <ListGroup variant="flush">
                              <ListGroup.Item>Pokémon</ListGroup.Item>
                              <ListGroup.Item>{card.set.name}</ListGroup.Item>
                              <ListGroup.Item>{card.rarity}</ListGroup.Item>
                              <ListGroup.Item>${card.cardmarket.prices.averageSellPrice}</ListGroup.Item>
                            </ListGroup>
                          </Card.Body>
                        </Card>  
                      </div> 
                    )
                  }
                )
              }
            </div>
          </div>
        </div>
      );
    }
}
  // Component that calls the CardsBase function
  class AllCards extends React.Component {
      render() {
            return (
              <div>
              <CardsBase />
              </div>
          );
      }
  }
    
  
  
  
  export default AllCards
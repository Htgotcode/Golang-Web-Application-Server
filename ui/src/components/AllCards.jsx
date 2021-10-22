import React, { useState, useEffect } from 'react';
import { Search } from 'react-bootstrap-icons';
import Card from 'react-bootstrap/Card';
import pokemon from 'pokemontcgsdk'
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import axios from 'axios';


pokemon.configure({apiKey: '8c44560c-5a0a-4e9e-828a-898992ad6345'})

function CardsBase() {
    const [pokemonCards, setPokemonCards] = useState([])
    const [marketCards, setMarketCards] = useState([])
    const [amountAvailable, setAmountAvailable] = useState(0)
    const [isLoading, setLoading] = useState(true);
    const [Name, setName] = useState('');
    
    //Call const arrow functions when parent function is called.
    useEffect(() => {
      getAllCards();
      getCards();
    }, [])

    const getAllCards = () => {
      pokemon.card.where({ q:'' })
      .then(response => {
        setPokemonCards(response.data); 
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
            setPokemonCards(response.data); 
            setLoading(false);
      })
    }


    //Match market cards with full pokemon database
    const getCards = () => {
      axios.get('/card')
        .then((response) => {
          setMarketCards(response.data);
      });
    };
  
    var uniqueMarketCards = marketCards.reduce( (acc, o) => (acc[o.name] = (acc[o.name] || 0)+1, acc), {} );
    //Fix: For each to match card count causes render to call too many times. Parse into individaul card amount.
    // for (let pokemonCard of pokemonCards) {
    //   for (let marketCard of marketCards) {
    //     if(pokemonCard.name.toLowerCase() == marketCard.name.toLowerCase() ){
           
    //     }
    //   } 
    // }
    console.log(uniqueMarketCards)

    if (isLoading){
      return <div className="App">Loading...</div>;
    } else {
      return (
        <Container className=" p-0">
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

          <Container className="mt-3">
            <Row>
                {pokemonCards.map((card) => {
                
                return (
                    <Col xs={4} md={4} lg={2} xl={2} xxl={2} sm={4} className="shadow rounded m-3" key={card._id}>
                        <Card>
                          <Card.Header as="h5">{card.name} - {amountAvailable} in Market</Card.Header>
                          <Card.Body>
                          <Card.Img variant="top" src={card.images.small}  />
                          </Card.Body>
                      </Card>  
                    </Col> 
                    )
                  }
                )
              }
            </Row>
          </Container>
        </Container>
      );
    }
}
  
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
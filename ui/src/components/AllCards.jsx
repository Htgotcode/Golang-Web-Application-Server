import React, { useState, useEffect } from 'react';
import { Search, ViewList } from 'react-bootstrap-icons';
import Card from 'react-bootstrap/Card';
import pokemon from 'pokemontcgsdk'
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import RenderCardListing from '../components/CardListing';
import axios from 'axios';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

pokemon.configure({apiKey: '8c44560c-5a0a-4e9e-828a-898992ad6345'})

function CardsBase() {
    const [CARDS, setCARDS] = useState([])
    const [CardList, setCardList] = useState([])
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
    
    const createCardListing = () => {
      if(cardListingCount > 0){
        axios.post(`/card-listing-add`, {cards: CardList }, {headers: {'Content-Type': 'application/json'}})
          .then(res => {
            console.log(res);
            console.log(res.data);
      })
      }
      
    }

    const addToCardListings = (card) => {
      setCart((CardList) => [...CardList, card]);
      setCartCount(cardListingCount + 1);
    };

    const clearCardListings = () => {
      setCart([]);
      setCardListCount(0);
      alert("Card Listing cleared.");
    };

    // const addCardListing = (card) => {
      
    //   axios.post(`/card-listing-add`, card._id, {headers: {'Content-Type': 'application/json'}})
    //   .then(res => {
    //     console.log(res);
    //     console.log(res.data);
    //   })
    // }

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
                {CARDS.map((card) => {
                return (
                    <Col xs={4} md={4} lg={2} xl={2} xxl={2} sm={4} className="shadow rounded m-3" key={card._id}>
                      
                          <Card>
                          <Card.Header as="h5">{card.name}</Card.Header>
                          <Card.Link href="#">
                          <Card.Body>
                          <Card.Img variant="top" src={card.images.small} />
                          </Card.Body>
                          </Card.Link>
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
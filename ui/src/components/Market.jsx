import React, { useState } from 'react'
import MaterialIcon from 'material-icons-react';
import Card from 'react-bootstrap/Card'
import ListGroup from 'react-bootstrap/ListGroup'
import Button from 'react-bootstrap/Button'


const CARDS = [
  {ID: 1, Brand: 'Pokemon', Image: "./images/logo192.png", Url: "", Name: "Pikachu", Set: "Celebrations", Rarity: "Holo Rare - #005/025", SellingPrice: 100, UploadedAt: "2021-10-18", CardID: "1"},
  {ID: 2, Brand: 'Pokemon', Image: "./images/logo192.png", Url: "", Name: "Riachu", Set: "Celebrations", Rarity: "Common", SellingPrice: 125, UploadedAt: "2021-10-18", CardID: "2"},
  {ID: 3, Brand: 'Pokemon', Image: "./images/logo192.png", Url: "", Name: "Gastly", Set: "Celebrations", Rarity: "Rare", SellingPrice: 35, UploadedAt: "2021-10-18", CardID: "3"},
];

function Market() {
  const [Name, setName] = useState('');
  const [foundCard, setFoundCard] = useState(CARDS);

  const filter = (e) => {
    const keyword = e.target.value;

    if (keyword !== '') {
      const RESULTS = CARDS.filter((card) => {
        return card.Name.toLowerCase().startsWith(keyword.toLowerCase());
      });
      setFoundCard(RESULTS);
    } else {
      setFoundCard(CARDS);
      // If the text field is empty, show all cards
    }

    setName(keyword);
  };
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
        size="25"
      /></span>
      
        <div className="row">
          {foundCard && foundCard.length > 0 ? (
            foundCard.map((card) => (
              <Card style={{ width: '18rem' }}>
              <Card.Img variant="top" src={card.Image} />
              <Card.Body>
                <Card.Title>{card.Name}</Card.Title>
                <Card.Text>
                  Some quick example text to build on the card title and make up the bulk of
                  the card's content.
                </Card.Text>
                  <ListGroup variant="flush">
                  <ListGroup.Item>{card.Brand}</ListGroup.Item>
                  <ListGroup.Item>{card.Set}</ListGroup.Item>
                  <ListGroup.Item>{card.Rarity}</ListGroup.Item>
                  <ListGroup.Item>R{card.SellingPrice}</ListGroup.Item>
                  <ListGroup.Item>{card.UploadedAt}</ListGroup.Item>
                  <ListGroup.Item>{card.Rarity}</ListGroup.Item>
                </ListGroup>
                <Card.Link href={card.Url}><Button variant="primary" >View Card</Button></Card.Link>
              </Card.Body>
            </Card>
              
            ))
          ) : (
            <p>Pokemon not found.</p>
          )}
        </div>
    </div>
  );
}
  



export default Market
import React, { useState }  from 'react';
import "bootstrap/dist/css/bootstrap.min.css";
import Card from 'react-bootstrap/Card'
import ListGroup from 'react-bootstrap/ListGroup'
import Button from 'react-bootstrap/Button'

const CARDS = [
    {ID: 1, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Pikachu", Set: "Celebrations", Rarity: "Holo Rare - #005/025", SellingPrice: 100, UploadedAt: "2021-10-18", CardID: "1"},
    {ID: 2, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Riachu", Set: "Celebrations", Rarity: "Common", SellingPrice: 125, UploadedAt: "2021-10-18", CardID: "2"},
    {ID: 3, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Gastly", Set: "Celebrations", Rarity: "Rare", SellingPrice: 35, UploadedAt: "2021-10-18", CardID: "3"},
];

//Card display
class RenderCards extends React.Component {
    render(){
        return(
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
        )
    }
}

export default RenderCards
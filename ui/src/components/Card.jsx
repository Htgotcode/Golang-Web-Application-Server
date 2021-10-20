import React, {useState, useEffect} from 'react';
import axios from "axios";
import {Button, Form, Container, Modal } from 'react-bootstrap'
import "bootstrap/dist/css/bootstrap.min.css";
import Card from 'react-bootstrap/Card'
import ListGroup from 'react-bootstrap/ListGroup'
import Button from 'react-bootstrap/Button'

const Card = ({cardData}) => {
  return(
    <Card style={{ width: '18rem' }} id={card.CardID}>
    <Card.Img variant="top" src={card.Image} />
    <Card.Body>
      <Card.Title>{ cardData !== undefined && cardData.name}</Card.Title>
      <Card.Text>{ cardData !== undefined && cardData.description}</Card.Text>
        <ListGroup variant="flush">
        <ListGroup.Item>{ cardData !== undefined && cardData.brand}</ListGroup.Item>
        <ListGroup.Item>{ cardData !== undefined && cardData.setName}</ListGroup.Item>
        <ListGroup.Item>{ cardData !== undefined && cardData.rarity}</ListGroup.Item>
        <ListGroup.Item>R{ cardData !== undefined && cardData.sellingPrice}</ListGroup.Item>
        <ListGroup.Item>{ cardData !== undefined && cardData.uploadedAt}</ListGroup.Item>
        <ListGroup.Item>{ cardData !== undefined && cardData.rarity}</ListGroup.Item>
      </ListGroup>
      <Card.Link><Button variant="primary" >View Card</Button></Card.Link>
    </Card.Body>
  </Card>
  );
}
/*
type Card struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Brand        string             `json:"brand" validate:"required"`
	SetName      string             `json:"set_name" validate:"required"`
	Rarity       string             `json:"rarity" validate:"required"`
	SellingPrice uint               `json:"selling_price" validate:"required"`
	UploadedAt   time.Time          `json:"uploaded_at"`
	CardID       string             `json:"card_id"`
}
*/
export default RenderCards
import React from 'react';
import axios from 'axios';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
// Component used to add card to the marketplace 
class AddCard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      description: '',
      brand: '',
      setname: '',
      rarity: '',
      sellingprice: '',
      imageurl: '',
      ownerid: '',
    };

    this.handleName = this.handleName.bind(this);
    this.handleDescription = this.handleDescription.bind(this);
    this.handleBrand = this.handleBrand.bind(this);
    this.handleSetName = this.handleSetName.bind(this);
    this.handleRarity = this.handleRarity.bind(this);
    this.handleSellingPrice = this.handleSellingPrice.bind(this);
    this.handleImageURL = this.handleImageURL.bind(this);
    this.handleOwnerID = this.handleOwnerID.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
 
  handleName = event => {
    this.setState({ name: event.target.value });
  }
  handleDescription = event => {
    this.setState({ description: event.target.value });
  }
  handleBrand = event => {
    this.setState({ brand: event.target.value });
  }
  handleSetName= event => {
    this.setState({ setname: event.target.value });
  }
  handleRarity = event => {
    this.setState({ rarity: event.target.value });
  }
  handleSellingPrice = event => {
    this.setState({ sellingprice: event.target.value });
  }
  handleImageURL = event => {
    this.setState({ imageurl: event.target.value });
  }
  handleOwnerID = event => {
    this.setState({ ownerid: event.target.value });
  }
 
  handleSubmit = event => {
    event.preventDefault();
    const card = {
      name: this.state.name,
      description: this.state.description,
      brand: this.state.brand,
      setname: this.state.setname,
      rarity: this.state.rarity,
      sellingprice: this.state.sellingprice,
      imageurl: this.state.imageurl,
      ownerid: this.state.ownerid,
    };
 
    alert = () => {
      alert("Card uploaded");
    }

    axios.post(`/card-add`, card, {headers: {'Content-Type': 'application/json'}})
      .then(res => {
        console.log(res);
        console.log(res.data);
      })
  }
 
  render() {
    return (
      <div className="container">
        <Form onSubmit={this.handleSubmit}>
          <Form.Group className="mb-3" controlId="formName">
            <Form.Label>Card Name</Form.Label>
            <Form.Control type="text" value={this.state.name} name="name" onChange={this.handleName} placeholder="Enter card name" required/>
            <Form.Text className="text-muted">
              eg. Pikachu
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formDesc">
            <Form.Label>Card Description</Form.Label>
            <Form.Control type="text" value={this.state.description} name="description" onChange={this.handleDescription} placeholder="Enter description" required/>
            <Form.Text className="text-muted">
              eg. A lightning type from the classic series.
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBrand">
            <Form.Label>Brand Name</Form.Label>
            <Form.Control type="text" value={this.state.brand} name="brand" onChange={this.handleBrand} placeholder="Enter brand name" required/>
            <Form.Text className="text-muted">
              eg. Pokemon
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formSetName">
            <Form.Label>Set Name</Form.Label>
            <Form.Control type="text" value={this.state.setname} name="setname" onChange={this.handleSetName} placeholder="Enter set name" required/>
            <Form.Text className="text-muted">
              eg. Celebrations
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formRarity">
            <Form.Label>Rarity</Form.Label>
            <Form.Control type="text"value={this.state.rarity} name="rarity" onChange={this.handleRarity} placeholder="Enter rarity" required/>
            <Form.Text className="text-muted">
              eg. Common
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formSellingPrice">
            <Form.Label>Card Selling Price</Form.Label>
            <Form.Control type="text" value={this.state.sellingprice} name="sellingprice" onChange={this.handleSellingPrice} placeholder="Enter selling price" required/>
            <Form.Text className="text-muted">
              eg. 5.26
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formImageURL">
            <Form.Label>Image URL</Form.Label>
            <Form.Control type="text" value={this.state.imageurl} name="imageurl" onChange={this.handleImageURL} placeholder="Enter image url" required/>
            <Form.Text className="text-muted">
              eg. https://sixprizes.com/wp-content/uploads/pikachu-next-destinies-nde-39.jpg
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formOwnerID">
            <Form.Label>Owner ID</Form.Label>
            <Form.Control type="text" value={this.state.ownerid} name="ownerid" onChange={this.handleOwnerID} placeholder="Enter owner ID" required/>
            <Form.Text className="text-muted">
              eg. Your personal ID MAKE THIS AUTOMATIC
            </Form.Text>
          </Form.Group>

          <Button variant="primary" type="submit" onClick={() => alert()}>
            Add
          </Button>
        </Form>
      </div>
    )
  }
}
 
export default AddCard
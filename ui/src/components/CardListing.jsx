import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';

// Function to read in card listing data and render it
function CardListing() {
    const [isLoading, setLoading] = useState(true);
    const [CardList, setCardList] = useState([]);
    
    useEffect(() => {
      readCardListing();
    }, []);
  
    const readCardListing = (cardListing) => {
      console.log(cardListing)
      axios.get('/card-listing', cardListing, {headers: {'Content-Type': 'application/json'}})
        .then((response) => {
        setCardList(response.data);
        setLoading(false);
      });
    };
    
    if (isLoading){
      return <div className="App">Loading...</div>;
    } else {
      return (
          <div className="container">
            <div className="row">
            {CardList.map((listItem) => {
                    return (
                    <div className="col-3" key={listItem.name}>
                          <Card className="m-1">
                          <div className="">
                            <Card.Img variant="top" src={listItem.imageurl} />
                          </div>
                          <Card.Body>
                            <Card.Title>{listItem.name}</Card.Title>
                            <Card.Text>
                            {listItem.description}
                            </Card.Text>
                              <ListGroup variant="flush">
                              <ListGroup.Item>{listItem.brand}</ListGroup.Item>
                              <ListGroup.Item>{listItem.setname}</ListGroup.Item>
                              <ListGroup.Item>{listItem.rarity}</ListGroup.Item>
                              <ListGroup.Item>{listItem.sellingprice}</ListGroup.Item>
                              <ListGroup.Item>{listItem.ownerid}</ListGroup.Item>
                            </ListGroup>
                            <input type ="submit" value="Buy" onClick={() => buyCard(listItem)}/>
                          </Card.Body>
                        </Card>
                        
                        {/* Card Listings Table */}
                        <h2>Card Listing</h2>
                                <table title="Card Listing" class="table table-striped">
                                <title>Card Listing</title>
                                  <thead>
                                      <tr>
                                      <th>Name</th>
                                      <th>Description</th>
                                      <th>Brand</th>
                                      <th>SetName</th>
                                      <th>Rarity</th>
                                      <th>SellingPrice</th>
                                      <th>UploadedAt</th>
                                      <th>CardID</th>
                                      </tr>
                                  </thead>
                                  <tbody>
                                      <tr>
                                      <td>{listItem.Name}</td>
                                      <td>{listItem.Description}</td>
                                      <td>{listItem.Brand}</td>
                                      <td>{listItem.SetName}</td>
                                      <td>{listItem.Rarity}</td>
                                      <td>${listItem.SellingPrice}</td>
                                      <td>{listItem.UploadedAt}</td>
                                      <td>{listItem.CardID}</td>
                                      </tr>
                                  </tbody>
                                </table>  
                      </div> 
                    )
                  })
                } 
            </div>
          </div>
      );
    }  
  }
  // Component calls the CardListing function
  class RenderCardListing extends React.Component {
      render() {
            return (
              <div>
                <CardListing />
              </div>
          );
      }
  }
      
export default RenderCardListing
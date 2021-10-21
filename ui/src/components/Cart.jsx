import React from 'react';
import axios from 'axios';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';

function Cart() {
    const [isLoading, setLoading] = useState(true);
    const [Cart, setCart] = useState([]);
    
    useEffect(() => {
        getCart();
    }, []);
  
    const getCart = () => {
      axios.get('/cart')
        .then((response) => {
        setCart(response.data);
        setLoading(false);
      });
    };

    const buyCard = (cartItem) => {
        axios.delete(`/card-remove`, cartItem._id, {headers: {'Content-Type': 'application/json'}})
            .then(res => {
        console.log(res);
        console.log(res.data);
      })
    };
    
  
    if (isLoading){
      return <div className="App">Loading...</div>;
    } else {
      return (
          <div className="container">
            <div className="row">
            {Cart.map((cartItem) => {
                    return (
                    <div className="col-3" key={cartItem._id}>
                          <Card className="m-1">
                          <div className="">
                            <Card.Img variant="top" src={cartItem.imageurl} />
                          </div>
                          <Card.Body>
                            <Card.Title>{cartItem.name}</Card.Title>
                            <Card.Text>
                            {cartItem.description}
                            </Card.Text>
                              <ListGroup variant="flush">
                              <ListGroup.Item>{cartItem.brand}</ListGroup.Item>
                              <ListGroup.Item>{cartItem.setname}</ListGroup.Item>
                              <ListGroup.Item>{cartItem.rarity}</ListGroup.Item>
                              <ListGroup.Item>{cartItem.sellingprice}</ListGroup.Item>
                              <ListGroup.Item>{cartItem.ownerid}</ListGroup.Item>
                            </ListGroup>
                            <input type ="submit" value="Buy" onClick={() => buyCard(cartItem)}/>
                          </Card.Body>
                        </Card>  
                      </div> 
                    )
                  })
                } 
            </div>
          </div>
      );
    }
    
  }
  
  class RenderCart extends React.Component {
      render() {
            return (
              <div>
              <Market />
              </div>
          );
      }
  }
    
  
export default RenderCart
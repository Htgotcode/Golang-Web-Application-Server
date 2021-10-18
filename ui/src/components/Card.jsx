import React, { useState }  from 'react';
import "bootstrap/dist/css/bootstrap.min.css";
import MaterialIcon from 'material-icons-react';

const CARDS = [
    {ID: 1, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Pikachu", Set: "Celebrations", Rarity: "Holo Rare - #005/025", SellingPrice: 100, UploadedAt: "2021-10-18", CardID: "1"},
    {ID: 2, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Riachu", Set: "Celebrations", Rarity: "Common", SellingPrice: 125, UploadedAt: "2021-10-18", CardID: "2"},
    {ID: 3, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Gastly", Set: "Celebrations", Rarity: "Rare", SellingPrice: 35, UploadedAt: "2021-10-18", CardID: "3"},
];

//Card display
class RenderCards extends React.Component {
    render(){
        return(
            <div className="card" style={{width: 20 + 'rem' }}>
                <img className="card-img-top" src={card.Image} />
                <div className="card-body">
                    <h1 className="card-title">{card.Name}</h1>
                    <h2 className="card-text">text</h2>
                    <ul className="list-group list-group-flush">
                        <li className="list-group-item">{card.Brand}</li>
                        <li className="list-group-item">{card.Set}</li>
                        <li className="list-group-item">{card.Rarity}</li>
                        <li className="list-group-item">R{card.SellingPrice}</li>
                        <li className="list-group-item">{card.UploadedAt}</li>
                        <li className="list-group-item"><a href={card.Url} target="_blank" rel="noopener noreferrer" className="card-link">Card Information</a></li>
                    </ul>
                </div>
            </div>
        )
    }
}

export default RenderCards
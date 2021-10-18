import React, { useState } from 'react'
import MaterialIcon from 'material-icons-react';

const CARDS = [
  {ID: 1, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Pikachu", Set: "Celebrations", Rarity: "Holo Rare - #005/025", SellingPrice: 100, UploadedAt: "2021-10-18", CardID: "1"},
  {ID: 2, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Riachu", Set: "Celebrations", Rarity: "Common", SellingPrice: 125, UploadedAt: "2021-10-18", CardID: "2"},
  {ID: 3, Brand: 'Pokemon', Image: "./logo192.png", Url: "", Name: "Gastly", Set: "Celebrations", Rarity: "Rare", SellingPrice: 35, UploadedAt: "2021-10-18", CardID: "3"},
];

function Market() {
  const [Name, setName] = useState('');
  const [foundCard, setFoundCard] = useState(CARDS);

  const filter = (e) => {
    const keyword = e.target.value;

    if (keyword !== '') {
      const results = CARDS.filter((card) => {
        return card.Name.toLowerCase().startsWith(keyword.toLowerCase());
      });
      setFoundCard(results);
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
            ))
          ) : (
            <p>Pokemon not found.</p>
          )}
        </div>
    </div>
  );
}
  



export default Market
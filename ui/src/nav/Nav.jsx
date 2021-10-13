import React from 'react';
 
import { NavLink } from 'react-router-dom';
 
const Navigation = () => {
    return (
       <div>
          <NavLink to="/">Home</NavLink>
          <NavLink to="/card-sell">Sell Card</NavLink>
       </div>
    );
}
 
export default Navigation;
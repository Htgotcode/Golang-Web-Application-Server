import React, { Component } from 'react';
import { BrowserRouter, Link, NavLink, Route, Switch } from 'react-router-dom';
class PurchaserHistory extends React.Component {
    constructor(props) {    super(props);    this.state = {      value: null,    };  }
    render() {
      return (
        <button className="View Purchase History" onClick={() => console.log('click')}>
          {this.props.value}
        </button>
      );
    }
  }

  export default PurchaserHistory;
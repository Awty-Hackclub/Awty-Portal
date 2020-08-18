import React from 'react'
import 'react-bootstrap';

import './../styles/navbar.css'
import { FaBeer } from 'react-icons/fa';
// Doesn't work ^

function NavBar() {
    return (
        <div className="navbar">
            <h1 class="nav-header" href="/">Awty Portal </h1>
            <input className="searchBar"></input>
            <div className="buttons">
                <button className="home"></button>
                <button className="add"></button>
                <button className="user"></button>
            </div>
            <h3> Lets go for a <FaBeer />? </h3>
        </div>
    )
}

export default NavBar
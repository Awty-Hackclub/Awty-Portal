import React from 'react'
import { Link } from 'react-router-dom';
import './../styles/navbar.css'
//import 'react-icons';

function NavBar() {
    return (
        <div className="navbar">
            <Link class="nav-header">Awty Portal </Link>
            <input className="searchBar"></input> 
        </div>
    )
}

export default NavBar
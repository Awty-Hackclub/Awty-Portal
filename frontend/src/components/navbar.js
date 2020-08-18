import React from 'react'
import './../styles/navbar.css'
//import 'react-icons';
// Doesn't work ^

function NavBar() {
    return (
        <div className="navbar">
            <h1 class="nav-header" href="/">Awty Portal </h1>
            <input className="searchBar" onKeyDown={(e) => handleEnter(e)}></input>
            <div className="buttons">
                <button className="home"></button>
                <button className="add"></button>
                <button className="user"></button>
            </div>
                
        </div>
    )
}

export default NavBar
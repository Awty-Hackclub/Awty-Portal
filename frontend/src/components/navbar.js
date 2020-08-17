import React from 'react'
import './../styles/navbar.css'
//import 'react-icons';
// Doesn't work ^

function NavBar() {
    const handleEnter=(event) => {
        if (event.keyCode === 13) {
            // Handle enter
            console.log('enter')
        }
    }
    return (
        <div className="navbar">
            <h1 class="nav-header" href="/">Awty Portal </h1>
            <input className="searchBar" onKeyDown={(e) => handleEnter(e)}></input>
            <div className="buttons">
                <button className="home" onClick={() => console.log("home button clicked")}></button>
                <button className="add" onClick={() => console.log("home button clicked")}></button>
                <button className="user" onClick={() => console.log("home button clicked")}></button>
            </div>
                
        </div>
    )
}

export default NavBar
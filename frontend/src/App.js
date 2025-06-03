import React from "react"
import Navigation from "./Components/Navigation"
import "./App.css"

const App = () => {
  return (
    <div className="App">
      <h1>Endpoint Monitoring Service</h1>
      <Navigation />
      <a href="https://github.com/chensxb97" className="github-link" target="_blank" rel="noopener noreferrer">
        <img src="https://upload.wikimedia.org/wikipedia/commons/9/91/Octicons-mark-github.svg" alt="GitHub Logo" className="github-logo" />
        <span>Check out my github for more projects!</span>
      </a>
    </div >
  )
}

export default App
import './App.css';
import React, { useState } from 'react';
import axios from 'axios';

function App() {

  const [data, setData] = useState({
    url: "",
    expiration: "",
  });

  const [response, setResponse] = useState(null);

  const handleSubmit = (e) => {
    e.preventDefault();

    const payload = {
      "url": data.url,
      "expiration": data.expiration
    }

    axios.post(`http://localhost:6666/encode`, payload)
      .then(response => {
        console.log(response);
        setResponse(response);
      })
      .catch(err => {
        console.log(err);
      })
  }

  return (
    <div className="App">
      <div className="content">
        <div className="content-inside">
          <header className="header">
       	    <h1>Short URL</h1>
          </header>
          <div className="url">
            <form>
              <label>URL:</label>
              <input type="text" />
              <label>Expire in:</label>
              <select>
                <option value="never">Never</option>
                <option value="5">5 min</option>
                <option value="10">10 min</option>
                <option value="15">15 min</option>
                <option value="30">30 min</option>
                <option value="hour">1 Hour</option>
              </select>
            </form>
          </div>
        </div>
      </div>
      <footer className="footer">
        <h2>Footer</h2>
      </footer>
    </div>
  );
}

export default App;

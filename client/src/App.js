import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';


import Home from './Pages/Home/Home.js'
import Info from './Pages/Info/Info.js'

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/" exact>
          <Home />
        </Route>
        <Route path="/:id/info">
          <Info />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;

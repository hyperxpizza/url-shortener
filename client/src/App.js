import './App.css';
import React, { useState } from 'react';
import axios from 'axios';

import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    maxWidth: 1000,
  },
  rootForm: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: '25ch',
    },
  },
}));


const expirationOptions = [
  {
    value: "never",
    label: "Never",
  },
  {
    value: "5",
    label: "5 min.",
  },
  {
    value: "10",
    label: "10 min.",
  },
  {
    value: "15",
    label: "15 min."
  },
  {
    value: "30",
    label: "30 min."
  },
  {
    value: "hour",
    label: "1 hour",
  }
];

function App() {

  const classes = useStyles();
  const [payload, setPayload] = useState({
    "url": "",
    "expiration": "5",
  });

  const handleChange = (e) => {
    const {id, value} = e.target
    setPayload(prevPayload => ({
      ...prevPayload,
      [id]: value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(payload);
    axios.post('http://localhost:8888/encode', {
      "url": payload.url,
      "expiration": payload.expiration
    })
    .then((response) => {
      console.log(response);
    }, (error) => {
      console.log(error);
    });
  }

  return (
    <div className="App">
      <Container className={classes.root}>
        <Typography variant="h2" component="h2" align="center">
          Short URL
        </Typography>
        <Container className={classes.root}>
          <form className={classes.rootForm} noValidate>
            <TextField 
              label="URL"
              autoFocus
              id="url"
              value={payload.url}
              onChange={handleChange}
            />
            <TextField 
              id="expiration"
              select
              label="Expire at"
              value={payload.expiration}
              onChange={handleChange}
              >
                {expirationOptions.map((option) => {
                  <MenuItem key={option.value} value={option.value}>
                    {option.label}
                  </MenuItem>
                })}
            </TextField>
            <Button
              type="submit"
              color="primary"
              onClick={handleSubmit}
            >
            Short it!
            </Button>
          </form>
        </Container>
      </Container>
    </div>
  );
}

export default App;

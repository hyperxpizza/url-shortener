import React, { useState } from 'react';
import axios from 'axios';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  root: {
    marginTop: '50px',
    flexGrow: 1,
  },
  main: {
    marginTop: theme.spacing(8),
    marginBottom: theme.spacing(2),
  },
  rootForm: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: '25ch',
    },
    alignItems: 'center',
  },
  mainContainer: {
    paddingTop: '10px',
  },
  select: {
    paddingTop: '23.5px',
  }
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

const Home = () => {
    const classes = useStyles();
  const [url, setUrl] = useState("");
  const [expiration, setExpiration] = useState();
  const [resp, setResp] = useState();

  const handleChangeURL = (e) => {
    setUrl(e.target.value);
  }

  const handleChangeExpiration = (e) => {
    setExpiration(e.target.value);
  }

  const handleSubmit = (e) => {
    e.preventDefault();

    axios.post('http://localhost:8888/encode', {
      "url": url,
      "expiration": expiration
    })
    .then((response) => {
      console.log(response);
      let url = "http://localhost:8888/" + response.data.code;
      setResp(url);
      setUrl("");
      setExpiration();
    }, (error) => {
      console.log(error);
    });

  }

  return (
    <div className={classes.root}>
      <Container>
        <Typography variant="h2" component="h2" align="center">
          Short URL
        </Typography>
        <Grid container className={classes.root} spacing={1} justify="center" alignItems="center">
          <form className={classes.rootForm} validate>
            <TextField 
              label="URL"
              autoFocus
              id="url"
              value={url}
              onChange={handleChangeURL}
              required
            />
            <Select
              id="expiration"
              value={expiration}
              onChange={handleChangeExpiration}
              className={classes.select}
            >
              {expirationOptions.map(option => (
                <MenuItem value={option.value} key={option.value}>{option.label}</MenuItem>
              ))}
            </Select>
            <Button
              type="submit"
              color="primary"
              onClick={handleSubmit}
            >
            Short it!
            </Button>
          </form>
        </Grid>
        <Container className={classes.mainContainer}>
          <Typography variant="h4" component="h4" align="center">
            <a href={resp} target="_blank">{resp}</a>
          </Typography>
        </Container>
      </Container>
    </div>
  );
}

export default Home;
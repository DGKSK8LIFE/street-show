import React from 'react';
import './App.css';
import { TextField, makeStyles, Button } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: '25ch',
    },
  },
  button: {
    '& > *': {
      margin: theme.spacing(1),
    },
  },
}));


function App() {
  const classes = useStyles()

  return (
    <div style={{ display: `flex`, justifyContent: `center`, alignItems: `center` }}>
      <header className="App-header">
        <form>
          <TextField className={classes.root} id="outlined-search" label="Search by name" varient="filled" type="search" />
          <Button className={classes.button} variant="outlined" color="default" type="submit" size="small">
            Search
          </Button>
        </form>
      </header>
    </div>
  );
}

export default App;

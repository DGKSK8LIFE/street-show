import React from 'react';
import './App.css';
import { TextField, makeStyles } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: '25ch',
    },
  },
}));


function App() {
  const classes = useStyles()

  return (
    <div className={classes.root} style={{ display: `flex`, justifyContent: `center`, alignItems: `center` }}>
      <header className="App-header">
        <form>
          <TextField id="outlined-search" label="Search by name" varient="filled" type="search" />
          <input class="btn btn-secondary btn-lg" type="submit" value="search"></input>
        </form>
      </header>
    </div>
  );
}

export default App;

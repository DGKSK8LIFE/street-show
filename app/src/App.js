import React from 'react';
import './App.css';
import { TextField,  } from '@material-ui/core';
import { useStyles, ColorButton } from './Styles';

function App() {
  const classes = useStyles()

  return (
    <div style={{ display: `flex`, justifyContent: `center`, alignItems: `center` }}>
      <header className="App-header">
        <form>
          <TextField className={classes.root} id="outlined-search" label="Search for Buskers near you" varient="filled" type="search" size="large" />
          <ColorButton variant="contained" color="primary" className={classes.margin}>
            Search
          </ColorButton>
        </form>
      </header>
    </div>
  );
}

export default App;

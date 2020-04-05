import React from 'react';
import './App.css';
import { TextField, makeStyles, Button, withStyles } from '@material-ui/core';
import { purple } from '@material-ui/core/colors';

const useStyles = makeStyles((theme) => ({
  root: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: '25ch',
    },
  },
  margin: {
    margin: theme.spacing(),
  },
}));

const ColorButton = withStyles((theme) => ({
  root: {
    color: theme.palette.getContrastText(purple[500]),
    backgroundColor: purple[500],
    '&:hover': {
      backgroundColor: purple[700],
    },
  },
}))(Button);

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

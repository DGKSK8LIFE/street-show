import React from 'react';
import { TextField, } from '@material-ui/core';
import { useStyles, ColorButton } from './Styles';

function LandingPage() {
    const classes = useStyles()

    return (
        <div style={{ display: `flex`, justifyContent: `center` }}>
            <header className="App-header">
                <form>
                    <TextField className={classes.root} id="outlined-search" label="Search for Buskers near you" varient="filled" type="search" size="large" />
                    <ColorButton variant="contained" color="primary" className={classes.margin} type="submit">
                        Search
          </ColorButton>
                </form>
            </header>
        </div>
    );
}

export default LandingPage;
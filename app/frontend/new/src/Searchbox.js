import React from 'react';
import { TextField, } from '@material-ui/core';
import { useStyles, ColorButton } from './Styles';
import './App.css';
import SimpleBottomNavigation from './Footer';
export default function Searchbox() {
    const classes = useStyles()

    return (

        <div class="searchbox" >

            <header className="App-header">
                <form>
                    <TextField className={classes.root} id="outlined-search" placeholder="Search for Buskers near you" varient="filled" type="search" size="large" />
                    <ColorButton variant="contained" color="primary" className={classes.margin} type="submit" id="button">
                        Search
                    </ColorButton>
                </form>
            </header>
        </div>

    );
}
import React from 'react';
import EventsExample from './map'
import { TextField, } from '@material-ui/core';
import { useStyles, ColorButton } from './Styles';
import './App.css';
import SimpleBottomNavigation from './Footer';
import Searchbox from './Searchbox';

export default function LandingPage() {
    const classes = useStyles()

    return (
        <div class="wrapper">
            <EventsExample />
            <Searchbox />

        </div>
    );
}
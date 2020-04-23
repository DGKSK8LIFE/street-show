import React, { Component } from 'react';
import EventsExample from './map'
// import { TextField, } from '@material-ui/core';
// import { useStyles, ColorButton } from './Styles';
import './App.css';
// import SimpleBottomNavigation from './Footer';
import Searchbox from './Searchbox';
// import SearchBar from 'material-ui-search-bar'
// import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
class LandingPage extends Component {
    constructor() {
        super()
        this.state = {
            searchfield: ' ',
            buskers: [],
        }
    }
    onSearchChange = (event) => {
        this.setState({ searchfield: event });
    }
    render() {
        return (
            <div class="wrapper">
                <EventsExample />
                <Searchbox searchChange={this.onSearchChange} />
            </div>
        );
    }
}
export default LandingPage;
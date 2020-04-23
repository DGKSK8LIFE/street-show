import React from 'react'
import SearchBar from 'material-ui-search-bar'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'

const Searchbox = ({ searchfield, searchChange }) => {
    return (

        <MuiThemeProvider>
            <SearchBar
                // onClick={() => console.log("clicked")}
                onChange={searchChange}
                onRequestSearch={() => console.log('onRequestSearch')}
                style={{
                    maxWidth: 400,
                    position: 'absolute',
                    zIndex: 1000,
                    left: 0,
                    right: 0,
                    marginLeft: 'auto',
                    marginRight: 'auto',
                    top: 10
                }}

            />

        </MuiThemeProvider>
    );
}
export default Searchbox;





// import React from 'react';
// import { TextField, } from '@material-ui/core';
// import { useStyles, ColorButton } from './Styles';
// import './App.css';
// import SimpleBottomNavigation from './Footer';
// export default function Searchbox() {
//     const classes = useStyles()

//     return (

//         <div class="searchbox" >

//             <header className="App-header">
//                 <form>
//                     <TextField className={classes.root} id="outlined-search" placeholder="Search for Buskers near you" varient="filled" type="search" size="large" />
//                     <ColorButton variant="contained" color="primary" className={classes.margin} type="submit" id="button">
//                         Search
//                     </ColorButton>
//                 </form>
//             </header>
//         </div>

//     );
// }



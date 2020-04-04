import React, { Component } from 'react';
import ReactDOM from 'react-dom';

class App extends Component {
	render() {
		return (
			<div class="root">
				<h1>HELLO WORLD!</h1>
			</div>
		)	
	}	
}

ReactDOM.render (
	App,
	document.getElementById('root')
);


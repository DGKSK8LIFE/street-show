import { Map, Marker, Popup, TileLayer } from 'react-leaflet';
import './LandingPage.css';
import React, { createRef, Component } from 'react';

// export default function LandingPage() {
//     return (
//         <Map center={[45.4, -75.7]} zoom={12}>
//             <TileLayer
//                 url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
//                 attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
//             />
//         </Map>
//     );
// }
type State = {
    hasLocation: boolean,
    latlng: {
        lat: number,
        lng: number,
    },
}

export default class EventsExample extends Component<{}, State> {
    state = {
        hasLocation: false,
        latlng: {
            lat: 51.505,
            lng: -0.09,
        },
    }

    mapRef = createRef()

    handleClick = () => {
        const map = this.mapRef.current
        if (map != null) {
            map.leafletElement.locate()
        }
    }

    handleLocationFound = (e: Object) => {
        this.setState({
            hasLocation: true,
            latlng: e.latlng,
        })
    }

    render() {
        const marker = this.state.hasLocation ? (
            <Marker position={this.state.latlng}>
                <Popup>You are here</Popup>
            </Marker>
        ) : null

        return (
            <Map
                center={this.state.latlng}
                length={4}
                onClick={this.handleClick}
                onLocationfound={this.handleLocationFound}
                ref={this.mapRef}
                zoom={13}>
                <TileLayer
                    attribution='&amp;copy <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                />
                {marker}
            </Map>
        )
    }
}
// import React, { Component } from 'react';
// import Modal from 'react-awesome-modal';
// import './Modal.css'
// export default class MODAL extends Component {
//     constructor(props) {
//         super(props);
//         this.state = {
//             visible: false
//         }
//     }
//     openModal() {
//         this.setState({
//             visible: true
//         });
//     }

//     closeModal() {
//         this.setState({
//             visible: false
//         });
//     }

//     render() {
//         return (
//             <div
//                 style={{
//                     position: 'absolute',
//                     zIndex: -1,
//                     bottom: 10,
//                     marginLeft: 'auto',
//                     marginRight: 'auto',
//                     left: 0,
//                     right: 0,
//                     width: 'fit-content',
//                 }}
//             >
//                 <input type="button" value="UP" onClick={() => this.openModal()} />
//                 <Modal
//                     visible={this.state.visible}
//                     width="100%"
//                     height="80%"
//                     effect="fadeInUp"
//                     onClickAway={() => this.closeModal()}
//                 >
//                     <div>
//                         <h1>Title</h1>
//                         <p>Some Contents</p>
//                         <a href="javascript:void(0);" onClick={() => this.closeModal()}>Close</a>
//                     </div>
//                 </Modal>
//             </div>
//         );
//     }
// }
import React, { useState, useCallback } from 'react';
import { render } from 'react-dom';
import useModal from 'react-hooks-use-modal';

const MODAL = () => {
    const [Modal, open, close, isOpen] = useModal('root', {
        preventScroll: true
    });
    return (
        <div
            style={{
                position: 'absolute',
                zIndex: 800,
                bottom: 10,
                marginLeft: 'auto',
                marginRight: 'auto',
                left: 0,
                right: 0,
                width: 'fit-content',
            }}
        >
            <p>Modal is Open? {isOpen ? 'Yes' : 'No'}</p>
            <button onClick={open}>UP</button>
            <Modal>
                <div>
                    <h1>Title</h1>
                    <p>This is a customizable modal.</p>
                    <button onClick={close}>CLOSE</button>
                </div>
            </Modal>
        </div>
    );
};
export default MODAL;
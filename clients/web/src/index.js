import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

import * as battleship from './battleship';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

let { gameID, playerID } = await battleship.newGame();
console.log('gameID', gameID);
console.log('playerID', playerID);

console.log('exporting gameID and playerID to window');
window.gameID = gameID;
window.playerID = playerID;

// We only do this here to get us to a place of being able to place ships. In
// the real world, we would wait for another player to join, and would need to
// be notified when we're able to proceed placing ships.
let { playerID: playerTwoID } = await battleship.joinGame(gameID);
console.log('playerTwoID', playerTwoID);

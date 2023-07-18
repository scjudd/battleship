import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {NewGameRequest} from './proto/battleship_pb.js';
import {BattleshipClient} from './proto/battleship_grpc_web_pb.js';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

function newGame() {
	console.log("starting a new game");

	let client = new BattleshipClient("http://localhost:8080")
	let request = new NewGameRequest();
	client.newGame(request, {}, function(err, response) {
		if (err) {
			console.log(err);
			return
		}

		window.gameID = response.getGameid();
		window.playerID = response.getPlayerid();

		console.log("Game ID", window.gameID);
		console.log("Player ID", window.playerID);
	});
}

newGame();

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

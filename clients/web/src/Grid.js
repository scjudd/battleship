import "./Grid.css";

import {NewGameRequest} from './proto/battleship_pb.js';
import {BattleshipClient} from './proto/battleship_grpc_web_pb.js';

function fire(x, y) {
	console.log("Firing at (" + x + "," + y + ")!");

	let client = new BattleshipClient("http://localhost:8080")
	let request = new NewGameRequest();
	client.newGame(request, {}, function(err, response) {
		if (err) {
			console.log(err);
		} else {
			console.log("Game ID", response.getGameid());
			console.log("Player ID", response.getPlayerid());
		}
	});
}

function Grid() {
	let rows = [];
	for (let y = 0; y < 10; y++) {
		let row = []
		for (let x = 0; x < 10; x++) {
			row.push(<div className="Grid-cell" key={x + "," + y} onClick={() => fire(x, y)}>{x + "," + y}</div>)
		}
		rows.push(<div className="Grid-row" key={y}>{row}</div>)
	}

	return (
		<div>{rows}</div>
	);
}

export default Grid;

import "./Grid.css";
import * as battleship from './battleship'

function fire(x, y) {
	console.log("Firing at (" + x + "," + y + ")!");
}

async function placeShip(gameID, playerID, name, x, y, position) {
	console.log("placing a " + name + " at (" + x + "," + y + ")")
	await battleship.placeShip(gameID, playerID, name, x, y, position)
}

function Grid() {
	let rows = [];
	for (let y = 0; y < 10; y++) {
		let row = []
		for (let x = 0; x < 10; x++) {
			row.push(<div className="Grid-cell" key={x + "," + y} onClick={() => placeShip(window.gameID, window.playerID, "Battleship", x, y, false)}>{x + "," + y}</div>)
		}
		rows.push(<div className="Grid-row" key={y}>{row}</div>)
	}

	return (
		<div>{rows}</div>
	);
}

export default Grid;

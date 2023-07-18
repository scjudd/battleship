import "./Grid.css";

function fire(x, y) {
	console.log("Firing at (" + x + "," + y + ")!");
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

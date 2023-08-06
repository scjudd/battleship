import { useState } from 'react';

import Grid from '../Grid';
import Ship from '../Ship';

export default function NewGame() {
  let [dragging, setDragging] = useState(null);

  return (
    <div className="flex justify-center">
      <Grid dragging={dragging} />
      <div className="flex flex-wrap content-start">
        <Ship kind="Carrier" setDragging={setDragging} />
        <Ship kind="Battleship" setDragging={setDragging} />
        <Ship kind="Cruiser" setDragging={setDragging} />
        <Ship kind="Submarine" setDragging={setDragging} />
        <Ship kind="Destroyer" setDragging={setDragging} />
      </div>
    </div>
  );
}

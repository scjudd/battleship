import { useState } from 'react';

import Grid from '../Grid';
import Ship from '../Ship';

export default function NewGame() {
  let [dragging, setDragging] = useState(null);

  return (
    <div className="flex justify-center">
      <Grid dragging={dragging} />
      <div className="flex justify-center">
        <Ship kind="Cruiser" setDragging={setDragging} />
      </div>
    </div>
  );
}

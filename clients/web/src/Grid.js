import { useState } from 'react';

import './Grid.css';

function Cell({ isOver, onDragOver }) {
  const [duration] = useState(`${Math.random() + 3}s`);
  const [delay] = useState(`${Math.random() - 4}s`);

  return (
    <div
      onDragOver={onDragOver}
      className={
        'inline-block w-[24px] h-[24px] m-0.5 bg-gray-600 border border-gray-400 rounded-sm' +
        (isOver ? ' bg-red-500' : '')
      }
      style={{ animation: `${duration} ease ${delay} infinite grid-wave` }}
    />
  );
}

export default function Grid({ dragging }) {
  let [over, setOver] = useState(null);

  function handleDragOver({ x: overX, y: overY }) {
    if (dragging === null) {
      return;
    }

    const newOver = {
      startX: overX - dragging.dragX,
      startY: overY - dragging.dragY,
      length: dragging.length,
      vertical: dragging.vertical,
    };

    if (
      over === null ||
      newOver.startX !== over.startX ||
      newOver.startY !== over.startY ||
      newOver.length !== over.length ||
      newOver.vertical !== over.vertical
    ) {
      setOver(newOver);
    }
  }

  function isOver({ x, y }) {
    if (over === null) {
      return false;
    }

    if (over.vertical === false) {
      return (
        y === over.startY && x >= over.startX && x < over.startX + over.length
      );
    }

    throw new Error('unimplemented: not currently handling vertical ships');
  }

  let rows = [];
  for (let y = 0; y < 10; y++) {
    let cells = [];
    for (let x = 0; x < 10; x++) {
      cells.push(
        <Cell
          key={'x-' + x}
          isOver={isOver({ x, y })}
          onDragOver={() => handleDragOver({ x, y })}
        />,
      );
    }
    rows.push(
      <div key={'y-' + y} className="flex">
        {cells}
      </div>,
    );
  }

  return (
    <div className="p-2 bg-gray-900 rounded border border-gray-600">{rows}</div>
  );
}

import { useState } from 'react';

const cellSize = 26;

export default function Ship({ kind, setDragging }) {
  let [isDragging, setIsDragging] = useState(false);

  function handleDragStart(e) {
    // If the ship were broken into cellSize squares, compute which one our
    // mouse would be in during a drag. When the mouse is over a grid cell
    // during a drag event, these dimensions will inform how many cells up,
    // right, down or left need to be handled.
    const {offsetX, offsetY} = e.nativeEvent;
    const [dragX, dragY] = [Math.floor(offsetX / cellSize), Math.floor(offsetY / cellSize)];
    setDragging({length: 5, vertical: false, dragX, dragY});
    setIsDragging(true);
  }

  function handleDragEnd(e) {
    setDragging(null);
    setIsDragging(false);
  }

  function handleDrop(e) {
    setIsDragging(false);
  }

  const base = 'inline w-[136px] h-[24px] m-0.5 bg-sky-700 border rounded-sm cursor-grab active:cursor-grabbing';
  const dragging = 'opacity-25 border-dotted';

  function classNames() {
    if (isDragging) {
      return [base, dragging].join(' ');
    }
    return base;
  }

  return (
    <div
      onDragStart={handleDragStart}
      onDragEnd={handleDragEnd}
      onDrop={handleDrop}
      draggable='true'
      className={classNames()}
    />
  );
}

import { useState, useEffect } from 'react';

const cellSize = 26;

const Orientation = Object.freeze({
  Horizontal: Symbol('horizontal'),
  Vertical: Symbol('vertical'),
});

export default function Ship({ kind, setDragging }) {
  const [orientation, setOrientation] = useState(Orientation.Horizontal);
  const [cursor, setCursor] = useState({ x: 0, y: 0, dragging: false });

  useEffect(() => {
    const setPositionFromEvent = (e) => {
      if (cursor.dragging) {
        setCursor({ x: e.clientX, y: e.clientY, dragging: true });
      }
    };

    const stopDragging = () => {
      setCursor({ x: cursor.x, y: cursor.y, dragging: false });
    };

    window.addEventListener('mousemove', setPositionFromEvent);
    window.addEventListener('mouseup', stopDragging);

    return () => {
      window.removeEventListener('mousemove', setPositionFromEvent);
      window.removeEventListener('mouseup', stopDragging);
    };
  }, [cursor]);

  let length = null;
  if (kind === 'Carrier') {
    length = 5;
  } else if (kind === 'Battleship') {
    length = 4;
  } else if (kind === 'Cruiser' || kind === 'Submarine') {
    length = 3;
  } else if (kind === 'Destroyer') {
    length = 2;
  } else {
    throw new Error(`Unexpected ship kind: ${kind}`);
  }

  let dimensions = null;
  if (length === 1) {
    dimensions = { w: 24, h: 24 };
  } else if (length === 2 && orientation === Orientation.Horizontal) {
    dimensions = { w: 52, h: 24 };
  } else if (length === 2 && orientation === Orientation.Vertical) {
    dimensions = { w: 24, h: 52 };
  } else if (length === 3 && orientation === Orientation.Horizontal) {
    dimensions = { w: 80, h: 24 };
  } else if (length === 3 && orientation === Orientation.Vertical) {
    dimensions = { w: 24, h: 80 };
  } else if (length === 4 && orientation === Orientation.Horizontal) {
    dimensions = { w: 108, h: 24 };
  } else if (length === 4 && orientation === Orientation.Vertical) {
    dimensions = { w: 24, h: 108 };
  } else if (length === 5 && orientation === Orientation.Horizontal) {
    dimensions = { w: 136, h: 24 };
  } else if (length === 5 && orientation === Orientation.Vertical) {
    dimensions = { w: 24, h: 136 };
  } else {
    throw new Error(`Unexpected ship length: ${length}`);
  }

  // These class strings must be static in order for tailwind to find them.
  const dimensionClasses = (length, orientation) => {
    if (length === 1) {
      return 'w-[24px] h-[24px]';
    } else if (length === 2 && orientation === Orientation.Horizontal) {
      return 'w-[52px] h-[24px]';
    } else if (length === 2 && orientation === Orientation.Vertical) {
      return 'w-[24px] h-[52px]';
    } else if (length === 3 && orientation === Orientation.Horizontal) {
      return 'w-[80px] h-[24px]';
    } else if (length === 3 && orientation === Orientation.Vertical) {
      return 'w-[24px] h-[80px]';
    } else if (length === 4 && orientation === Orientation.Horizontal) {
      return 'w-[108px] h-[24px]';
    } else if (length === 4 && orientation === Orientation.Vertical) {
      return 'w-[24px] h-[108px]';
    } else if (length === 5 && orientation === Orientation.Horizontal) {
      return 'w-[136px] h-[24px]';
    } else if (length === 5 && orientation === Orientation.Vertical) {
      return 'w-[24px] h-[136px]';
    } else {
      throw new Error(`Unexpected ship length: ${length}`);
    }
  };

  const baseClasses = [
    'inline',
    dimensionClasses(length, orientation),
    'm-0.5',
    'bg-sky-700',
    'border',
    'rounded-sm',
    'cursor-grab',
  ].join(' ');

  const activeClasses = [baseClasses, 'cursor-grabbing'].join(' ');

  const ghostClasses = [baseClasses, 'border-dotted', 'opacity-25'].join(' ');

  const handleDragStart = (e) => {
    e.preventDefault();
  };

  const handleMouseDown = (e) => {
    if (e.button !== 0) return;
    e.preventDefault();
    setCursor({ x: e.clientX, y: e.clientY, dragging: true });
  };

  const handleMouseUp = (e) => {
    e.preventDefault();
    setCursor({ x: 0, y: 0, dragging: false });
  };

  return (
    <>
      <div
        onDragStart={handleDragStart}
        onMouseDown={handleMouseDown}
        onMouseUp={handleMouseUp}
        className={cursor.dragging ? ghostClasses : baseClasses}
      />
      {cursor.dragging && (
        <div
          onDragStart={handleDragStart}
          className={activeClasses}
          style={{
            position: 'absolute',
            top: cursor.y - dimensions.h / 2,
            left: cursor.x - dimensions.w / 2,
          }}
        />
      )}
    </>
  );
}

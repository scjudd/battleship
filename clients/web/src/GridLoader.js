const grid = createAnimation(
  'GridLoader',
  '0% {transform: scale(1)} 50% {transform: scale(0.5); opacity: 0.7} 100 {transform: scale(1); opacity: 1}',
  'grid',
);

const random = (top) => Math.random() * top;

function createAnimation(loaderName, frames, suffix) {
  const animationName = `react-spinners-${loaderName}-${suffix}`;

  if (typeof window == 'undefined' || !window.document) {
    return animationName;
  }

  const styleEl = document.createElement('style');
  document.head.appendChild(styleEl);
  const styleSheet = styleEl.sheet;

  const keyFrames = `
    @keyframes ${animationName} {
      ${frames}
    }
  `;

  if (styleSheet) {
    styleSheet.insertRule(keyFrames, 0);
  }

  return animationName;
}

function GridLoader({ loading = true }) {
  const wrapper = {
    width: '480px',
    fontSize: 0,
    display: 'inline-block',
  };

  const speedMultiplier = 1;

  const style = (rand) => {
    return {
      display: 'inline-block',
      backgroundColor: 'rgba(0, 0, 0, 0.25)',
      width: '40px',
      height: '40px',
      margin: '4px',
      animationFillMode: 'both',
      animation: `${grid} ${(rand / 100 + 0.6) / speedMultiplier}s ${
        rand / 100 - 0.2
      }s infinite ease`,
    };
  };

  if (!loading) {
    return null;
  }

  let cells = [];
  for (let i = 0; i < 100; i++) {
    cells.push(<span style={style(random(100))} />)
  }

  return (
    <span style={wrapper}>
      {cells}
    </span>
  );
}

export default GridLoader;

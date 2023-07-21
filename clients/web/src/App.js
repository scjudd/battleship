import Grid from './Grid';
import GridLoader from './GridLoader';

function App() {
  return (
    <div>
      <header>
        <GridLoader />
        <Grid />
        <a
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;

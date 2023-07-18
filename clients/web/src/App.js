import './App.css';
import Grid from './Grid';
import GridLoader from './GridLoader';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <GridLoader />
        <Grid />
        <a
          className="App-link"
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

import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';

import Layout from './pages/Layout';
import NewGame from './pages/NewGame';
import JoinGame from './pages/JoinGame';
import NoPage from './pages/NoPage';

import GameMenu from './GameMenu';
import Grid from './Grid';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route exact path="/" element={<Navigate to="/new" />} />
          <Route path="/new" element={<NewGame />} />
          <Route path="/join" element={<JoinGame />} />
          <Route path="*" element={<NoPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;

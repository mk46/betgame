import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './components/LoginPage';
import BetGamePage from './components/BetgamePage';
import GamePage from './components/GamePage';
import PastBetsPage from './components/PastBetsPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/betgame" element={<BetGamePage />} />
        <Route path="/game" element={<GamePage />} />
        <Route path="/past-bets" element={<PastBetsPage />} />
      </Routes>
    </Router>
  );
}

export default App;

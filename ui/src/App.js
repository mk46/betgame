import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './components/LoginPage';
import BetGamePage from './components/BetgamePage';
import GamePage from './components/GamePage';
import PastBetsPage from './components/PastBetsPage';
import AdminView from './components/AdminView';
import AdminLogin from './components/AdminLogin';
import AddGame from './components/AddGame';
import ListGames from './components/ListGames';
import ListUser from './components/ListUser';
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/betgame" element={<BetGamePage />} />
        <Route path="/game" element={<GamePage />} />
        <Route path="/past-bets" element={<PastBetsPage />} />
        <Route path="/admin" element={<AdminView />} />
        <Route path="/login" element={<AdminLogin />} />
        <Route path="/admin/addgame" element={<AddGame />} />
        <Route path="/admin/listgames" element={<ListGames />} />
        <Route path="/admin/getusers" element={<ListUser />} />
      </Routes>
    </Router>
  );
}

export default App;

import React from 'react';
import { Box, Button, Heading } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

function BetGamePage() {
  const navigate = useNavigate();

  return (
    <Box p="4">
      <Heading mb="6">Bet Game</Heading>
      <Button colorScheme="teal" onClick={() => navigate('/game')}>
        Start a New Game
      </Button>
      <Button colorScheme="blue" onClick={() => navigate('/past-bets')} ml="4">
        View Past Bets
      </Button>
    </Box>
  );
}

export default BetGamePage;

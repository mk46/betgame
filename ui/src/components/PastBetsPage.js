import React, { useState, useEffect } from 'react';
import { Box, Button, Heading, List, ListItem, Text, Divider } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

function PastBetsPage() {
  const navigate = useNavigate();
  const [pastBets, setPastBets] = useState([]);

  useEffect(() => {
    const getPastBets = async () => {
      const uid = localStorage.getItem('uid');
      const response = await fetch(`http://localhost:8080/getpastbets/${uid}`);
      const data = await response.json();
      setPastBets(data.data);
    }

    getPastBets();
  }, []);

  return (
    <Box
      p="6"
      maxW="lg"
      mx="auto"
      mt="10"
      bg="white"
      boxShadow="md"
      borderRadius="md"
      overflowX="hidden"
    >
      <Heading mb="4" textAlign="center">Past Bets</Heading>
      <Divider mb="4" />

      <List spacing={3}>
        {pastBets.map((bet) => (
          <ListItem
            key={bet.id}
            padding="4"
            borderWidth="1px"
            borderRadius="md"
            bg={bet.winner ? 'green.50' : 'red.50'}
          >
            <Text><b>Bet ID:</b> {bet.placed_bet}</Text>
            <Text><b>Result time:</b> {bet.result_time}</Text>
            <Text><b>Bet Amount:</b> ${bet.bet_amount}</Text>
          </ListItem>
        ))}
      </List>

      <Button
        colorScheme="teal"
        onClick={() => navigate('/betgame')}
        mt="4"
        width="full"
      >
        Back to Bet Game
      </Button>
    </Box>
  );
}

export default PastBetsPage;

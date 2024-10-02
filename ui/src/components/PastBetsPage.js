import React from 'react';
import { Box, Button, Heading, List, ListItem, Text, Divider } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

function PastBetsPage() {
  const navigate = useNavigate();

  return (
    <Box p="6" maxW="lg" mx="auto" mt="10" boxShadow="lg" borderRadius="md">
      <Heading mb="4" textAlign="center">Past Bets</Heading>
      <Divider mb="4" />

      <List spacing={3}>
        <ListItem padding="4" borderWidth="1px" borderRadius="md" bg="gray.50">
          <Text>Bet on Game 1: $10</Text>
        </ListItem>
        <ListItem padding="4" borderWidth="1px" borderRadius="md" bg="gray.50">
          <Text>Bet on Game 2: $20</Text>
        </ListItem>
        <ListItem padding="4" borderWidth="1px" borderRadius="md" bg="gray.50">
          <Text>Bet on Game 3: $5</Text>
        </ListItem>
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
import React from 'react';
import { Box, Button, Heading, Text, Divider, useToast } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

function GamePage() {
  const navigate = useNavigate();
  const toast = useToast(); // Initialize toast

  const placeBet = () => {
    // Logic for placing a bet
    toast({
      title: "Bet Placed.",
      description: "Your bet of $10 has been placed successfully.",
      status: "success",
      duration: 5000,
      isClosable: true,
      position: "top", // You can set this to "top", "bottom", etc.
    });
  };

  return (
    <Box p="6" maxW="lg" mx="auto" mt="10" boxShadow="lg" borderRadius="md">
      <Heading mb="4" textAlign="center">Game Page</Heading>
      <Divider mb="4" />
      <Text mb="4" textAlign="center">Place your bets!</Text>

      <Button colorScheme="blue" onClick={placeBet} width="full" mb="4">
        Bet $10
      </Button>

      <Button
        colorScheme="teal"
        onClick={() => navigate('/betgame')}
        width="full"
      >
        Back to Bet Game
      </Button>
    </Box>
  );
}

export default GamePage;

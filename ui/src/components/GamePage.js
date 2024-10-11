import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Box,
    Heading,
    Text,
    Button,
    Flex,
    FormLabel,
    FormControl,
    Grid,
    GridItem,
    Container,
    NumberInput,
    NumberInputField,
    NumberInputStepper,
    NumberIncrementStepper,
    NumberDecrementStepper,
    useToast,
} from '@chakra-ui/react';
import axios from 'axios';

export default function GamePage() {
    const [games, setGames] = React.useState({ active: [] });
    const [showForm, setShowForm] = useState(false);
    const [currentGame, setCurrentGame] = useState(null);
    const [luckyNumber, setLuckyNumber] = useState(1);
    const [betAmount, setBetAmount] = useState(100);
    const toast = useToast();
    const navigate = useNavigate();

    React.useEffect(() => {
        axios.post('http://localhost:8080/getgames')
            .then((response) => {
                const currentTime = new Date();
                const activeGames = [];
                response.data.data.forEach((game) => {
                    const startTime = new Date(game.start);
                    const endTime = new Date(game.end);
                    if (currentTime >= startTime && currentTime <= endTime) {
                        activeGames.push(game);
                    }
                });
                setGames({ active: activeGames });
            })
            .catch((error) => console.error(error));
    }, []);

    const addBet = (game) => {
        setCurrentGame(game);
        setShowForm(true);
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        const uid = localStorage.getItem('uid');
        const data = { gameid: currentGame.id, number: parseInt(luckyNumber), amount: parseInt(betAmount) };
        axios.post(`http://localhost:8080/addbet/${uid}`, data)
            .then(() => {
                toast({
                    title: 'Bet Placed',
                    description: 'Your bet has been placed successfully',
                    status: 'success',
                    duration: 1000,
                    isClosable: true,
                    position: 'top',
                });
                setShowForm(false);
                setCurrentGame(null);
            })
            .catch((error) => {
                if (error.response.status === 403 && error.response.data.message.includes('insufficient fund')) {
                    toast({
                        title: 'Insufficient Balance',
                        description: error.response.data.message + "\n" + error.response.data.data,
                        status: 'error',
                        duration: 1000,
                        isClosable: true,
                        position: 'top',
                    });
                } else {
                    toast({
                        title: 'Unknown issue',
                        description: String(error.response.data.data),
                        status: 'error',
                        duration: 1000,
                        isClosable: true,
                        position: 'top',
                    });
                }
            });
    };

    const handleCancel = () => {
        setShowForm(false);
        setCurrentGame(null);
    };

    return (
        <Container maxW="container.md" py={4}>
            <Flex justify="space-between" align="center" mb={4}>
                <Heading textAlign="center">Active Games</Heading>
                <Button onClick={() => navigate('/betgame')} colorScheme="teal">Back</Button>
            </Flex>
            {showForm && currentGame ? (
                <Box mt={4}>
                    <FormControl as="form" onSubmit={handleSubmit}>
                        <FormLabel>Lucky Number:</FormLabel>
                        <NumberInput
                            value={luckyNumber}
                            onChange={(value) => setLuckyNumber(value)}
                            min={1}
                            max={100}
                            step={1}
                        >
                            <NumberInputField />
                            <NumberInputStepper>
                                <NumberIncrementStepper />
                                <NumberDecrementStepper />
                            </NumberInputStepper>
                        </NumberInput>
                        <FormLabel>Bet Amount:</FormLabel>
                        <NumberInput
                            value={betAmount}
                            onChange={(value) => setBetAmount(value)}
                            min={1}
                            step={1}
                        >
                            <NumberInputField />
                            <NumberInputStepper>
                                <NumberIncrementStepper />
                                <NumberDecrementStepper />
                            </NumberInputStepper>
                        </NumberInput>
                        <Flex justify="space-between" mt={4} mb={4}>
                            <Button type="submit" w="full" bg="green.500" color="white" mr={2}>
                                Place Bet
                            </Button>
                            <Button onClick={handleCancel} w="full" variant="outline" colorScheme="red" ml={2}>
                                Cancel
                            </Button>
                        </Flex>
                    </FormControl>
                </Box>
            ) : (
                <Grid templateColumns="repeat(1, 1fr)" gap={4} mt={4}>
                    {games.active.map((game) => (
                        <GridItem key={game.id}>
                            <Box
                                p={4}
                                bg="white"
                                boxShadow="md"
                                borderRadius="md"
                                display="flex"
                                flexDirection="column"
                            >
                                <Flex align="center" justify="space-between">
                                    <Text flex="1" fontWeight="bold">{game.name}</Text>
                                    <Button
                                        onClick={() => addBet(game)}
                                        colorScheme="teal"
                                        variant="outline"
                                    >
                                        Add Bet
                                    </Button>
                                </Flex>
                            </Box>
                        </GridItem>
                    ))}
                </Grid>
            )}
        </Container>
    );
}





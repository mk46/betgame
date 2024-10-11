import React, { useState } from 'react';
import { Box, Button, Flex, FormControl, FormLabel, Input, Heading, useToast } from '@chakra-ui/react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const AddGame = () => {
    const toast = useToast();
    const navigate = useNavigate();
    const [name, setName] = useState('');
    const [start, setStart] = useState('');
    const [end, setEnd] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();
        if (!name || !start || !end) {
            toast({
                title: "Error",
                description: "All fields are required.",
                status: "error",
                duration: 1000,
                isClosable: true,
                position: "top"
            });
            return;
        }
        const startUtc = new Date(start).toISOString();
        const endUtc = new Date(end).toISOString();
        const game = { name: name, start: startUtc, end: endUtc };

        axios.post('http://localhost:8080/addgame', game)
            .then((response) => {
                console.log(response.data);
                setName('');
                setStart('');
                setEnd('');
                toast({
                    title: "Game added.",
                    description: `Game ${name} added successfully.`,
                    status: "success",
                    duration: 1000,
                    isClosable: true,
                    position: "top"
                });
            })
            .catch((error) => {
                toast({
                    title: "Error",
                    description: error.message,
                    status: "error",
                    duration: 1000,
                    isClosable: true,
                    position: "top"
                });
            });
    };
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
            <Heading as="h2" size="lg">Add Game</Heading>
            <FormControl>
                <FormLabel>Name:</FormLabel>
                <Input type="text" value={name} onChange={(event) => setName(event.target.value)} />
            </FormControl>
            <FormControl>
                <FormLabel>Start:</FormLabel>
                <Input type="datetime-local" value={start} onChange={(event) => setStart(event.target.value)} />
            </FormControl>
            <FormControl>
                <FormLabel>End:</FormLabel>
                <Input type="datetime-local" value={end} onChange={(event) => setEnd(event.target.value)} />
            </FormControl>
            <Flex justifyContent="space-between" alignItems="center" mt={4}>
                <Button type="submit" colorScheme="teal" onClick={handleSubmit}>Add Game</Button>
                <Button colorScheme="blue" variant="outline" onClick={() => navigate('/admin')}>Go back to admin page</Button>
            </Flex>
        </Box>
    );
};

export default AddGame;


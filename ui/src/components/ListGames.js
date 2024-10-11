import axios from 'axios';
import { useEffect, useState } from 'react';
import { Container, Box, Button, Input, Heading, FormControl, FormLabel, IconButton, Flex, Text, useToast } from '@chakra-ui/react';
import { FaPen, FaTrash, FaBullseye } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';

const ListGames = () => {
    const [games, setGames] = useState([]);
    const [editing, setEditing] = useState(null);
    const [declaringResult, setDeclaringResult] = useState(null);
    const toast = useToast();
    const navigate = useNavigate();

    useEffect(() => {
        const getGames = async () => {
            const response = await axios.post('http://localhost:8080/getgames');
            setGames(response.data.data);
        }

        getGames();
    }, []);

    const deleteGame = async (id) => {
        alert("Delete Game: " + id);
        return
        const response = await axios.delete(`http://localhost:8080/deletegame/${id}`);
        if (response.status === 200) {
            const newGames = games.filter(game => game.id !== id);
            setGames(newGames);
        }
    }

    const startEditing = (game) => {
        setEditing({ ...game, start: new Date(game.start), end: new Date(game.end) });
    }

    const saveEdit = async (game) => {
        const response = await axios.post(`http://localhost:8080/reschedulegame/`, {
            id: game.id,
            name: game.name,
            start: game.start.toISOString(),
            end: game.end.toISOString()
        });
        if (response.status === 202) {
            const newGames = games.map(g => g.id === game.id ? game : g);
            setGames(newGames);
            setEditing(null);
            toast({
                title: "Game updated",
                description: "Game updated successfully",
                status: "success",
                duration: 1000,
                isClosable: true,
                position: "top"
            });
            navigate('/admin/listgames');
        }
    }

    const cancelEdit = () => {
        setEditing(null);
    }

    const declareResult = async (game) => {
        setDeclaringResult(game);
    }

    const saveResult = async (result) => {
        const data = {
            gameid: declaringResult?.id,
            result: result
        };
        console.log(data)
        axios.post('http://localhost:8080/declairewinner/', data)
            .then(response => {
                if (response.status === 202) {
                    toast({
                        title: "Result declared",
                        description: `Result declared for game ${declaringResult?.name}`,
                        status: "success",
                        duration: 1000,
                        isClosable: true,
                        position: "top"
                    });
                }
            })
            .catch(error => {
                if (error.response.status === 403 && error.response.data.message.includes('failed to remove Bets')) {
                    toast({
                        title: "Result declared",
                        description: `There is no any bet placed for game: ${declaringResult?.name}`,
                        status: "error",
                        duration: 1000,
                        isClosable: true,
                        position: "top"
                    });
                }

            });
        setDeclaringResult(null);
        navigate('/admin/listgames');

    }

    const cancelResult = () => {
        setDeclaringResult(null);
    }

    return (
        <Container mt={10}>
            <Flex justify="space-between" align="center">
                <Heading>Games</Heading>
                <Button onClick={() => navigate('/admin')}>Back to Admin</Button>
            </Flex>
            {editing ? (
                <Box mt={2} p={4} boxShadow="md" borderRadius="md">
                    <Flex justify="space-between" align="center">
                        <Box>
                            <Heading size="sm">
                                <FormControl>
                                    <FormLabel>Name</FormLabel>
                                    <Input value={editing.name} onChange={(e) => setEditing({ ...editing, name: e.target.value })} />
                                </FormControl>
                            </Heading>
                            <Text>
                                <FormControl>
                                    <FormLabel>Start</FormLabel>
                                    <Input type="datetime-local" value={editing.start.toISOString().slice(0, 16)} onChange={(e) => setEditing({ ...editing, start: new Date(e.target.value) })} />
                                </FormControl>
                            </Text>
                            <Text>
                                <FormControl>
                                    <FormLabel>End</FormLabel>
                                    <Input type="datetime-local" value={editing.end.toISOString().slice(0, 16)} onChange={(e) => setEditing({ ...editing, end: new Date(e.target.value) })} />
                                </FormControl>
                            </Text>
                        </Box>
                    </Flex>
                    <Flex mt={2} justify="flex-start">
                        <Button mr={2} colorScheme="green" onClick={() => saveEdit(editing)}>Save</Button>
                        <Button mr={2} colorScheme="red" onClick={cancelEdit}>Cancel</Button>
                    </Flex>
                </Box>
            ) : declaringResult ? (
                <Box mt={2} p={4} boxShadow="md" borderRadius="md">
                    <Flex justify="space-between" align="center">
                        <Box>
                            <Heading size="sm">Declare Result for {declaringResult.name}</Heading>
                            <FormControl>
                                <FormLabel>Result</FormLabel>
                                <Input type="number" value={declaringResult.result} onChange={(e) => setDeclaringResult({ ...declaringResult, result: parseFloat(e.target.value) })} />
                            </FormControl>
                        </Box>
                    </Flex>
                    <Flex mt={2} justify="flex-start">
                        <Button mr={2} colorScheme="green" onClick={() => saveResult(declaringResult.result)}>Save</Button>
                        <Button mr={2} colorScheme="red" onClick={cancelResult}>Cancel</Button>
                    </Flex>
                </Box>
            ) : (

                games.map(game => (
                    <Box mt={2} key={game.id} p={4} boxShadow="md" borderRadius="md">
                        <Flex justify="space-between" align="center">
                            <Box>
                                <Heading size="sm">{game.name}</Heading>
                                <Text>Start: {game.start.toLocaleString()}</Text>
                                <Text>End: {game.end.toLocaleString()}</Text>
                            </Box>
                            <Flex>
                                <IconButton
                                    aria-label="Edit"
                                    icon={<FaPen color="green" />}
                                    size="sm"
                                    onClick={() => startEditing(game)}
                                />
                                <IconButton
                                    aria-label="Declare Result"
                                    icon={<FaBullseye color="green" />}
                                    size="sm"
                                    onClick={() => declareResult(game)}
                                />
                                <IconButton
                                    aria-label="Delete"
                                    icon={<FaTrash color="red" />}
                                    size="sm"
                                    onClick={() => deleteGame(game.id)}
                                />
                            </Flex>
                        </Flex>
                    </Box>
                ))
            )}
        </Container>
    );
}

export default ListGames;




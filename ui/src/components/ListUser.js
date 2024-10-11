import axios from 'axios';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Box,
    Heading,
    Text,
    Grid,
    GridItem,
    Button,
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton,
    Input,
    FormLabel,
    useToast,
    Flex,
    Center,
} from '@chakra-ui/react';

function ListUser() {
    const toast = useToast();
    const [users, setUsers] = useState([]);
    const [isOpen, setIsOpen] = useState(false);
    const [selectedUserId, setSelectedUserId] = useState(null);
    const [amount, setAmount] = useState('');
    const [action, setAction] = useState(''); // add or withdraw

    const navigate = useNavigate();
    useEffect(() => {
        setAmount('');
    }, [isOpen]);

    const getUsers = () => {
        axios.get('http://localhost:8080/getusers')
            .then((response) => {
                setUsers(response.data.data);
            })
            .catch((error) => {
                console.error(error);
                toast({
                    title: 'Error',
                    description: error.message,
                    status: 'error',
                    duration: 1000,
                    isClosable: true,
                    position: 'top',
                });
            });
    };

    useEffect(() => {
        getUsers();
    }, []);

    const handleAddCash = (userId) => {
        setSelectedUserId(userId);
        setIsOpen(true);
        setAction('addcash');
    };

    const handleWithdrawCash = (userId) => {
        setSelectedUserId(userId);
        setIsOpen(true);
        setAction('withdrawcash');
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        if (selectedUserId) {
            axios.post(`http://localhost:8080/${action}/${selectedUserId}`, {
                amount: Math.abs(amount),
            })
                .then((response) => {
                    toast({
                        title: 'Success',
                        description: response.data.message,
                        status: 'success',
                        duration: 1000,
                        isClosable: true,
                        position: 'top',
                    });
                    setIsOpen(false);
                    getUsers();
                })
                .catch((error) => {
                    console.error(error);
                    toast({
                        title: 'Error',
                        description: "Failed to add/withdraw cash. " + error.message,
                        status: 'error',
                        duration: 1000,
                        isClosable: true,
                        position: 'top',
                    });
                });
        }
    };

    return (
        <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
        >
            <Center>
                <Box>
                    <Heading size="lg" m="4" alignItems="center">
                        List of User
                    </Heading>

                    <Flex justifyContent="flex-end" mt={4}>
                        <Button
                            colorScheme="blue"
                            onClick={() => navigate('/admin')}
                        >
                            back to Admin
                        </Button>
                    </Flex>

                    <Grid templateColumns="repeat(2, 1fr)" gap={2}>
                        {users.map((user) => (
                            <GridItem key={user.id} p="2" borderRadius="md" bg="gray.100" display="flex" alignItems="center">
                                <Text fontSize="md" m={4}>
                                    {user.name}
                                </Text>
                                <Text fontSize="md" m={4}>
                                    {user.phone}
                                </Text>
                                <Text fontSize="md" m={4}>
                                    ${user.balance}
                                </Text>
                                <Button
                                    size="sm"
                                    colorScheme="green"
                                    onClick={() => handleAddCash(user.id)}
                                >
                                    Add Cash
                                </Button>
                                <Button
                                    size="sm"
                                    colorScheme="red"
                                    onClick={() => handleWithdrawCash(user.id)}
                                    ml={2}
                                >
                                    Withdraw Cash
                                </Button>
                            </GridItem>
                        ))}
                    </Grid>
                    <Modal
                        isOpen={isOpen}
                        onClose={() => setIsOpen(false)}
                    >
                        <ModalOverlay />
                        <ModalContent>
                            <ModalHeader>{action === 'addcash' ? 'Add Cash' : 'Withdraw Cash'}</ModalHeader>
                            <ModalCloseButton />
                            <ModalBody>
                                <form onSubmit={handleSubmit}>
                                    <FormLabel>
                                        Amount
                                    </FormLabel>
                                    <Input
                                        type="number"
                                        value={amount}
                                        onChange={(event) => setAmount(event.currentTarget.value)}
                                        placeholder="Enter amount"
                                    />
                                    <Button
                                        type="submit"
                                        colorScheme="blue"
                                        mt={4}
                                    >
                                        Submit
                                    </Button>
                                </form>
                            </ModalBody>
                        </ModalContent>
                    </Modal>

                </Box>
            </Center>
        </Box>
    );
}

export default ListUser;


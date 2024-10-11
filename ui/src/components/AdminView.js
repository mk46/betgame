import React from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Button,
    Heading,
    Flex,
} from '@chakra-ui/react';

const AdminView = () => {
    const navigate = useNavigate();
    return (
        <Flex direction="column" alignItems="center" justifyContent="center">
            <Heading mt={8}>Admin View</Heading>
            <Flex mt={4} justify="space-between">
                <Button mr={4} colorScheme="teal" onClick={() => { navigate('/admin/addgame') }}>Add Game</Button>
                <Button mr={4} colorScheme="teal" onClick={() => { navigate('/admin/listgames') }}>List Games</Button>
                <Button mr={4} colorScheme="teal" onClick={() => { navigate('/admin/getusers') }}>List Users</Button>
            </Flex>
        </Flex>
    );
};

export default AdminView;


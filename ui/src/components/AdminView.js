import React from 'react';
import {
    Box,
    Button,
    Heading,
    Flex,
} from '@chakra-ui/react';

const AdminView = () => {
    return (
        <Flex direction="column" alignItems="center" justifyContent="center">
            <Heading mt={8}>Admin View</Heading>
            <Flex mt={4} justify="space-between">
                <Button mr={4} colorScheme="teal">Add Game</Button>
                <Button mr={4} colorScheme="teal">List Games</Button>
                <Button mr={4} colorScheme="teal">Reschedule Game</Button>
                <Button mr={4} colorScheme="teal">Withdraw Cash</Button>
                <Button mr={4} colorScheme="teal">Add Cash</Button>
                <Button mr={4} colorScheme="teal">Announce Result</Button>
            </Flex>
        </Flex>
    );
};

export default AdminView;


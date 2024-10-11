import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Input,
  Heading,
  Flex,
} from '@chakra-ui/react';

const AdminLogin = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = () => {
    // authenticate admin credentials
    if (username === 'admin' && password === 'password') {
      // redirect to admin view page
      navigate('/admin');
    } else {
      alert('Invalid credentials');
    }
  };

  return (
    <Box
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
      w="30%"
      m="auto"
    >
      <Heading alignSelf="center">Admin Login</Heading>
      <FormControl mt={4}>
        <FormLabel>Username:</FormLabel>
        <Input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
      </FormControl>
      <FormControl mt={4}>
        <FormLabel>Password:</FormLabel>
        <Input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </FormControl>
      <Flex mt={4} justifyContent="space-between" alignItems="center">
        <Button
          onClick={handleLogin}
          colorScheme="teal"
        >
          Login
        </Button>
        <Box mx={2} />
        <Button
          onClick={() => {
            setUsername('');
            setPassword('');
          }}
          variant="solid"
          colorScheme="red"
        >
          Clear
        </Button>
      </Flex>
    </Box>
  );
};


export default AdminLogin;


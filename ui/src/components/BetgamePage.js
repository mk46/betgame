import React from 'react';
import { useNavigate } from 'react-router-dom';
import { Flex, Box, Heading, Button, Spacer, Input, Text, IconButton, Icon, Menu, MenuItem, MenuList, MenuButton, Container, List, ListItem } from '@chakra-ui/react';
import { FaUser } from 'react-icons/fa';

function BetGamePage() {
  const navigate = useNavigate();
  const [user, setUser] = React.useState({ name: '', phone: '', email: '', balance: 0 });
  const [isEditing, setIsEditing] = React.useState(false);
  const [bets, setBets] = React.useState([]);

  React.useEffect(() => {
    const sessionId = localStorage.getItem('session-id');

    if (!sessionId) {
      // Redirect to login page if session ID is not present
      navigate('/');
      // don't continue with the rest of the code
      return;
    }

    const options = {
      method: 'GET',
      headers: {
        'Authorization': sessionId,
      },
    };

    fetch('http://localhost:8080/getuser', options)
      .then(response => response.json())
      .then(data => {
        setUser(data.data);
        localStorage.setItem('uid', data.data.id);
      })
      .catch(error => console.error(error));

    fetch(`http://localhost:8080/getbets/${localStorage.getItem('uid')}`)
      .then(response => response.json())
      .then(data => {
        setBets(data.data);
      })
      .catch(error => console.error(error));
  }, []);

  const handleUpdate = (event) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const name = formData.get('name');
    const email = formData.get('email');

    if (name === user.name && email === user.email) {
      // if name and email are same, don't update and close the form
      setIsEditing(false);
      return;
    }

    const options = {
      method: 'POST',
      headers: {
        'Authorization': localStorage.getItem('session-id'),
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: name,
        email: email,
        phone: user.phone,
        balance: user.balance,
      }),
    };

    fetch('http://localhost:8080/updateuser', options)
      .then(response => response.json())
      .then(data => {
        setUser(data.data);
        setIsEditing(false);
      })
      .catch(error => console.error(error));
  };

  const handleLogout = () => {
    localStorage.removeItem('session-id');
    localStorage.removeItem('uid');
    navigate('/');
  };

  return (
    <Container direction="column" align="center" overflowX="hidden" maxW="container.md">
      <Box p="8" display="flex" justifyContent="space-between" alignItems="center">
        <Heading mb="6" textAlign="center">Bet Game</Heading>
        <Flex justify="space-between" align="center">
          <Button colorScheme="teal" mr="4" onClick={() => navigate('/game')}> Add Bet</Button>
          <Button colorScheme="teal" mr="4" onClick={() => navigate('/past-bets')}>Past Bets</Button>
          <Text mr="4"><b>Cash:</b> ${user.balance}</Text>
          <Box align="right" mr="4">
            <Menu>
              <MenuButton as={IconButton} icon={<Icon as={FaUser} />} />
              <MenuList align="right">
                <MenuItem onClick={() => setIsEditing(!isEditing)}>Edit Profile</MenuItem>
                <MenuItem onClick={handleLogout}>Logout</MenuItem>
              </MenuList>
            </Menu>
          </Box>
        </Flex>
      </Box>
      <Spacer />
      {isEditing && (
        <Box p="8" mt="6">
          <form onSubmit={handleUpdate}>
            <Box>
              <Input type="text" name="name" defaultValue={user.name} placeholder="Name" mr="4" />
              <Input type="email" name="email" defaultValue={user.email} placeholder="Email" mr="4" />
              <Input type="number" name="balance" value={user.balance || 0} isDisabled />
              <Button type="submit" colorScheme="teal" mr="4">Update</Button>
              <Button onClick={() => setIsEditing(false)} colorScheme="blue" variant="outline">Cancel</Button>
            </Box>
          </form>
        </Box>
      )}
      {!isEditing && bets.length > 0 && (
        <Box p="8" mt="6">
          <Heading mb="4" textAlign="center">Your Bets</Heading>
          <List spacing={3} align="left">
            {bets.map(bet => (
              <ListItem key={bet.id} padding="4" borderWidth="1px" borderRadius="md" bg="gray.50">
                <Text><b>Game ID:</b> {bet.gameid}</Text>
                <Text><b>Amount:</b> {bet.amount}</Text>
                <Text><b>Number:</b> {bet.number}</Text>
                <Text><b>Placed Time:</b> {new Date(bet.placed_at).toLocaleString()}</Text>
              </ListItem>
            ))}
          </List>
        </Box>
      )}
    </Container>
  );
}

export default BetGamePage;

import React from 'react'
import {
    Container,
    CssBaseline,
    Box,
    Avatar,
    Typography,
    TextField,
    Button,
    Grid,
  } from "@mui/material";
import axios from 'axios';

const Home = () => {
// Function to start the server
const startServer = async () => {
  console.log('Cookies:', document.cookie);
  try {
    const response = await axios.get('http://localhost:5432/vm/server/start', {
      withCredentials: true, // Include cookies
    });
    console.log('Server started:', response.data);
    alert('Server started successfully!');
  } catch (error) {
    console.error('Error starting server:', error);
    alert('Failed to start the server.');     
  }
};

// Function to stop the server
const stopServer = async () => {
  try {
    const response = await axios.get('http://localhost:5432/vm/server/stop', {
      withCredentials: true, // Include cookies
    });
    console.log('Server stopped:', response.data);
    alert('Server stopped successfully!');
  } catch (error) {
    console.error('Error stopping server:', error);
    alert('Failed to stop the server.');
  }
};

return (
  <Box
    sx={{
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      justifyContent: 'center',
      height: '100vh',
      backgroundColor: '#fafafa', // Set background color to a consistent light gray
      textAlign: 'center', // Center text horizontally
    }}
  >
    <Typography variant="h3" sx={{ marginBottom: 4 }}>
      Minecraft Server Dashboard
    </Typography>
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        gap: 3, // Adds space between buttons
      }}
    >
      <Button
        variant="contained"
        color="primary"
        size="large"
        onClick={startServer}
        sx={{ padding: '15px 30px', fontSize: '16px' }}
      >
        Start Server
      </Button>
      <Button
        variant="contained"
        color="secondary"
        size="large"
        onClick={stopServer}
        sx={{ padding: '15px 30px', fontSize: '16px' }}
      >
        Stop Server
      </Button>
    </Box>
  </Box>
);
}

export default Home
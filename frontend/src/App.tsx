import React from 'react';
import axios from 'axios';
import { Button, Container, Typography, Box } from '@mui/material';

const App: React.FC = () => {
  // Define your backend API endpoints for starting and stopping the server
  const startServerUrl = 'http://your-backend-vm-ip/start-server';
  const stopServerUrl = 'http://your-backend-vm-ip/stop-server';

  // Function to start the server
  const startServer = async () => {
    try {
      const response = await axios.post(startServerUrl);
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
      const response = await axios.post(stopServerUrl);
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
};

export default App;

import React, { useState } from 'react'
import {
    Container,
    CssBaseline,
    Box,
    Avatar,
    Typography,
    TextField,
    Button,
    Grid,
    LinearProgress,
    Snackbar,
  } from "@mui/material";
import axios from 'axios';


const Home = () => {
  const [loading, setLoading] = useState(false);
  const [openSnackbar, setOpenSnackbar] = useState(false); 
  const [snackbarMessage, setSnackbarMessage] = useState('');

// Function to start the server
const startServer = async () => {
  console.log('Cookies:', document.cookie);
  try {
    const response = await axios.get('http://localhost:5432/vm/server/start', {
      withCredentials: true, // Include cookies
    });
    console.log('Server started:', response.data);
    setSnackbarMessage('Server started successfully!');
    setOpenSnackbar(true);
  } catch (error) {
    console.error('Error starting server:', error);
    setSnackbarMessage('Failed to start the server.');
    setOpenSnackbar(true);    
  }
};

// Function to stop the server
const stopServer = async () => {
  setLoading(true);
  try {
    const response = await axios.get('http://localhost:5432/vm/server/stop', {
      withCredentials: true, // Include cookies
    });
    console.log('Server stopped:', response.data);
    setSnackbarMessage('Server stopped successfully!');
  } catch (error) {
    console.error('Error stopping server:', error);
    setSnackbarMessage('Failed to stop the server.');
  } finally {
    setLoading(false); // Stop loading bar
    setOpenSnackbar(true);
  }
};

const handleCloseSnackbar = () => {
  setOpenSnackbar(false);
};

return (
  <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        backgroundColor: '#fafafa',
        textAlign: 'center',
      }}
    >
      <Typography variant="h3" sx={{ marginBottom: 4 }}>
        Minecraft Server Dashboard
      </Typography>
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          gap: 3,
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

      {/* Show loading bar when loading is true */}
      {loading && <LinearProgress sx={{ width: '20%', marginTop: 4 }} />}
      <Snackbar
        open={openSnackbar}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
        message={snackbarMessage}
      />
    </Box>
);
}

export default Home
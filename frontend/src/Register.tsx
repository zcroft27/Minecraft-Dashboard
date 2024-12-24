import {
    Avatar,
    Box,
    Button,
    Container,
    CssBaseline,
    Grid,
    TextField,
    Typography,
} from "@mui/material";
import { LockOutlined } from "@mui/icons-material";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from 'axios';

const Register = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleRegister = async () => {
        try {
            const response = await axios.post('http://localhost:5432/sign-up', {
                email: email,
                password: password,
                credentials : true
            });
            
            navigate('/login');
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <>
            <Container maxWidth="xs">
                <CssBaseline />
                <Box
                    sx={{
                        mt: 20,
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                    }}
                >
                    <Avatar sx={{ m: 1, bgcolor: "primary.light" }}>
                        <LockOutlined />
                    </Avatar>
                    <Typography variant="h5">Register</Typography>
                    <Box sx={{ mt: 3 }}>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>
                                <TextField
                                    required
                                    fullWidth
                                    id="email"
                                    label="Email Address"
                                    name="email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField
                                    required
                                    fullWidth
                                    name="password"
                                    label="Password"
                                    type="password"
                                    id="password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                />
                            </Grid>
                        </Grid>
                        <Button
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                            onClick={handleRegister}
                        >
                            Register
                        </Button>
                        <Grid container justifyContent="flex-end">
                            <Grid item>
                                <Link to="/login">Already have an account? Login</Link>
                            </Grid>
                        </Grid>
                    </Box>
                </Box>
            </Container>
        </>
    );
};

export default Register;

import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Box,
    Button,
    FormControl,
    FormLabel,
    Input,
    Heading,
} from '@chakra-ui/react';

function LoginPage() {
    const [step, setStep] = useState(1); // 1 for mobile input, 2 for OTP input
    const [mobile, setMobile] = useState('');
    const [otp, setOtp] = useState('');
    const navigate = useNavigate();

    const handleMobileSubmit = (e) => {
        e.preventDefault();
        console.log('Sending OTP to:', mobile);
        setStep(2); // Move to OTP step
    };

    const handleOtpSubmit = (e) => {
        e.preventDefault();
        console.log('Validating OTP:', otp);
        // Redirect to the bet game page on successful validation
        navigate('/betgame');
    };

    return (
        <Box width="400px" mx="auto" mt="100px" p="4" boxShadow="md" borderRadius="md">
            <Heading mb="6">{step === 1 ? 'Login' : 'Verify OTP'}</Heading>
            <form onSubmit={step === 1 ? handleMobileSubmit : handleOtpSubmit}>
                {step === 1 ? (
                    <FormControl id="mobile" mb="4" isRequired>
                        <FormLabel>Mobile Number</FormLabel>
                        <Input
                            type="tel"
                            placeholder="Enter your mobile number"
                            value={mobile}
                            onChange={(e) => setMobile(e.target.value)}
                        />
                    </FormControl>
                ) : (
                    <FormControl id="otp" mb="4" isRequired>
                        <FormLabel>Enter OTP</FormLabel>
                        <Input
                            type="text"
                            placeholder="Enter the OTP sent to your mobile"
                            value={otp}
                            onChange={(e) => setOtp(e.target.value)}
                        />
                    </FormControl>
                )}
                <Button colorScheme="teal" type="submit" width="full">
                    {step === 1 ? 'Send OTP' : 'Verify OTP'}
                </Button>
            </form>
        </Box>
    );
}

export default LoginPage;

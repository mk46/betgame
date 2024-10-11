import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Box,
    Button,
    FormControl,
    FormLabel,
    Input,
    Heading,
    InputGroup,
} from '@chakra-ui/react';

function LoginPage() {
    const [step, setStep] = useState(1); // 1 for mobile input, 2 for OTP input
    const [mobile, setMobile] = useState('');
    const [otp, setOtp] = useState('');
    const [countryCode, setCountryCode] = useState('+91');
    const navigate = useNavigate();

    React.useEffect(() => {
        const sessionId = localStorage.getItem('session-id');

        if (sessionId) {
            // Redirect to the bet game page if session ID is present
            navigate('/betgame');
            // don't continue with the rest of the code
            return;
        }
    }, []);

    const handleMobileSubmit = (e) => {
        e.preventDefault();
        const requestBody = { phone: countryCode + mobile };
        fetch('http://localhost:8080/loginuser', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestBody),
        }).then((res) => res.json())
            .then((data) => {
                console.log(data);
                setStep(2); // Move to OTP step
            })
            .catch((err) => console.log(err));
        console.log('Sending OTP to:', countryCode + mobile);
        setStep(2); // Move to OTP step    };
    }
    const handleOtpSubmit = (e) => {
        e.preventDefault();
        const requestBody = { phone: countryCode + mobile, code: otp };
        fetch('http://localhost:8080/verifyOTP', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestBody),
        }).then((res) => res.json())
            .then((data) => {
                console.log(data);
                localStorage.setItem('session-id', data.data);
                // Redirect to the bet game page on successful validation
                navigate('/betgame');
            })
            .catch((err) => console.log(err));
        console.log('Validating OTP:', otp);
    };

    return (
        <Box width="400px" mx="auto" mt="100px" p="4" boxShadow="md" borderRadius="md">
            <Heading mb="6">{step === 1 ? 'Login' : 'Verify OTP'}</Heading>
            <form onSubmit={step === 1 ? handleMobileSubmit : handleOtpSubmit}>
                {step === 1 ? (
                    <FormControl id="mobile" mb="4" isRequired>
                        <FormLabel>Mobile Number</FormLabel>
                        <InputGroup>
                            <Input
                                type="tel"
                                placeholder="Enter your country code"
                                value={countryCode}
                                onChange={(e) => setCountryCode(e.target.value)}
                                width="120px"
                            />
                            <Input
                                type="tel"
                                placeholder="Enter your mobile number"
                                value={mobile}
                                onChange={(e) => setMobile(e.target.value)}
                            />
                        </InputGroup>
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

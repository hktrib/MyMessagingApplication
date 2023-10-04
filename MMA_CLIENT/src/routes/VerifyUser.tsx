import { useState } from "react";
import { useLocation } from "react-router-dom";
import { useSearchParams } from "react-router-dom";
import axios, { AxiosError, AxiosResponse } from 'axios';

import * as React from 'react';
import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import { green } from '@mui/material/colors';
import Button from '@mui/material/Button';
import Fab from '@mui/material/Fab';
import CheckIcon from '@mui/icons-material/Check';
import CloseIcon from '@mui/icons-material/Close';
import '../styles/VerifyEmail.scss'


interface VerifyUserReq {
    username: string,
    secretcode: string
}

interface Error {
    errorMessage: string
}

const VerifyUser = () => {
    
    const [loading, setLoading] = React.useState(false);
    const [success, setSuccess] = React.useState(false);
    const [searchParams, setSearchParams] = useSearchParams();
    const timer = React.useRef<number>();

    const buttonSx = {
    ...(success && {
        bgcolor: green[500],
        '&:hover': {
        bgcolor: green[700],
        },
    }),
    };

    React.useEffect(() => {
        return () => {
            clearTimeout(timer.current);
        };
    }, []);

    const handleButtonClick = () => {
        if (!loading) {
            setSuccess(false);
            setLoading(true);
            axios.put<VerifyUserReq>("http://localhost:8080/verifyUser?",
                JSON.stringify({username: username_val, secret_code: secret_code_val}), config)
                .then(res => {
                    const response : any = res.
                    if (response instanceof Error){
                        // setLoad(true)
                    }
                })
                .catch(error => {
                    console.log(error)
                })
        }
    };
    
    
    
    const verifyQueryParams = () => {
        const config = {
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            }
        }
        const username_val = searchParams.get("username")
        const secret_code_val = searchParams.get("secret_code")


    }


    return (
        <div className="VerifyEmail">
            <h3>Click Below to Verify your Registration</h3>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Box sx={{ m: 1, position: 'relative' }}>
                <Fab
                  aria-label="save"
                  color="primary"
                  sx={buttonSx}
                  onClick={handleButtonClick}
                >
                  {success ? <CheckIcon /> : <CloseIcon/>}
                </Fab>
                {loading && (
                  <CircularProgress
                    size={68}
                    sx={{
                      color: green[500],
                      position: 'absolute',
                      top: -6,
                      left: -6,
                      zIndex: 1,
                    }}
                  />
                )}
              </Box>
              <Box sx={{ m: 1, position: 'relative' }}>
                <Button
                  variant="contained"
                  sx={buttonSx}
                  disabled={loading}
                  onClick={handleButtonClick}
                >
                  Accept terms
                </Button>
                {loading && (
                  <CircularProgress
                    size={24}
                    sx={{
                      color: green[500],
                      position: 'absolute',
                      top: '50%',
                      left: '50%',
                      marginTop: '-12px',
                      marginLeft: '-12px',
                    }}
                  />
                )}
              </Box>
            </Box>
        </div>
    )
}


export default VerifyUser;
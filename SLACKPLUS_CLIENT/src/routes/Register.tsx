import { useRef, useState, useEffect } from "react";
import CloseIcon from '@mui/icons-material/Close';
import CheckIcon from '@mui/icons-material/Check';
import InfoIcon from '@mui/icons-material/Info';
import * as React from 'react';
import '../styles/Register.scss'
import { useNavigate } from "react-router-dom";
import axios, { AxiosError, AxiosResponse } from 'axios';

const NAME_REGEX = /.+/;
const USER_REGEX = /^[A-z][A-z0-9-_]{3,23}$/;
const PWD_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/;
const EMAIL_REGEX = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/


interface UserRequest {
    username: string;
    emailAddress: string;
    password: string;
}

interface AvailableUsername {
    isAvailable : boolean
}
interface EmailAlreadyUsed {
    isUsed : boolean
}


const Register = () => {

    
    const userRef = useRef<HTMLInputElement>(null);
    const errRef = useRef<HTMLParagraphElement>(null);


    const [backendMessage, setBackendMessage] = useState('')
    const [username, setUsername] = useState('');
    const [validUsername, setValidUsername] = useState(false);
    const [usernameFocus, setUsernameFocus] = useState(false);


    // const [firstName, setFirstName] = useState('');
    // const [validFirstName, setValidFirstName] = useState(false);
    // const [firstNameFocus, setFirstNameFocus] = useState(false);

    // const [lastName, setLastName] = useState('');
    // const [validLastName, setValidLastName] = useState(false);
    // const [lastNameFocus, setLastNameFocus] = useState(false);

    const [email, setEmail] = useState('');
    const [validEmail, setValidEmail] = useState(false);
    const [emailFocus, setEmailFocus] = useState(false);

    const [pwd, setPwd] = useState('');
    const [validPwd, setValidPwd] = useState(false);
    const [pwdFocus, setPwdFocus] = useState(false);

    const [matchPwd, setMatchPwd] = useState('');
    const [validMatch, setValidMatch] = useState(false);
    const [matchFocus, setMatchFocus] = useState(false);

    const [errMsg, setErrMsg] = useState('');
    const [success, setSuccess] = useState(false);

    useEffect(() => {
        if (userRef.current)
            userRef.current.focus();
    }, [])
    useEffect(() => {
        const valid = USER_REGEX.test(username)
        // if (valid === true) {
        //     valid = availableUsername(username)
        // }
        setValidUsername(valid);
    }, [username])

    useEffect(() => {
        const valid = EMAIL_REGEX.test(email)
        // if (valid === true) {
        //     valid = emailUsed(email)
        // }
        setValidEmail(valid);
    }, [email])

    useEffect(() => {
        setValidPwd(PWD_REGEX.test(pwd));
        setValidMatch(pwd === matchPwd);
    }, [pwd, matchPwd])

    useEffect(() => {
        setErrMsg('');
    }, [username, pwd, matchPwd])


    const availableUsername = (username: string) : boolean => {
        console.log(username)
        axios.get(`http://localhost:8080/register?username=${username}`)
        .then((res : AxiosResponse<AvailableUsername>) => {
            const availUser : AvailableUsername = res.data;
            console.log("username -> entered .then")
            console.log(availUser)
            if (availUser.isAvailable == true) {
                setBackendMessage('')
                return true
            }
        })
        .catch((error: AxiosError<string>) => {
            // Handle error here
            console.error('An error occurred:', error.message);
        }); 
        
        setBackendMessage("Error: Username is not Available!!!")
        return false
    }
    const emailUsed = (email: string) : boolean => {
        
        axios.get(`http://localhost:8080/register?email=${email}`)
        .then((res : AxiosResponse<EmailAlreadyUsed>) => {
            const emailUsed : EmailAlreadyUsed = res.data
            console.log("email -> entered .then");
            console.log(emailUsed)
            if (emailUsed.isUsed == true) {
                setBackendMessage('')
                return true
            }
        })
        .catch((error: AxiosError<string>) => {
            // Handle error here
            console.error('An error occurred:', error.message);
        }); 
        
        setBackendMessage("Error: Email is already Registered....try Signing In instead!!!")
        return false
    }

    const handleSubmit = (e : React.FormEvent<HTMLFormElement>) : void => {
        e.preventDefault()
        console.log("Submit Button Clicked")
        const v1 = USER_REGEX.test(username);
        const v2 = PWD_REGEX.test(pwd);
        const v3 = EMAIL_REGEX.test(email)
        if (!v1 || !v2 || !v3) {
            setErrMsg("Invalid Entry");
        }
        else {
            console.log("Reached Else")
            const config = {
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                }
            }
            axios.post<UserRequest>('http://localhost:8080/register', 
                JSON.stringify({username : username, emailAddress: email, password: pwd}), config)
                .then(res => {
                    console.log("Post Request Sent....got back response");
                    console.log(JSON.stringify(res.data));

                    //clear state and controlled inputs
                    setSuccess(true);
                    setUsername('');
                    setPwd('');
                    setMatchPwd('');
                    return 1;
                })
                .catch(error => {
                    console.log("ERROR",error)
                    return 0;
                });
        }
    }

    const emptyString = (str : string) : boolean => {
        console.log("validMatch: ", validMatch)
        console.log("validPwd: ", validPwd)
        console.log("validMatch: ", validMatch)
        if (str == "") {
            console.log("str: empty")
            return false
        } else {
            return true
        }
    }

    return (
        <div className="register">
            {success ? (
                <section>
                    <h1>Success!</h1>
                    <p>
                        <button className="link__button">
                            <a href="#" >Sign In</a>
                        </button>
                    </p>
                </section>
            ) : (
                <section>
                    <p ref={errRef} className={errMsg ? "errmsg" : "offscreen"} aria-live="assertive">{errMsg}</p>
                    <h1>Register</h1>
                    <form onSubmit={handleSubmit}>
                        <label htmlFor="username">
                            Username:
                            <CheckIcon className={validUsername ? "valid" : "hide"} />
                            <CloseIcon className={validUsername || !validUsername ? "hide" : "invalid"} />
                        </label>
                        <input
                            type="text"
                            id="username"
                            ref={userRef}
                            autoComplete="off"
                            onChange={(e) => setUsername(e.target.value)}
                            value={username}
                            required
                            aria-invalid={validUsername ? "false" : "true"}
                            aria-describedby="uidnote"
                            onFocus={() => setUsernameFocus(true)}
                            onBlur={() => setUsernameFocus(false)}
                        />
                        <p id="uidnote" className={usernameFocus && username && !validUsername ? "instructions" : "offscreen"}>
                            <InfoIcon/>
                            4 to 24 characters.<br />
                            Must begin with a letter.<br />
                            Letters, numbers, underscores, hyphens allowed.
                        </p>

                        <label htmlFor="email">
                            Email:
                            <CheckIcon className={validEmail ? "valid" : "hide"} />
                            <CloseIcon className={validEmail || !email ? "hide" : "invalid"} />
                        </label>
                        <input
                            type="text"
                            id="email"
                            ref={userRef}
                            autoComplete="off"
                            onChange={(e) => setEmail(e.target.value)}
                            value={email}
                            required
                            aria-invalid={validUsername ? "false" : "true"}
                            aria-describedby="uidnote"
                            onFocus={() => setEmailFocus(true)}
                            onBlur={() => setEmailFocus(false)}
                        />
                        <p id="uidnote" className={emailFocus && email && !validEmail ? "instructions" : "offscreen"}>
                            <InfoIcon/>
                            Must be a valid email address of the format<br />
                            blahblah@example.com<br />
                            Letters, numbers, underscores, hyphens allowed.
                        </p>


                        <label htmlFor="password">
                            Password:
                            <CheckIcon className={validPwd ? "valid" : "hide"} />
                            <CloseIcon className={validPwd || !validUsername ? "hide" : "invalid"} />
                        </label>
                        <input
                            type="password"
                            id="password"
                            onChange={(e) => setPwd(e.target.value)}
                            value={pwd}
                            required
                            aria-invalid={validPwd ? "false" : "true"}
                            aria-describedby="pwdnote"
                            onFocus={() => setPwdFocus(true)}
                            onBlur={() => setPwdFocus(false)}
                        />
                        <p id="pwdnote" className={pwdFocus && !validPwd ? "instructions" : "offscreen"}>
                            <InfoIcon/>
                            8 to 24 characters.<br />
                            Must include uppercase and lowercase letters, <br/>
                            a number and a special character.<br />
                            Allowed special characters: <span aria-label="exclamation mark">!</span> <span aria-label="at symbol">@</span> <span aria-label="hashtag">#</span> <span aria-label="dollar sign">$</span> <span aria-label="percent">%</span>
                        </p>


                        <label htmlFor="confirm_pwd">
                            Confirm Password:
                            <CheckIcon className={validMatch && validPwd ? "valid" : "hide"} />
                            <CloseIcon className={(validMatch) || (!emptyString(matchPwd) && !emptyString(pwd)) ? "hide" : "invalid"} />
                            {/* false false -> true and true */}
                        </label>
                        <input
                            type="password"
                            id="confirm_pwd"
                            onChange={(e) => setMatchPwd(e.target.value)}
                            value={matchPwd}
                            required
                            aria-invalid={validMatch ? "false" : "true"}
                            aria-describedby="confirmnote"
                            onFocus={() =>  setMatchFocus(true)}
                            onBlur={() => setMatchFocus(false)}
                        />
                        <p id="confirmnote" className={matchFocus && !validMatch ? "instructions" : "offscreen"}>
                            <InfoIcon/>
                            Must match the first password input field.
                        </p>

                        <button disabled={!validUsername || !validPwd || !validMatch ? true : false}>Sign Up</button>
                    </form>
                    <p>
                        Already registered?<br />
                        <span className="line">
                            {/*put router link here*/}
                            <a href="#">Sign In</a>
                        </span>
                    </p>
                        {backendMessage != '' ? (
                            <div className="backendMessage">
                                <p>{backendMessage}</p>
                            </div>  
                        ) : null}
                </section>
            )}
        </div>
    )
}

export default Register;
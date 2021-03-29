import React, {useState} from 'react';
import './login.css';

const Login = () => {
    // initial state
    const [state, setState] = useState<any>(null);
    const [isError, setIsError] = useState<boolean>(false);
    const url = '#';

    // helper
    const handleChange = (key: string, value: string) => {
        setIsError(false);
        setState({...state, [key]: value})
    }

    const handleLogin = (event: any) => {
        event.preventDefault();
        setIsError(prevState => !prevState);
    }

    return (
        <div className='login-container'>
            <div className='login-wrapper'>
                <div className='app-icon'>
                    <img src='/images/login/doge.png' alt='Open Your World'/>
                </div>
                <div className='input-wrapper'>
                    <label>Username</label>
                    <input onChange={(e) => handleChange('username', e.target.value)}/>
                </div>
                <div className='input-wrapper mt-20'>
                    <label>Password</label>
                    <input type='password' onChange={(e) => handleChange('pwd', e.target.value)}/>
                </div>
                <div className='help-link'>
                    <a href={url}>
                        Need help?
                    </a>
                </div>
                <div className={isError ? 'error-message error-active' : 'error-message'}>
                    Invalid username or password,<br/>please try again.
                </div>
                <a className='btn' href={url} onClick={(e) => handleLogin(e)}>
                    Login
                </a>
            </div>
        </div>
    )
}

export default Login
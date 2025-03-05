import React, { ChangeEvent, useState } from "react"
import axios from "axios"

interface SignupProp {
    username: string,
    password: string,
    email: string,
}


const Signup:React.FC = () => {
    const [signupData, setSignupData] = useState<SignupProp>({username: "", email: "", password: ""})

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        setSignupData({ ...signupData, [e.target.name]: e.target.value });
    }

    const handleSignup = async () => {
        try {
            console.log(signupData)
            const response = await axios.post("http://localhost:8080/addUser", signupData);
            console.log(response);
        } catch (error) {
            console.log("problem" + error)
        }
    }

    return (
        <div>
            <div>
            <form onSubmit={handleSignup}>
        <div>
          <label>Name:</label>
          <input type="text" name="username" value={signupData.username} onChange={handleChange} required />
        </div>

        <div>
          <label>Email:</label>
          <input type="email" name="email" value={signupData.email} onChange={handleChange} required />
        </div>

        <div>
          <label>Password:</label>
          <input type="password" name="password" value={signupData.password} onChange={handleChange} required />
        </div>

        <button type="submit" style={{ marginTop: "10px", padding: "10px", cursor: "pointer" }}>
          Sign Up
        </button>
        </form>
            </div>
        </div>
    )
}

export default Signup;
'use client';

import Link from "next/link"
import "../../../style/login.css"

export default function register() {

  const submitForm = async (e) => {
    e.preventDefault()

    const formData = new FormData(e.target)
    try {

      const resp = await fetch('http://localhost:8080/register', {
        method: 'POST',
        body: formData
      })

      const data = await resp.json()

      console.log(data)
      
    } catch (error) {
      console.error(error)
    }

    
      
  }

  return (
   
      <div id="register-container" >
        <div className="info-side">
          <h2>Create an account</h2>
          <p>Join us and enjoy all the benefits of our platform</p>
          <ul className="feature-list">
            <li>Customer Service 24/7</li>
            <li>Interface simple et intuitive</li>
            <li>Protection of your personal data</li>
            <li>Regular feature updates</li>
          </ul>
        </div>

        <div className="register">
          <h1>Create Your Account</h1>
          <form id="register-form" onSubmit={submitForm}>
            <div className="name-row">
              <div className="form-group">
                <label htmlFor="firstName">First Name</label>
                <input type="text" id="firstName" placeholder="John" />
              </div>
              <div className="form-group">
                <label htmlFor="lastName">Last Name</label>
                <input type="text" id="lastName" placeholder="Doe" />
              </div>
            </div>

            <div className="form-group">
              <label htmlFor="age">Age</label>
              <input type="number" id="age" placeholder="25" />
            </div>

            <div className="form-group">
              <label htmlFor="gender">Gender</label>
              <select id="gender" defaultValue="">
                <option value="" disabled>
                  Select gender
                </option>
                <option value="male">Male</option>
                <option value="female">Female</option>
              </select>
            </div>

            <div className="form-group">
              <label htmlFor="nickname">Nickname</label>
              <input type="text" id="nickname" placeholder="johndoe" />
            </div>

            <div className="form-group">
              <label htmlFor="email">Email Address</label>
              <input type="email" id="email" placeholder="john@example.com" />
            </div>

            <div className="form-group">
              <label htmlFor="password">Password</label>
              <input type="password" id="password" placeholder="••••••••" />
            </div>

            <div className="fill">
              <span>Fill in all fields</span>
            </div>

            <p id="error-reg"></p>

            <button type="submit" id="creat-btn">Create Account</button>
            <span className="have">Already have an account?</span>
            <Link href="/login" id="log">Login</Link>
          </form>
        </div>
      </div>
  )
}
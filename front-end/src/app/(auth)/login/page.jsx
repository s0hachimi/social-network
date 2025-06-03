'use client';

import Link from "next/link";
import "../../../style/login.css"
import { useRouter } from "next/navigation";

export default function Login() {

  const router = useRouter()


  const submitForm = async (e) => {
    e.preventDefault()

    const msgErr = document.getElementById('error-log')

    const formData = new FormData(e.target)


    const json = {
      Email: formData.get("email"),
      Password: formData.get("password")
    }

    if (!json.Password || !json.Email) {
      msgErr.textContent = "Please fill up fields"
      msgErr.style.color = "red"
      return
    }

    try {

      const resp = await fetch('http://localhost:8080/login', {
        method: 'POST',
        credentials: "include",
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(json)
      })

      const data = await resp.json()

      console.log(data);

      if (data.status) {
        router.push("/Posts")
      } else {
        msgErr.textContent = data.error
        msgErr.style.color = "red"
      }

    } catch (error) {
      console.error(error)
    }

  }


  return (
    <div>
      <div id="login-container">
        <div className="info-side">
          <h2>Welcome back!</h2>
          <p>Log in to access your account</p>
          <p>Take advantage of all our exclusive services and features.</p>
        </div>

        <div className="login-form">
          <h1>Login</h1>
          <form id="login-form" onSubmit={submitForm} >
            <div className="form-group">
              <label>Nickname / Email</label>
              <input
                type="text"
                id="userInput"
                name="email"
                placeholder="Nickname or Email"
              />
            </div>

            <div className="form-group">
              <label htmlFor="paswd">Password</label>
              <input
                type="password"
                id="paswd"
                name="password"
                placeholder="••••••••"
                required
              />
            </div>

            <p id="error-log"></p>
            <button type="submit" id="login-btn">Login</button>

            <div className="register-link">
              No account yet ?
              <Link href="/register" id="resgesterlogin">
                Create an account
              </Link>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}


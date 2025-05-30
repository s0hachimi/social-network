import Link from "next/link";
import "../../../style/login.css"

export default function Login() {

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
          <form id="login-form" method="post">
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
              Pas encore de compte?{" "}
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


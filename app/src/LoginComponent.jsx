import React from 'react';

function LoginComponent() {
  return (
    <div>
      <h1>Login Page</h1>
      <form action="/login" method="post">
        <label htmlFor="username">Username:</label>
        <input type="text" id="username" name="username" /><br /><br />
        <label htmlFor="password">Password:</label>
        <input type="password" id="password" name="password" /><br /><br />
        <button type="submit">Login</button>
      </form>
    </div>
  );
}

export default LoginComponent;
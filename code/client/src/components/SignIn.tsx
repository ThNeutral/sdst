import React from "react";
import "../css/SignIn.css"; // Import the CSS file here

const SignIn: React.FC = () => {
  return (
    <div className="signin-container">
      {/* Header */}
      <header className="signin-header">
        <div className="logo">
          <img
            src="/icons/PlasmaLogo.png"
            alt="Plasma Logo"
            className="logo-image"
          />
          <h1 className="header-title">Collaborative Editor Plasma</h1>
        </div>
      </header>

      {/* Create Account Section */}
      <main className="create-account-section">
        <h2 className="create-account-title">Create an Account</h2>
        <p className="instruction-text">
          Please enter your details to sign up and start collaborating!
        </p>

        {/* Input Fields */}
        <form className="input-fields">
          <input type="text" placeholder="Username" className="input-field" />
          <input
            type="email"
            placeholder="Email Address"
            className="input-field"
          />
          <input
            type="password"
            placeholder="Password"
            className="input-field"
          />
          <button type="submit" className="submit-button">
            Sign Up
          </button>
        </form>
      </main>
    </div>
  );
};

export default SignIn;

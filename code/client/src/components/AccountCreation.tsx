import React from "react";
import "../css/AccountCreation.css"; // Import the styles for this page

const AccountCreation: React.FC = () => {
  return (
    <div className="account-creation-container">
      <div className="header">
        <div className="logo">
          <div className="logo-img"></div>
        </div>
        <p className="header-title">Collaborative Editor Plasma</p>
      </div>

      <div className="form-container">
        <p className="form-title">Creating Account</p>
        <p className="form-instruction">Please fill in the details below</p>

        {/* Input Fields */}
        <div className="input-field">
          <input type="text" placeholder="Username" />
        </div>
        <div className="input-field">
          <input type="email" placeholder="Email" />
        </div>
        <div className="input-field">
          <input type="password" placeholder="Password" />
        </div>

        {/* Submit Button */}
        <button className="submit-button">Create Account</button>
      </div>
    </div>
  );
};

export default AccountCreation;

import React from "react";
import "../css/AccountCreationErrorCase.css"; // Import the specific styles for this page

const AccountCreationErrorCase: React.FC = () => {
  return (
    <div className="error-container">
      <div className="header">
        <div className="logo">
          <div className="logo-img"></div>
        </div>
        <p className="header-title">Collaborative Editor Plasma</p>
      </div>

      <div className="form-container">
        <p className="error-message">
          An error occurred while creating your account. Please try again.
        </p>

        {/* Error Input Field */}
        <div className="input-field">
          <input type="text" placeholder="Username" />
        </div>

        {/* Submit Button */}
        <button className="submit-button">Retry</button>
      </div>
    </div>
  );
};

export default AccountCreationErrorCase;

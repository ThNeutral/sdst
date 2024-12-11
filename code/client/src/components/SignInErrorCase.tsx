import React from "react";
import "../css/SignInErrorCase.css"; // Ensure CSS is imported

const ErrorContainer: React.FC = () => {
  return (
    <div className="error-container">
      {/* Header */}
      <header className="error-header">
        <div className="logo">
          <img
            src="—Pngtree—magic plasma ball abstract translucent_15853701.png"
            alt="Plasma Logo"
            className="logo-image"
          />
          <h1 className="header-title">Collaborative Editor Plasma</h1>
        </div>
      </header>

      {/* Error Input Field Section */}
      <div className="error-input-field-container">
        <div className="input-field-error">
          <input
            type="text"
            placeholder="Enter some data..."
            className="error-input" // Connects with the CSS
          />
        </div>
        <p className="error-message">This field is required.</p>{" "}
        {/* Error message */}
      </div>
    </div>
  );
};

export default ErrorContainer;

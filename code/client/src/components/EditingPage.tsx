import React, { useState, useEffect } from "react";
import "../css/EditingPage.css";

const EditingPage: React.FC = () => {
  // Retrieve theme from localStorage or default to "dark"
  const [theme, setTheme] = useState<string>(
    localStorage.getItem("theme") || "dark"
  );

  // Update theme in localStorage and document whenever the theme changes
  useEffect(() => {
    document.documentElement.setAttribute("data-theme", theme); // Apply theme to <html>
    localStorage.setItem("theme", theme); // Save theme to localStorage
  }, [theme]); // This effect runs whenever the theme changes

  const toggleTheme = () => {
    // Toggle between "light" and "dark" themes
    setTheme((prevTheme) => (prevTheme === "dark" ? "light" : "dark"));
  };

  return (
    <div className="editing-page">
      {/* Header */}
      <header className="header">
        <div className="logo"></div>
        <button className="run-button">Run</button>
      </header>

      {/* Terminal Button */}
      <div className="terminal-button">
        <button className="button">
          <div className="terminal-icon"></div>
          Terminal
        </button>
      </div>

      {/* Invite Button */}
      <div className="invite-button">
        <button className="button">
          <div className="user-icon"></div>
          Invite
        </button>
      </div>

      {/* Theme Toggle */}
      <div className="theme-toggle">
        <button onClick={toggleTheme}>
          {theme === "dark" ? "Switch to Light" : "Switch to Dark"}
        </button>
      </div>

      {/* Main content */}
      <div className="tv">
        <div className="tv-content">{/* Your content goes here */}</div>
      </div>

      {/* Search Input */}
      <div className="search-input">
        <label>Search:</label>
        <input type="text" placeholder="Search" />
      </div>
    </div>
  );
};

export default EditingPage;

import { Outlet } from "react-router-dom";
import React from "react";
import { useTheme } from "../context/ThemeContext";

export const Header: React.FC = () => {
  const { theme, toggleTheme } = useTheme();

  const isEditorRoute = window.location.pathname === "/editor";

  return (
    <>
      <header className={isEditorRoute ? "header-editor" : "header"}>
        {!isEditorRoute ? (
          <>
            <p className="header-text-small">collaborative editor</p>
            <p className="header-text-big">PLASMA</p>
            <button onClick={toggleTheme} className="theme-toggle-button">
              {theme === "dark"
                ? "Switch to Light Mode"
                : "Switch to Dark Mode"}
            </button>
          </>
        ) : (
          <>
            <p className="header-text-big-editor">repo name</p>
            <button className="header-button-run">
              <div className="header-button-run-circle"></div>
              <div className="header-button-run-text">RUN</div>
            </button>
            <button className="header-button-invite">
              <div className="header-button-run-text">Invite</div>
            </button>
          </>
        )}
      </header>
      <Outlet />
    </>
  );
};

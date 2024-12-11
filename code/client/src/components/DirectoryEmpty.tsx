import React from "react";
import "../css/DirectoryEmpty.css";

const DirectoryEmpty: React.FC = () => {
  return (
    <div className="directory-empty">
      {/* Header */}
      <header className="header">
        <div className="logo"></div>
        <div className="plasma">PLASMA</div>
      </header>

      {/* Plasma Editor */}
      <div className="plasma-editor">Where ideas ignite</div>

      {/* Main Content */}
      <div className="tv"></div>

      {/* Menu Options */}
      <div className="menu-options">
        <div className="new-file">New file...</div>
        <div className="open-file">Open file...</div>
        <div className="open-folder">Open folder...</div>
        <div className="clone-git-repo">Clone Git Repository...</div>
      </div>

      {/* Icons (Add, Import, etc.) */}
      <div className="icons">
        <div className="add-icon"></div>
        <div className="import-icon"></div>
        <div className="git-icon"></div>
        <div className="folder-icon"></div>
      </div>

      {/* Gradient or Decorative Element */}
      <div className="gradient-element"></div>
    </div>
  );
};

export default DirectoryEmpty;

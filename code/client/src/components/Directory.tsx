import React from "react";
import "../css/Directory.css"; // Make sure to import the updated CSS

const DirectoryEmpty: React.FC = () => {
  return (
    <div className="directory-empty">
      {/* Header */}
      <div className="header">
        <div className="logo"></div>
        <div className="plasma">Plasma</div>
      </div>

      {/* TV Section */}
      <div className="tv">
        <div className="tv-content">{/* Your content goes here */}</div>
      </div>

      {/* Buttons Section */}
      <div className="buttons">
        <div className="button add">Add</div>
        <div className="button import">Import</div>
        <div className="button folder">Folder</div>
        <div className="button git">Git</div>
      </div>

      {/* File Options */}
      <div className="file-options">
        <div className="file-option">New File...</div>
        <div className="file-option">Open File...</div>
        <div className="file-option">Open Folder...</div>
        <div className="file-option">Clone Git Repository...</div>
      </div>
    </div>
  );
};

export default DirectoryEmpty;

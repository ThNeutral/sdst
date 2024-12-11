import React from "react";
import "../css/ErrorPage.css";

const Error404Page: React.FC = () => {
  return (
    <div className="error-page">
      <div className="error-code">404</div>
      <div className="error-message">THE PAGE NOT FOUND</div>
    </div>
  );
};

export default Error404Page;

import { Outlet } from "react-router-dom";

export function Header() {
  return (
    <>
      {window.location.pathname != "/editor" ? (
        <div className="header">
          <p className="header-text-small">collaborative editor</p>
          <p className="header-text-big">PLASMA</p>
        </div>
      ) : (
        <div className="header-editor">
          <p className="header-text-big-editor">repo name</p>
          <button className="header-button-run">
            <div className="header-button-run-circle"></div>
            <div className="header-button-run-text">RUN</div>
          </button>
          <button className="header-button-invite">
            <div className="header-button-run-text">Invite</div>
          </button>
        </div>
      )}
      <Outlet />
    </>
  );
}

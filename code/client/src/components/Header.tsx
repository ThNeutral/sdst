import { Outlet } from "react-router-dom";

export function Header() {
  return (
    <>
      <div className="header">
        <p className="header-text-small">collaborative editor</p>
        <p className="header-text-big">PLASMA</p>
      </div>
      <Outlet />
    </>
  );
}

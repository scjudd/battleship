import { Outlet, Link } from 'react-router-dom';

function Layout() {
  return (
    <>
      <nav className="text-white p-4 font-mono">
        <span>Battleship</span>
      </nav>

      <Outlet />
    </>
  );
}

export default Layout;

import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useSelector } from 'react-redux';

const SideMenu = () => {
  const [isOpen, setIsOpen] = useState(false);

  const pathname = useSelector(({ router }) => router.location.pathname);

  useEffect(() => setIsOpen(false), [pathname]);

  const links = [
    {
      label: '非同期通信の確認画面',
      url: '/',
    },
    {
      label: 'Not Found',
      url: 'not_found',
    },
  ];

  const linksList = links => {
    const linkArray = links.map(item => (
      <li key={item.url}>
        <Link to={item.url}>{item.label}</Link>
      </li>
    ));
    return linkArray;
  };

  return (
    <aside className="fade-in">
      <nav>
        <input
          type="checkbox"
          className="check"
          id="checked"
          checked={isOpen}
          onChange={() => setIsOpen(!isOpen)}
        />
        <label className="menu-btn" htmlFor="checked">
          <span
            className={
              isOpen
                ? 'hamburger hamburger--3dx is-active'
                : 'hamburger hamburger--3dx'
            }
          >
            <span className="hamburger-box">
              <span className="hamburger-inner" />
            </span>
          </span>
        </label>
        <label className="close-menu" htmlFor="checked" />
        <div className="drawer-menu">
          <ul className="sub">{linksList(links)}</ul>
        </div>
      </nav>
    </aside>
  );
};

export default SideMenu;

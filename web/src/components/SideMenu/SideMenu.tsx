import * as React from 'react';
import { Link } from 'react-router-dom';
import { LinksPathnames } from '../../constant/commonConstant';

import './style.scss';

type LinkType = {
  label: string;
  url: string;
};
export const SideMenu = (): JSX.Element => {
  const [isOpen, setIsOpen] = React.useState(false);

  const links = [
    {
      label: LinksPathnames.HOME.name,
      url: LinksPathnames.HOME.path,
    },
    {
      label: LinksPathnames.IMAGE.name,
      url: LinksPathnames.IMAGE.path,
    },
    {
      label: LinksPathnames.TEST.name,
      url: LinksPathnames.TEST.path,
    },
  ];

  const linksList = (links: LinkType[]) => {
    const linkArray = links.map(item => (
      <li key={item.url}>
        <Link to={item.url}>{item.label}</Link>
      </li>
    ));
    return linkArray;
  };

  return (
    <aside>
      <nav>
        <input
          type="checkbox"
          className="check"
          id="checked"
          checked={isOpen}
          onClick={() => setIsOpen(!isOpen)}
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

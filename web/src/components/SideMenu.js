import React from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { Link } from "react-router-dom";
import { sideMenuAction } from "../actions/sideMenuAction";

export class SideMenu extends React.Component {
  componentWillUnmount() {
    this.props.initialize();
  }

  render() {
    const { initialize, isSideOpen, changeState } = this.props;

    const hamburgerClassName = isSideOpen
      ? "hamburger hamburger--3dx is-active"
      : "hamburger hamburger--3dx";
    const droawerClassName = isSideOpen
      ? "drawer-menu drawer-open"
      : "drawer-menu";

    const links = [
      {
        label: "Top",
        url: "/"
      },
      {
        label: "Not Found",
        url: "not_found"
      }
    ];

    const linksList = links => {
      const linkArray = links.map(item => (
        <li key={item.url} onClick={() => initialize()}>
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
            onChange={() => changeState({ isSideOpen: !isSideOpen })}
          />
          <label className="menu-btn" htmlFor="checked">
            <span className={hamburgerClassName}>
              <span className="hamburger-box">
                <span className="hamburger-inner" />
              </span>
            </span>
          </label>
          <label className="close-menu" htmlFor="checked" />
          <div className={droawerClassName}>
            <ul className="sub">{linksList(links)}</ul>
          </div>
        </nav>
      </aside>
    );
  }
}

const mapStateToProps = ({ sideMenu }) => sideMenu;
const mapDispatchToProps = dispatch =>
  bindActionCreators(sideMenuAction, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(SideMenu);

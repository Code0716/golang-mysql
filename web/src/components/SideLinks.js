import React from "react";
import PropTypes from "prop-types";

export class SideLink extends React.Component {
  render() {
    return (
      <ul className="sub">
        <li>
          <a href="/">Top</a>
        </li>
        <li>
          <a href="Member">Member</a>
        </li>
      </ul>
    );
  }
}

//-------------------------------------------------------------------------

export default SideLink;

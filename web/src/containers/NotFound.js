import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { Link } from 'react-router-dom';

class NotFound extends React.Component {
  render() {
    return (
      <React.Fragment>
        <div className="contents fade-in">
          <div className="form_container">
            <h1>ページが見つかりません。</h1>
            <Link to="/">Top</Link>
          </div>
        </div>
      </React.Fragment>
    );
  }
}
export default NotFound;

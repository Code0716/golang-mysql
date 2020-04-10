import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { homeActions } from '../actions/homeActions';

export class Home extends React.Component {
  render() {
    const { getData, postData, changeEntity, get, post } = this.props;
    const { getUrl, postUrl, postEntity } = this.props.entity;
    return (
      <React.Fragment>
        <div className="contents fade-in">
          <div className="form_container">
            {getData}
            {postData}
            <h1>inputサンプル</h1>
            <label>GET</label>
            <input
              value={getUrl}
              type="text"
              onChange={e => changeEntity({ getUrl: e.target.value })}
            ></input>
            <button className="action_button" onClick={() => get(getUrl)}>
              <span>GET button</span>
            </button>
            <label>POST</label>
            <input
              value={postUrl}
              type="text"
              onChange={e => changeEntity({ postUrl: e.target.value })}
            ></input>
            <textarea
              style={{
                width: '500px',
                height: '300px',
                display: 'block',
                margin: '0 auto',
              }}
              value={postEntity}
              onChange={e => changeEntity({ postEntity: e.target.value })}
            />
            <button
              className="action_button"
              onClick={() => post(postUrl, postEntity)}
            >
              <span>POST button</span>
            </button>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

//-------------------------------------------------------------------------

const mapStateToProps = ({ home }) => home;
const mapDispatchToProps = dispatch =>
  bindActionCreators(homeActions, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(Home);

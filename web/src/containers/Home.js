import React, { useState, useMemo } from 'react';
import { homeActions } from '../actions/homeActions';

const Home = () => {
  const [getInput, setGetInput] = useState('');
  const [postInput, setPostInput] = useState('');
  const [textArea, setTextArea] = useState('');

  const { getData, postData, get, post } = homeActions();

  const country = useMemo(() => {
    return getData.map((val, index) => <li key={index}>{val.name}</li>);
  }, [getData]);

  return (
    <React.Fragment>
      <div className="contents fade-in">
        <div className="form_container">
          <div>非同期通信の確認画面</div>
          {country}
          {postData}
          <h1>inputサンプル</h1>
          <label>GET</label>
          <input
            value={getInput}
            type="text"
            onChange={e => setGetInput(e.target.value)}
          />
          <button className="action_button" onClick={() => get(getInput)}>
            <span>GET button</span>
          </button>
          <label>POST</label>
          <input
            value={postInput}
            type="text"
            onChange={e => setPostInput(e.target.value)}
          ></input>
          <textarea
            style={{
              width: '500px',
              height: '300px',
              display: 'block',
              margin: '0 auto',
            }}
            value={textArea}
            onChange={e => setTextArea(e.target.value)}
          />
          <button
            className="action_button"
            onClick={() => post(postInput, textArea)}
          >
            <span>POST button</span>
          </button>
        </div>
      </div>
    </React.Fragment>
  );
};

export default Home;

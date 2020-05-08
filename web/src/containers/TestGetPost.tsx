import React, { useState } from 'react';
import { AutoSizer, Column, Table } from 'react-virtualized';
import { Select } from '../components/Select';
import { FormContainer } from '../components/FormContainer';
import { testGetPostActions } from '../actions/testGetPostActions';
import { shotMessageAction } from '../actions/shotMessageActions';

const TestGetPost = () => {
  const [getInput, setGetInput] = useState('/city');
  const [postInput, setPostInput] = useState('');
  const [textArea, setTextArea] = useState('');

  const { getData, postData, get, post } = testGetPostActions();
  const { showMessage } = shotMessageAction();

  const selctData: { value: string }[] = [
    { value: '/country' },
    { value: '/country/japan' },
    { value: '/city' },
    { value: '/city/tokyo' },
    { value: '/city/kyoto' },
  ];

  const disabled: boolean = true;

  return (
    <FormContainer>
      <div className="test_center">
        <div>非同期通信の確認画面</div>
        {postData}
        <h1>inputサンプル</h1>
        <label>GET</label>
        <Select
          value={getInput}
          option={selctData}
          onChange={e => setGetInput(e.target.value)}
        />
        <button className="action_button" onClick={() => get(getInput)}>
          <span>GET button</span>
        </button>
        <label>POST(TODO)</label>
        <input
          value={postInput}
          type="text"
          onChange={() => {}}
          disabled={disabled}
        />
        <textarea
          className="test_textarea"
          disabled={disabled}
          value={textArea}
          onChange={e => setTextArea(e.target.value)}
        />
        <button
          className="action_button"
          onClick={() => post(postInput, textArea)}
          disabled={disabled}
        >
          <span>POST button</span>
        </button>
      </div>
      <button onClick={() => showMessage('test')}>test</button>
      <button onClick={() => showMessage('rest')}>rest</button>
      <AutoSizer>
        {({ width, height }) => (
          <Table
            data={getData}
            height={height}
            width={width}
            rowHeight={35}
            rowGetter={({ index }) => getData[index]}
            rowCount={getData.length}
            rowClassName="d-flex"
          >
            <Column width={200} flexGrow={1} label="name" dataKey="name" />
            <Column width={200} label="code" dataKey="code" />
          </Table>
        )}
      </AutoSizer>
    </FormContainer>
  );
};

export default TestGetPost;

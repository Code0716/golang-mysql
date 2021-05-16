import * as React from 'react';
import { useState } from 'react';
import { Select } from '../components/Select';
import { FormContainer } from '../components/FormContainer';
import { testGetPostActions } from '../actions/testGetPostActions';
import { shotMessageAction } from '../actions/shotMessageActions';

const TestGetPost = () => {
  const [getInput, setGetInput] = useState('/city');

  const { postData, get, migrate } = testGetPostActions();
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
      <div className="test_center mb-30">
        <div>非同期通信の確認画面</div>
        {postData}
        <h1>inputサンプル</h1>
        <label>GET</label>
        <Select
          value={getInput}
          option={selctData}
          onChange={(e: React.ChangeEvent<HTMLSelectElement>) =>
            setGetInput(e.target.value)
          }
        />
        <button className="action_button" onClick={() => get(getInput)}>
          <span>GET button</span>
        </button>
      </div>
      <div className="test_center mb-30">
        <label>Notification</label>
        <div className="d-flex mb-15">
          <button
            className="action_button mr-5"
            onClick={() => showMessage('test')}
          >
            test
          </button>
          <button className="action_button" onClick={() => showMessage('rest')}>
            rest
          </button>
        </div>
      </div>
      <div className="test_center mb-30">
        <label>migration</label>
        <button className="action_button" onClick={migrate}>
          migrate
        </button>
      </div>
    </FormContainer>
  );
};

export default TestGetPost;

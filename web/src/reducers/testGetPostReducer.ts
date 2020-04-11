import { ActionTypes } from '../actions/testGetPostActions';

export interface State {
  getData: Array<Object>;
  postData: Array<Object>;
}
const initialState: State = {
  getData: [],
  postData: [],
};

export function testGetPostReducer(state = initialState, { type, payload }) {
  switch (type) {
    case ActionTypes.CHANGE_STATE:
      return {
        ...state,
        ...payload,
      };

    default:
      return state;
  }
}

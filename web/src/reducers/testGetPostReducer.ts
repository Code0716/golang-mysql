import { ActionTypes } from '../actions/testGetPostActions';

export interface State {
  getData: { code: string; name: string }[];
  postData: { id: number; code: string; name: string }[];
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

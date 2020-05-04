import { ActionTypes } from '../actions/testGetPostActions';

export type State = {
  getData: { code: string; name: string }[];
  postData: { id: number; code: string; name: string }[];
  entity: {};
};

export const initialState: State = {
  getData: [],
  postData: [],
  entity: {},
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

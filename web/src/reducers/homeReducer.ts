import { ActionTypes } from '../actions/homeActions';

export interface State {
  getData: Array<Object>;
  postData: Array<Object>;
}
const initialState: State = {
  getData: [],
  postData: [],
};

function homeReducer(state = initialState, { type, payload }) {
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

export default homeReducer;

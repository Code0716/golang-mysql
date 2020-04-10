import { CHANGE_STATE } from '../actions/homeActions';

const initialState = {
  getData: [],
  postData: [],
};

function homeReducer(state = initialState, { type, payload }) {
  switch (type) {
    case CHANGE_STATE:
      return {
        ...state,
        ...payload,
      };

    default:
      return state;
  }
}

export default homeReducer;

import {
  HOME_GET,
  HOME_POST,
  CHANGE_STATE,
  CHANGE_ENTITY
} from "../actions/homeActions";

const initialState = {
  getData: "",
  postData: "",
  entity: {
    getUrl: "",
    postUrl: "",
    postEntity: ""
  }
};

function homeReducer(state = initialState, { type, payload }) {
  switch (type) {
    case CHANGE_STATE:
      return {
        ...state,
        ...payload
      };
    case CHANGE_ENTITY:
      return {
        ...state,
        entity: { ...state.entity, ...payload }
      };
    default:
      return state;
  }
}

export default homeReducer;

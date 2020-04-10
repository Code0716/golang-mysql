import {
  INITALIZE,
  CHANGE_STATE,
  CHANGE_ENTITY
} from "../actions/sideMenuAction";

const initialState = {
  isSideOpen: false
};

function reducer(state = initialState, { type, payload }) {
  switch (type) {
    case INITALIZE:
      return { ...initialState };
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

export default reducer;

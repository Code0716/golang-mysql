import { ActionTypes } from '../actions/imageListActions';

// 型の定義
export interface State {
  images: { continent: string }[];
}

// 初期値
const initialState: State = {
  images: [],
};

export function imageListReducer(state = initialState, { type, payload }) {
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

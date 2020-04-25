import { ActionTypes } from '../actions/loadingActions';

// 型の定義
export type loading = {
  load: boolean;
};

const initialState: loading = {
  load: false,
};

// reducer
export function reducer(state = initialState, action) {
  switch (action.type) {
    case ActionTypes.LOADING:
      return { ...state, load: true };
    case ActionTypes.UNLOADING:
      return { ...state, load: false };
    default:
      return state;
  }
}

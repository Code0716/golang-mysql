import { ActionTypes } from '../actions/shotMessageActions';

// 型の定義
export type shotMessProps = {
  show: boolean;
  hide: boolean;
};

const initialState: shotMessProps = {
  show: false,
  hide: true,
};

// reducer
export function reducer(state = initialState, { type }) {
  switch (type) {
    case ActionTypes.SHOW:
      return { ...state, show: true, hide: false };
    case ActionTypes.HIDE:
      return { ...state, hide: true };
    case ActionTypes.DISAPPEAR:
      return { ...state, show: false };
    default:
      return state;
  }
}

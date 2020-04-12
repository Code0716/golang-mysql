import { ActionTypes } from '../actions/slideShowActions';

// 型の定義
export interface State {
  continents: { continent: string }[];
  cities: { code: string; name: string }[];
  countries: { id: number; code: string; name: string; continent: string }[];
  currentContinent: string;
  description: string;
}

// 初期値
const initialState: State = {
  continents: [],
  cities: [],
  countries: [],
  currentContinent: '',
  description: '',
};

export function slideShowReducer(state = initialState, { type, payload }) {
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

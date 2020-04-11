import { ActionTypes } from '../actions/slideShowActions';

export interface State {
  continents: { continent: string }[];
  cities: { code: string; name: string }[];
  countries: { id: number; code: string; name: string; continent: string }[];
}
const initialState: State = {
  continents: [],
  cities: [],
  countries: [],
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

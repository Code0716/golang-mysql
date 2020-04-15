import { ActionTypes } from '../actions/imageListActions';

// 型の定義
export interface State {
  images: { continent: string }[];
  preUploadImages: string[];
}

// 初期値
const initialState: State = {
  images: [],
  preUploadImages: [],
};

export function imageListReducer(state = initialState, { type, payload }) {
  switch (type) {
    case ActionTypes.POST_PRE_UPLOAD:
      return {
        ...state,
        preUploadImages: [...state.preUploadImages, ...payload],
      };
    case ActionTypes.CHANGE_STATE:
      return {
        ...state,
        ...payload,
      };

    default:
      return state;
  }
}

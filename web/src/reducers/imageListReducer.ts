import { ActionTypes } from '../actions/imageListActions';

// 型の定義
export type ImageInfo = {
  id: number;
  title: string;
  create: Date;
  path: string;
};

export interface LoadImage {
  img: string;
  info: ImageInfo;
}

export interface State {
  images: { continent: string }[];
  preUploadImages: LoadImage[];
}

// 初期値
const initialState: State = {
  images: [],
  preUploadImages: [],
};

export function imageListReducer(state = initialState, { type, payload }) {
  switch (type) {
    case ActionTypes.INITIALIZE:
      return initialState;

    case ActionTypes.UPDATE_PRE_UPLOAD:
      return {
        ...state,
        preUploadImages: [...state.preUploadImages, ...payload],
      };
    case ActionTypes.DELETE_PRE_UPLOAD:
      return {
        ...state,
        preUploadImages: [...payload],
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

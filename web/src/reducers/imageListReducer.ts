import { cloneDeep } from 'lodash';
import { ActionTypes } from '../actions/imageListActions';

// 型の定義
export type ImageInfo = {
  ID: number;
  Title: string;
  CreatedAt: Date;
  Path: string;
};

export interface LoadImage {
  img: string;
  info: ImageInfo;
}

export interface State {
  images: { image: string }[];
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
    /*  case ActionTypes.UPDATE_PRE_UPLOAD_IMAGE:
      const _copyData = cloneDeep(state.preUploadImages);

      _copyData.find((item: LoadImage) => {
        if (item.info.ID === payload.id) {
          item.img = payload.img;
        }
      });
      return {
        ...state,
        preUploadImages: [..._copyData],
      };*/
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

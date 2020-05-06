import { cloneDeep } from 'lodash';
import { ActionTypes } from '../actions/imageListActions';

// 型の定義 imageと
export type ImageInfo = {
  ID: number;
  Title: string;
  CreatedAt: Date;
  Path: string;
};

export type EachImage = {
  id: number;
  img: string;
};

export interface LoadImage {
  image: EachImage;
  info: ImageInfo;
}

export interface State {
  images: LoadImage[];
  preUploadImages: LoadImage[];
}

// 初期値
const initialState: State = {
  images: [],
  preUploadImages: [],
};

export function imageListReducer(state = initialState, { type, payload }) {
  switch (type) {
    case ActionTypes.CHANGE_STATE:
      return {
        ...state,
        ...payload,
      };
    case ActionTypes.INITIALIZE:
      return initialState;
    case ActionTypes.UPDATE_PRE_UPLOAD:
      return {
        ...state,
        preUploadImages: [...state.preUploadImages, ...payload],
      };
    case ActionTypes.UPDATE_UPLOAD_IMAGE:
      // imagesに保存するようにする。
      const _copyData = cloneDeep(state.images);
      _copyData.find((item: LoadImage) => {
        if (item.info.ID === payload.id) {
          item.image = payload;
        }
      });

      return {
        ...state,
        images: [..._copyData],
      };
    case ActionTypes.DELETE_PRE_UPLOAD:
      return {
        ...state,
        preUploadImages: [...payload],
      };

    default:
      return state;
  }
}

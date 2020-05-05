import { useEffect, useCallback } from 'react';
import { Action } from 'redux';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';
import { cloneDeep } from 'lodash';
import { RootState } from '../store/store';

import { LoadImage, ImageInfo } from '../reducers/imageListReducer';
import { loading, unloading } from '../actions/loadingActions';

import HttpRequest from '../service/api/HttpRequest';
// Actions
export const ActionTypes = {
  INITIALIZE: 'IMAGE_LIST_INITIALIZE',
  GET_LIST: 'IMAGE_LIST_GET_LIST',
  GET_UPLOAD: 'IMAGE_LIST_GET_UPLOAD_IMAGES',
  UPDATE_PRE_UPLOAD: 'IMAGE_LIST_UPDATE_PRE_UPLOAD',
  UPDATE_PRE_UPLOAD_IMAGE: 'IMAGE_LIST_UPDATE_PRE_UPLOAD_IMAGE',
  DELETE_PRE_UPLOAD: 'IMAGE_LIST_DELETE_PRE_UPLOAD',
  CHANGE_STATE: 'IMAGE_LIST_CHANGE_STATE',
} as const;

// Action Creators

// Actionの型 Actionを継承 TODO
interface GetList extends Action {
  type: typeof ActionTypes.GET_LIST;
}

export type HttpRequestActionTypes = GetList;

export const imageListActions = () => {
  const dispatch = useDispatch();

  //store
  const imageList = useSelector(({ imageList }: RootState) => imageList);

  const { images, preUploadImages } = imageList;

  const initialize = useCallback(
    () =>
      dispatch({
        type: ActionTypes.INITIALIZE,
      }),
    [dispatch],
  );

  //  画像一覧
  const getImages = useCallback(async () => {
    dispatch({
      type: ActionTypes.GET_LIST,
    });
    try {
      const data = await HttpRequest.get(`/image/upload`);
      dispatch({
        type: ActionTypes.CHANGE_STATE,
        payload: { images: data },
      });
    } finally {
      // TODO
    }
  }, [dispatch]);

  // load preupload image
  const getPreImagesInfo = useCallback(async () => {
    dispatch(loading());
    try {
      const response = await HttpRequest.get('/image/pre_upload');
      dispatch({
        type: ActionTypes.UPDATE_PRE_UPLOAD,
        payload: response,
      });
    } finally {
      dispatch(unloading());
    }
  }, [dispatch, preUploadImages]);

  // 画像を一枚づつ取得する。
  const getImage = useCallback(
    async (id: Number) => {
      try {
        const response = await HttpRequest.get(`/image/pre_upload/${id}`);

        dispatch({
          type: ActionTypes.UPDATE_PRE_UPLOAD_IMAGE,
          payload: { id: id, ...response },
        });
      } finally {
      }
    },
    [dispatch, preUploadImages],
  );

  // add preupload image
  const addPreUploadImages = useCallback(
    async (files: File[]) => {
      dispatch(loading());

      const postData = new FormData();
      files.forEach(element => {
        postData.append('images', element);
      });

      try {
        const response = await HttpRequest.postImg(
          '/image/pre_upload',
          postData,
        );
        dispatch({
          type: ActionTypes.UPDATE_PRE_UPLOAD,
          payload: response.data,
        });
      } finally {
        dispatch(unloading());
      }
    },
    [dispatch, preUploadImages],
  );

  //  PreUpload delete image
  const deletePreImage = useCallback(
    async (id: number) => {
      try {
        const _copyImages = cloneDeep(preUploadImages);
        const newState: LoadImage[] = [];
        let deleteImage: ImageInfo;

        _copyImages.forEach((elm: LoadImage) => {
          if (elm.info.ID === id) {
            delete elm.img;
            deleteImage = elm.info;
          } else {
            newState.push(elm);
          }
        });

        await HttpRequest.deleteImage(`/image/pre_upload/delete`, deleteImage);
        dispatch({
          type: ActionTypes.DELETE_PRE_UPLOAD,
          payload: newState,
        });
      } finally {
      }
    },
    [dispatch, preUploadImages],
  );

  const commitUpload = useCallback(async () => {
    dispatch(loading());
    try {
      const response = await HttpRequest.put('/image/upload');
      // TODO
      console.log(response);
      forwordToDirectory();
    } finally {
      dispatch(unloading());
    }
  }, [dispatch, preUploadImages]);

  const forwordToDirectory = (path: string = '') =>
    dispatch(push(`/images/${path}`));

  return {
    //state
    images,
    preUploadImages,
    //action
    initialize,
    getPreImagesInfo,
    getImages,
    addPreUploadImages,
    deletePreImage,
    commitUpload,
    forwordToDirectory,
  };
};

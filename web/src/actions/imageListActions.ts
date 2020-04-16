import { useCallback } from 'react';
import { Action } from 'redux';
import { useDispatch, useSelector } from 'react-redux';
import { cloneDeep } from 'lodash';

//import { push } from 'connected-react-router';

import HttpRequest from '../service/api/HttpRequest';
// Actions
export const ActionTypes = {
  GET_LIST: 'IMAGE_LIST_GET_LIST',
  GET_UPLOAD: 'IMAGE_LIST_GET_UPLOAD_IMAGES',
  UPDATE_PRE_UPLOAD: 'IMAGE_LIST_UPDATE_PRE_UPLOAD',
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
  const imageList = useSelector(({ imageList }) => imageList);

  const { images, preUploadImages } = imageList;

  //  画像一覧
  const getImages = useCallback(async () => {
    dispatch({
      type: ActionTypes.GET_LIST,
    });
    try {
      const data = await HttpRequest.get(`/image/list`);
      dispatch({
        type: ActionTypes.CHANGE_STATE,
        payload: { images: data },
      });
    } finally {
      // TODO
    }
  }, [dispatch]);

  //  preupload image
  const addPreUploadImages = useCallback(
    async (files: File[]) => {
      const submitData = new FormData();
      files.forEach(element => {
        submitData.append('images', element);
      });

      try {
        const response = await HttpRequest.postImg(
          '/image/pre_upload',
          submitData,
        );
        dispatch({
          type: ActionTypes.UPDATE_PRE_UPLOAD,
          payload: response.data,
        });
      } finally {
        // TODO
      }
    },
    [dispatch, preUploadImages],
  );

  //  PreUpload delete image
  const deletePreImage = useCallback(
    async (id: number) => {
      const copyImages = cloneDeep(preUploadImages);
      const newState = copyImages.filter(elm => elm.info.id !== id);
      dispatch({
        type: ActionTypes.DELETE_PRE_UPLOAD,
        payload: newState,
      });
    },
    [dispatch, preUploadImages],
  );

  return {
    //state
    images,
    preUploadImages,
    //action
    getImages,
    addPreUploadImages,
    deletePreImage,
  };
};

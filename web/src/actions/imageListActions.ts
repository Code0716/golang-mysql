import { useCallback } from 'react';
import { Action } from 'redux';
import { useDispatch, useSelector } from 'react-redux';
//import { push } from 'connected-react-router';

import HttpRequest from '../service/api/HttpRequest';
// Actions
export const ActionTypes = {
  GET_LIST: 'IMAGE_LIST_GET_LIST',
  GET_UPLOAD: 'IMAGE_LIST_GET_UPLOAD_IMAGES',
  POST_PRE_UPLOAD: 'IMAGE_LIST_POST_PRE_UPLOAD',
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

  //  画像upload
  const addUploadImages = useCallback(
    async files => {
      const submitData = new FormData();
      files.forEach(element => {
        submitData.append('images', element);
      });

      try {
        const response = await HttpRequest.postImg('/image/upload', submitData);
        dispatch({
          type: ActionTypes.POST_PRE_UPLOAD,
          payload: response.data,
        });
      } finally {
        // TODO
      }
    },
    [dispatch],
  );

  return {
    //state
    images,
    preUploadImages,
    //action
    getImages,
    addUploadImages,
  };
};

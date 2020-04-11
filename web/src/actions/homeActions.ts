import { useCallback } from 'react';
import { Action } from 'redux';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';

import HttpRequest from '../service/api/HttpRequest';
// Actions
export const ActionTypes = {
  GET: 'HOME_GET',
  POST: 'HOME_POST',
  CHANGE_STATE: 'HOME_CHANGE_STATE',
} as const;

// Action Creators

// Actionの型 Actionを継承 TODO
interface GetAction extends Action {
  type: typeof ActionTypes.GET;
}

interface PostAction extends Action {
  type: typeof ActionTypes.POST;
}

export type HttpRequestActionTypes = GetAction | PostAction;

export const homeActions = () => {
  const dispatch = useDispatch();
  //store
  const homeState = useSelector(({ home }) => home);

  const { getData, postData, entity } = homeState;

  const get = useCallback(
    async url => {
      dispatch({ type: ActionTypes.GET, url });
      const data = await HttpRequest.get(url);
      // Asiaのみ取得
      dispatch({ type: ActionTypes.CHANGE_STATE, payload: { getData: data } });
    },
    [dispatch],
  );

  const post = useCallback(
    // TODO
    (url, json) => dispatch({ type: ActionTypes.POST, payload: { url, json } }),
    [dispatch],
  );

  return {
    //state
    getData,
    postData,
    entity,
    //action
    get,
    post,
  };
};

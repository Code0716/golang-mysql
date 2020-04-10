import { push } from 'connected-react-router';
import { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import HttpRequest from '../service/api/HttpRequest';
// Actions

export const CHANGE_STATE = 'CHANGE_HOME_STATE';
export const CHANGE_ENTITY = 'CHANGE_HOME_ENTITY';
export const HOME_GET = 'HOME_GET';
export const HOME_POST = 'HOME_POST';

// Action Creators

export const homeActions = () => {
  const dispatch = useDispatch();
  //store
  const homeState = useSelector(({ home }) => home);

  const { getData, postData, entity } = homeState;

  const get = useCallback(
    async url => {
      dispatch({ type: HOME_GET, url });
      const data = await HttpRequest.get(url);
      // Asiaのみ取得
      dispatch({ type: CHANGE_STATE, payload: { getData: data } });
    },
    [dispatch],
  );

  const post = useCallback(
    // TODO
    (url, json) => dispatch({ type: HOME_POST, payload: { url, json } }),
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

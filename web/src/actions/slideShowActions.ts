import { useCallback } from 'react';
import { Action } from 'redux';
import { useDispatch, useSelector } from 'react-redux';
//import { push } from 'connected-react-router';

import HttpRequest from '../service/api/HttpRequest';
// Actions
export const ActionTypes = {
  GET_CONTINENTS: 'SLIDE_SHOW_GET_CONTINENTS',
  GET_COUNTRIES: 'SLIDE_SHOW_GET_COUNTRIES',
  GET_CITIES: 'SLIDE_SHOW_GET_CITIES',
  CHANGE_STATE: 'SLIDE_SHOW_CHANGE_STATE',
} as const;

// Action Creators

// Actionの型 Actionを継承 TODO
interface GetCountryAction extends Action {
  type: typeof ActionTypes.GET_COUNTRIES;
}

interface GetCityAction extends Action {
  type: typeof ActionTypes.GET_CITIES;
}

export type HttpRequestActionTypes = GetCountryAction | GetCityAction;

export const slideShowActions = () => {
  const dispatch = useDispatch();
  //store
  const slideShow = useSelector(({ slideShow }) => slideShow);

  const {
    currentContinent,
    description,
    continents,
    countries,
    cities,
  } = slideShow;

  // continent一覧
  const getContinents = useCallback(async () => {
    dispatch({ type: ActionTypes.GET_CONTINENTS });
    try {
      const data = await HttpRequest.get('/continent');
      // Asiaのみ取得
      dispatch({
        type: ActionTypes.CHANGE_STATE,
        payload: { continents: data },
      });
    } finally {
      //TODO
    }
  }, [dispatch]);

  // country一覧
  const getCountries = useCallback(async () => {
    dispatch({ type: ActionTypes.GET_COUNTRIES });
    try {
      const data = await HttpRequest.get('/country');
      // Asiaのみ取得
    } finally {
      // TODO
    }
  }, [dispatch]);

  // City一覧
  const getCities = useCallback(
    async url => {
      dispatch({ type: ActionTypes.GET_CITIES, url });
      try {
        const data = await HttpRequest.get(url);
        // Asiaのみ取得
        dispatch({ type: ActionTypes.CHANGE_STATE, payload: { cities: data } });
      } finally {
        // TODO
      }
    },
    [dispatch],
  );
  //  description
  const getContinentDesc = useCallback(
    async currentContinent => {
      dispatch({
        type: ActionTypes.CHANGE_STATE,
        payload: { currentContinent: currentContinent },
      });
      try {
        //        const data = await HttpRequest.get(`/scraping/${currentContinent}`);
        // Asiaのみ取得
        /*     dispatch({
          type: ActionTypes.CHANGE_STATE,
          payload: { description: data },
        });*/
      } finally {
        // TODO
      }
    },
    [dispatch, currentContinent],
  );

  return {
    //state
    currentContinent,
    description,
    continents,
    cities,
    countries,
    //action
    getContinents,
    getCountries,
    getCities,
    getContinentDesc,
  };
};

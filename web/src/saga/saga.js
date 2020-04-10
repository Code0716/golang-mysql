import HttpRequest from './api/HttpRequest';
import {
  call,
  put,
  fork,
  takeEvery,
  takeLtest,
  take,
} from 'redux-saga/effects';
import { HOME_GET, HOME_POST, homeActions } from '../actions/homeActions';

function* getRest() {
  while (true) {
    const actions = yield take(HOME_GET);
    //call(api,action)
    //callでactionとして、apiに引数を渡せる
    const { payload, error } = yield call(HttpRequest.get, actions);
    if (payload && !error) {
      yield put(homeActions.changeState({ getData: payload }));
    } else {
      //yield put(failureUser(error));
      console.log(error);
    }
  }
}

function* postRest() {
  while (true) {
    const actions = yield take(HOME_POST);
    //call(api,action)
    //callでactionとして、apiに引数を渡せる
    const { payload, error } = yield call(HttpRequest.post, actions);
    if (payload && !error) {
      yield put(homeActions.changeState({ postData: payload }));
    } else {
      //yield put(failureUser(error));
      console.log(error);
    }
  }
}

export default function* rootSaga() {
  yield fork(getRest);
  yield fork(postRest);
}

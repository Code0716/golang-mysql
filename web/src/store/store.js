import { createStore, combineReducers, applyMiddleware, compose } from "redux";
import reducers from "../reducers";

import { connectRouter, routerMiddleware } from "connected-react-router";
import { createBrowserHistory } from "history";

import createSagaMiddleware from "redux-saga";
import rootSaga from "../saga/saga";

//reduxのloggerの実装
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export const history = createBrowserHistory();
/* Storeの実装 */

export default function configureStore() {
  const sagaMiddleware = createSagaMiddleware();
  const store = createStore(
    combineReducers({ ...reducers, router: connectRouter(history) }),

    composeEnhancers(applyMiddleware(sagaMiddleware, routerMiddleware(history)))
  );
  sagaMiddleware.run(rootSaga);
  return store;
}

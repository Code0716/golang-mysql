import { createStore, combineReducers, applyMiddleware, compose } from "redux";

import { connectRouter, routerMiddleware } from "connected-react-router";
import { createBrowserHistory } from "history";
import reducers from "../reducers";
//import createSagaMiddleware from "redux-saga";
export const history = createBrowserHistory();
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
/* Storeの実装 */

export default function configureStore() {
  // const sagaMiddleware = createSagaMiddleware();
  const store = createStore(
    combineReducers({ ...reducers, router: connectRouter(history) }),

    composeEnhancers(
      applyMiddleware(/*sagaMiddleware,*/ routerMiddleware(history))
    )
  );
  //sagaMiddleware.run(rootSaga);
  return store;
}

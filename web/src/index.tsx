'use strict';
import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { Route, Switch } from 'react-router';
import { ConnectedRouter } from 'connected-react-router';
import configureStore, { history } from './store/store';

import { LinksPathnames } from './constant/commonConstant';

import { Home } from '../src/containers/Home';
import { SlideShow } from '../src/containers/SlideShow';
import { ImageList } from './containers/ImageList/ImageList';
import TestGetPost from '../src/containers/TestGetPost';
import NotFound from './containers/NotFound.js';
import SideMenu from '../src/components/SideMenu';
import './sass/style.scss';

const store = configureStore();

ReactDOM.render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <SideMenu />
      <Switch>
        <Route exact path={LinksPathnames.HOME.path} component={Home} />
        <Route exact path={LinksPathnames.SLIDE.path} component={SlideShow} />
        <Route
          exact
          path={LinksPathnames.IMAGE_LIST.path}
          component={ImageList}
        />
        <Route exact path={LinksPathnames.TEST.path} component={TestGetPost} />
        <Route component={NotFound} />
      </Switch>
    </ConnectedRouter>
  </Provider>,
  document.getElementById('root'),
);

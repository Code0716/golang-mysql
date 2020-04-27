import React from 'react';
import { useSelector } from 'react-redux';
import { Portal } from './Portal';
import { LoadingElm } from './LoadingElm';

export const Loading = () => {
  const isLoad = useSelector(state => state.loading.load);
  return (
    isLoad && (
      <Portal domId={'loading'}>
        <LoadingElm className={'full_display'} />
      </Portal>
    )
  );
};

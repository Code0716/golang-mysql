import React from 'react';
import { useSelector } from 'react-redux';
import { Portal } from '../Portal';
import { LoadingElm } from './LoadingElm';
import { RootState } from '../../store/store';

export const Loading = () => {
  const isLoad = useSelector((state: RootState) => state.loading.load);
  return (
    isLoad && (
      <Portal domId={'loading'}>
        <LoadingElm className={'full_display'} />
      </Portal>
    )
  );
};

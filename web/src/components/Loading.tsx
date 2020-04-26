import React from 'react';
import { useSelector } from 'react-redux';
import { Portal } from './Portal';

export const Loading = () => {
  const isLoad = useSelector(state => state.loading.load);
  return (
    isLoad && (
      <Portal domId={'loading'}>
        <div className="outLoading">
          <p className="sentence">NOW LODING </p>
          <span className="loading" />
          <div className="loader" />
        </div>
      </Portal>
    )
  );
};

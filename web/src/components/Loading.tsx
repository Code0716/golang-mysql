import * as React from 'react';

export const Loading = ({ isLoad }) => {
  return (
    <div className="outLoading" style={{ display: isLoad ? 'block' : 'none' }}>
      <p className="sentence">NOW LODING </p>
      <span className="loading" />
      <div className="loader" />
    </div>
  );
};

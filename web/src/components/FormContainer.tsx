import React from 'react';

export const FormContainer = props => {
  return (
    <div className="contents fade-in">
      <div className="form_container">{props.children}</div>
    </div>
  );
};

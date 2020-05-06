import { useCallback } from 'react';

import { useDispatch, useSelector } from 'react-redux';

// action type
export const ActionTypes = {
  SHOW: 'SHOW_SHOT_MESSAGE',
  HIDE: 'SHOT_MESSAGE_HIDE',
  DISAPPEAR: 'SHOT_MESSAGE_DISAPPEAR',
} as const;

// action creator

export const shotMessageAction = () => {
  const show = () => {
    return {
      type: ActionTypes.SHOW,
    };
  };

  const hide = () => {
    return {
      type: ActionTypes.HIDE,
    };
  };

  const disappear = () => {
    return {
      type: ActionTypes.DISAPPEAR,
    };
  };

  return {
    // action
    show,
    hide,
    disappear,
  };
};

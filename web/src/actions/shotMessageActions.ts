import { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store/store';

// action type
export const ActionTypes = {
  SHOW: 'SHOW_SHOT_MESSAGE',
  DISAPPEAR: 'SHOT_MESSAGE_DISAPPEAR',
} as const;

// action creator

export const shotMessageAction = () => {
  const dispatch = useDispatch();

  const { show, message } = useSelector(({ shotMess }: RootState) => shotMess);

  const showMessage = useCallback(
    (message: string) =>
      dispatch({
        type: ActionTypes.SHOW,
        message,
      }),
    [dispatch],
  );

  const disappear = useCallback(
    () => dispatch({ type: ActionTypes.DISAPPEAR }),
    [dispatch],
  );

  return {
    // stre
    show,
    message,
    // action
    showMessage,
    disappear,
  };
};

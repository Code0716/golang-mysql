// action type
export const ActionTypes = {
  LOADING: 'LOADING',
  UNLOADING: 'UNLOADING',
} as const;

// action creator
/**
 * 処理中
 */
export function load() {
  return {
    type: ActionTypes.LOADING,
  };
}

/**
 * 処理後
 */
export function unload() {
  return {
    type: ActionTypes.UNLOADING,
  };
}

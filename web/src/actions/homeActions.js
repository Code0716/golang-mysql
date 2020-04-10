/* Actionsの実装 */

// Action名の定義
export const CHANGE_STATE = "CHANGE_HOME_STATE";
export const CHANGE_ENTITY = "CHANGE_HOME_ENTITY";
export const HOME_GET = "HOME_GET";
export const HOME_POST = "HOME_POST";

// Action Creators

export const homeActions = {
  changeState: payload => ({ type: CHANGE_STATE, payload }),
  changeEntity: payload => ({ type: CHANGE_ENTITY, payload }),
  get: url => ({ type: HOME_GET, url }),
  post: (url, json) => ({ type: HOME_POST, payload: { url, json } })
};

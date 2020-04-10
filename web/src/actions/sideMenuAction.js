/* Actionsの実装 */

// Action名の定義
export const INITALIZE = "SIDE_MENU_INITALIZE";
export const CHANGE_STATE = "CHANGE_SIDE_MENU_STATE";
export const CHANGE_ENTITY = "CHANGE_SIDE_MENU_ENTITY";
export const SIDE_MENU_GET = "SIDE_MENU_GET";
export const SIDE_MENU_POST = "SIDE_MENU_POST";

// Action Creators

export const sideMenuAction = {
  initialize: () => ({ type: INITALIZE }),
  changeState: payload => ({ type: CHANGE_STATE, payload }),
  changeEntity: payload => ({ type: CHANGE_ENTITY, payload })
};

export class LinksPathnames {}

export namespace LinksPathnames {
  export const HOME: { name: string; path: string } = {
    name: 'HOME',
    path: '/',
  };

  export const SLIDE: { name: string; path: string } = {
    name: 'SLIDE SHOW',
    path: '/slide',
  };

  export const TEST: { name: string; path: string } = {
    name: '非同期通信の確認画面',
    path: '/test',
  };

  export const NOT_FOUND: { name: string; path: string } = {
    name: 'NOT FOUND',
    path: 'not_found',
  };
}

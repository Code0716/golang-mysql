import { FC, useEffect } from 'react';
import ReactDOM from 'react-dom';

type PortalProps = { children: any };

export const Portal: FC<PortalProps> = ({ children }) => {
  useEffect(() => {
    return () => document.body.removeChild(target);
  }, []);
  const target = document.createElement('div');
  target.setAttribute('class', 'portal');
  target.setAttribute('id', 'portal');

  const potalDom = document.getElementById('portal');

  if (!potalDom) {
    // BODY直下に新しいdiv要素を作成する
    document.body.appendChild(target);
    // 作成したdivの子要素としてchildrenを描写する
    return ReactDOM.createPortal(children, target);
  }

  return ReactDOM.createPortal(children, potalDom);
};

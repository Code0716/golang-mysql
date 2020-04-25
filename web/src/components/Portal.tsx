import { FC, useEffect } from 'react';
import ReactDOM from 'react-dom';

type PortalProps = { isApply: boolean };

export const Portal: FC<PortalProps> = ({ isApply, children }) => {
  useEffect(() => {
    return () => document.body.removeChild(target);
  }, []);

  const target = document.createElement('div');
  target.setAttribute('class', 'portal');
  target.setAttribute('id', 'portal');

  if (isApply) {
    // BODY直下に新しいdiv要素を作成する
    document.body.appendChild(target);
  }
  // 作成したdivの子要素としてchildrenを描写する
  return ReactDOM.createPortal(children, target);
};

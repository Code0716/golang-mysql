import * as React from 'react';
import { useEffect } from 'react';
import { imageListActions, RouteParams } from '../../actions/imageListActions';
import './style.scss';

export const ImageDetaile: React.FC = () => {
  const { imgBase64, getImage } = imageListActions();

  useEffect(() => {
    getImage();
  }, []);

  return (
    <div className="d-flex">
      <div className="image_box">
        {imgBase64 && <img src={`data:image/png;base64,${imgBase64}`} />}
      </div>
    </div>
  );
};

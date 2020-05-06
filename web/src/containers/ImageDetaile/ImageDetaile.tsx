import * as React from 'react';
import { useEffect, useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { imageListActions, RouteParams } from '../../actions/imageListActions';
import { LoadImage } from '../../reducers/imageListReducer';
import { LoadingElm } from '../../components/Loading/LoadingElm';
import './style.scss';

export const ImageDetaile: React.FC = () => {
  const params: RouteParams = useParams();

  const { getImage, images, forwordToDetaile } = imageListActions();

  useEffect(() => {
    getImage();
  }, [params.id]);

  // 現在の画像を返す
  const currnetImage = useMemo<string | undefined>(() => {
    const data = images.find((item: LoadImage) => {
      if (item.image && item.image.id === Number(params.id)) return item;
    });
    if (!data) return;
    return data.image.img;
  }, [images]);

  return (
    <div className="flex-box-center">
      <button
        className="img_left_arrow"
        onClick={() => forwordToDetaile(Number(params.id) - 1)}
      >
        {'←'}
      </button>
      <div className="image_box">
        {currnetImage ? (
          <div className="inner_box">
            <img src={`data:image/png;base64,${currnetImage}`} />
          </div>
        ) : (
          <LoadingElm className="loading_height" />
        )}
      </div>
      <button
        className="img_right_arrow"
        onClick={() => forwordToDetaile(Number(params.id) + 1)}
      >
        {'→'}
      </button>
    </div>
  );
};

import * as React from 'react';
import { useEffect, useMemo, useCallback } from 'react';
import { useParams } from 'react-router-dom';
import { imageListActions, RouteParams } from '../../actions/imageListActions';
import { LoadImage } from '../../reducers/imageListReducer';
import { LoadingElm } from '../../components/Loading/LoadingElm';
import './style.scss';

export const ImageDetaile: React.FC = () => {
  const params: RouteParams = useParams();
  const paramsID = Number(params.id);
  const { getImage, images, forwordToDetaile } = imageListActions();

  useEffect(() => {
    getImage();
  }, [params.id]);

  // 現在の画像を返す
  const currnetImage = useMemo<string | undefined>(() => {
    const data = images.find((item: LoadImage) => {
      if (item.image && item.image.id === paramsID) return item;
    });
    if (!data) return;
    return data.image.img;
  }, [images, paramsID]);

  const searchPrevNext = (num: number) =>
    images.some((item: LoadImage, index: number) => {
      // スマートじゃないよねー。
      if (paramsID === item.info.ID) {
        if (images[index + num]) return true;
      }
    });

  const isPrevExist = useMemo<boolean>(() => {
    const result = searchPrevNext(-1);
    return result;
  }, [images, paramsID]);

  const isNextExist = useMemo<boolean>(() => {
    const result = searchPrevNext(1);
    return result;
  }, [images, paramsID]);

  return (
    <div className="flex-box-center">
      {isPrevExist && (
        <button
          className="img_left_arrow"
          onClick={() => forwordToDetaile(paramsID - 1)}
        >
          {'←'}
        </button>
      )}
      <div className="image_box">
        {currnetImage ? (
          <div className="inner_box">
            <img src={`data:image/png;base64,${currnetImage}`} />
          </div>
        ) : (
          <LoadingElm className="loading_height" />
        )}
      </div>
      {isNextExist && (
        <button
          className="img_right_arrow"
          onClick={() => forwordToDetaile(paramsID + 1)}
        >
          {'→'}
        </button>
      )}
    </div>
  );
};

import * as React from 'react';
import { useState, useEffect, useMemo, useCallback, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { imageListActions } from '../../actions/imageListActions';

//const createObjectURL = (window.URL || window.webkitURL).createObjectURL;

export const UploadImage = () => {
  const {
    preUploadImages,
    addPreUploadImages,
    deletePreImage,
  } = imageListActions();

  const handleChangeFile = useCallback(
    (e: any) => {
      const target: HTMLInputElement = e.target as HTMLInputElement;
      let files: File[] = [];

      for (let index = 0; index < target.files.length; index++) {
        files.push(target.files.item(index));
      }

      addPreUploadImages(files);
    },
    [preUploadImages],
  );

  const preUploadImagesRender = useMemo(() => {
    const imgs = preUploadImages.map((imagePath: string, index: number) => (
      <div key={`pre-up-img${index}`} className="preupload_img">
        <button
          className="button delete_preup"
          onClick={() => {
            deletePreImage(index);
          }}
        >
          削除
        </button>
        <img src={'data:image/png;base64,' + imagePath} />
      </div>
    ));
    return <div className="preupload_imgbox">{imgs}</div>;
  }, [preUploadImages]);

  return (
    <React.Fragment>
      <label className="image_up_label">
        ＋写真を選択
        <input
          className="d-none"
          type="file"
          accept="image/*"
          multiple
          onChange={e => handleChangeFile(e)}
        />
      </label>
      <div>{preUploadImagesRender}</div>
      {preUploadImages.length !== 0 && (
        <button
          className="action_button"
          onClick={() => {
            // TODO
          }}
        >
          <span>Commit</span>
        </button>
      )}
    </React.Fragment>
  );
};

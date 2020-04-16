import * as React from 'react';
import { useState, useEffect, useMemo, useCallback, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { imageListActions } from '../../actions/imageListActions';
import { LoadImage } from '../../reducers/imageListReducer';

//const createObjectURL = (window.URL || window.webkitURL).createObjectURL;

export const UploadImage = () => {
  useEffect(() => {}, []);

  const {
    preUploadImages,
    addPreUploadImages,
    deletePreImage,
  } = imageListActions();

  const handleChangeFile = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
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
    const imgs = preUploadImages.map((elemnt: LoadImage, index: number) => (
      <React.Fragment key={`pre-up-img${index}`}>
        <div className="preupload_img">
          <button
            className="button delete_preup"
            onClick={() => {
              deletePreImage(elemnt.info.id);
            }}
          >
            削除
          </button>
          <img src={'data:image/png;base64,' + elemnt.img} />
        </div>
        <div className="preupload_description">
          <span className="d-block">Image : {elemnt.info.title}</span>
          <span className="d-block">ID : {elemnt.info.id}</span>
          <span className="d-block">Date : {elemnt.info.create}</span>
          <span className="d-block">Path : {elemnt.info.path}</span>
        </div>
      </React.Fragment>
    ));
    return <div className="preupload_imgbox flex_box">{imgs}</div>;
  }, [preUploadImages]);

  return (
    <React.Fragment>
      <label className="image_up_label">
        ＋写真を選択
        <input
          className="d-none"
          type="file"
          accept=".png,.jpg,.gif"
          multiple
          onChange={handleChangeFile}
        />
      </label>
      {preUploadImagesRender}
      {preUploadImages.length !== 0 && (
        <button
          className="action_button"
          onClick={() => {
            // TODO
          }}
        >
          <span>Commit Upload</span>
        </button>
      )}
    </React.Fragment>
  );
};

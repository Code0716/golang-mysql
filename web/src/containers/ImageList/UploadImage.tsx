import * as React from 'react';
import { useState, useEffect, useMemo, useCallback } from 'react';
import { Portal } from '../../components/Portal';
import { FormContainer } from '../../components/FormContainer';
import { imageListActions } from '../../actions/imageListActions';
import { LoadImage } from '../../reducers/imageListReducer';
import { LoadingElm } from '../../components/Loading/LoadingElm';
import { FormatUtil } from '../../util/FormatUtil';
import './style.scss';

export const UploadImage = () => {
  const {
    // state
    preUploadImages,
    // action
    initialize,
    addPreUploadImages,
    deletePreImage,
    getPreImagesInfo,
  } = imageListActions();

  useEffect(() => {
    getPreImagesInfo();
    return initialize;
  }, []);

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
      <div className="mb-10 flex-box" key={`pre-up-img${index}`}>
        <div className="preupload-img">
          <button
            className="button"
            onClick={() => {
              deletePreImage(elemnt.info.ID);
            }}
          >
            削除
          </button>
          {elemnt.info.Path ? (
            <img className="fade-in" src={elemnt.info.Path} />
          ) : (
            <LoadingElm className={'image_size_loading'} />
          )}
        </div>
        <div className="preupload-description">
          <span className="d-block">Title : {elemnt.info.Title}</span>
          <span className="d-block">ID : {elemnt.info.ID}</span>
          <span className="d-block">
            Date : {FormatUtil.formatedDate(elemnt.info.CreatedAt)}
          </span>
          <span className="d-block">Path : {elemnt.info.Path}</span>
        </div>
      </div>
    ));
    return <div className="preupload-imgbox mb-10">{imgs}</div>;
  }, [preUploadImages]);

  return (
    <React.Fragment>
      <label className="image_up_label mb-5">
        ＋写真を選択
        <input
          className="d-none"
          type="file"
          accept=".png,.jpg,.gif"
          multiple
          onChange={handleChangeFile}
        />
      </label>
      {preUploadImages.length !== 0 && (
        <Portal domId={'preupload'}>
          <FormContainer>
            {preUploadImagesRender}
            <div className="image_commit_box">
              <label className="image_up_label mb-5">
                ＋写真を選択
                <input
                  className="d-none"
                  type="file"
                  accept=".png,.jpg,.gif"
                  multiple
                  onChange={handleChangeFile}
                />
              </label>

              <button
                className="action_button"
                onClick={() => {
                  // TODO
                }}
              >
                <span>Commit Upload</span>
              </button>
            </div>
          </FormContainer>
        </Portal>
      )}
    </React.Fragment>
  );
};

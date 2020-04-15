import * as React from 'react';
import { useState, useEffect, useMemo, useCallback, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { FormContainer } from '../../components/FormContainer';
import { imageListActions } from '../../actions/imageListActions';

//const createObjectURL = (window.URL || window.webkitURL).createObjectURL;

export const UploadImage = () => {
  const { preUploadImages, addUploadImages } = imageListActions();

  const handleChangeFile = useCallback((e: any) => {
    const target: HTMLInputElement = e.target as HTMLInputElement;
    let files: File[] = [];

    for (let index = 0; index < target.files.length; index++) {
      files.push(target.files.item(index));
    }

    addUploadImages(files);
  }, []);

  const preUploadImagesRender = useMemo(() => {
    const imgs = preUploadImages.map((imagePath, index) => (
      <img
        key={`pre-up-img${index}`}
        src={'data:image/png;base64,' + imagePath}
      />
    ));
    return <div>{imgs}</div>;
  }, [preUploadImages]);

  return (
    <FormContainer>
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
    </FormContainer>
  );
};

import * as React from 'react';
import { useState, useEffect, useCallback, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { FormContainer } from '../../components/FormContainer';
import { imageListActions } from '../../actions/imageListActions';

//const createObjectURL = (window.URL || window.webkitURL).createObjectURL;

export const UploadImage = () => {
  const { addUploadImages } = imageListActions();

  const handleChangeFile = (e: any) => {
    const target: HTMLInputElement = e.target as HTMLInputElement;
    let files: File[] = [];

    for (let index = 0; index < target.files.length; index++) {
      files.push(target.files.item(index));
    }

    addUploadImages(files);
  };

  return <input type="file" multiple onChange={e => handleChangeFile(e)} />;
};

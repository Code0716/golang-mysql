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
    const file: File = target.files.item(0);
    addUploadImages(file);
  };

  return <input type="file" onChange={e => handleChangeFile(e)} />;
};

import * as React from 'react';
import { useParams } from 'react-router-dom';
import { FormContainer } from '../../components/FormContainer';
import { imageListActions } from '../../actions/imageListActions';
import { UploadImage } from './UploadImage';
import { ListImages } from './ListImages';

import './style.scss';

type RouteParams = {
  directory?: string;
};
export const ImageList = () => {
  const params = useParams();

  const {
    //state
    //action
    forwordToDirectory,
  } = imageListActions();

  return (
    <FormContainer>
      <div className="flex-box mb-10 p-10">
        <button className="button" onClick={() => forwordToDirectory('upload')}>
          Upload images
        </button>
        <button className="button" onClick={() => forwordToDirectory('list')}>
          Image list
        </button>
      </div>
      {params.directory === 'upload' && <UploadImage />}
      {params.directory === 'list' && <ListImages />}
    </FormContainer>
  );
};

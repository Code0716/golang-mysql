import * as React from 'react';
import { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { push } from 'connected-react-router';
import { AutoSizer, Column, Table } from 'react-virtualized';
import { FormContainer } from '../../components/FormContainer';
import { imageListActions } from '../../actions/imageListActions';
import { UploadImage } from './UploadImage';

export const ImageList = () => {
  const dispatch = useDispatch();
  const [selectContainer, setSelectContainer] = useState('');

  const [isHover, setIsHover] = useState(false);
  const {
    //state
    images,
    //action
    getImages,
  } = imageListActions();

  const makePath = continent => continent.replace(' ', '_').toLowerCase();

  useEffect(() => {}, []);

  return (
    <FormContainer>
      <div className="flex_box">
        <button className="button" onClick={() => setSelectContainer('upload')}>
          Upload images
        </button>
        <button className="button" onClick={() => setSelectContainer('list')}>
          Image list
        </button>
      </div>
      {selectContainer === 'upload' && <UploadImage />}
      {selectContainer === 'list' && (
        <AutoSizer>
          {({ width, height }) => (
            <div className="d-flex">
              <Table
                data={images}
                height={height}
                width={width}
                headerHeight={35}
                rowHeight={35}
                rowGetter={({ index }) => images[index]}
                rowCount={images.length}
                rowClassName="virtualized_row"
                onRowClick={({ rowData }) =>
                  dispatch(push(makePath(rowData.continent)))
                }
              >
                <Column
                  width={200}
                  label="Images"
                  dataKey="_"
                  className="colum_row"
                  cellRenderer={({ rowData }) => {}}
                />
              </Table>
            </div>
          )}
        </AutoSizer>
      )}
    </FormContainer>
  );
};

import * as React from 'react';
import { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { push } from 'connected-react-router';
import { AutoSizer, Column, Table } from 'react-virtualized';
import { FormContainer } from '../components/FormContainer';
import { slideShowActions } from '../actions/slideShowActions';

export const SlideShow = () => {
  const dispatch = useDispatch();

  const [isHover, setIsHover] = useState(false);
  const {
    // store
    currentContinent,
    description,
    continents,
    countries,
    cities,
    // actions
    getContinents,
    getCities,
    getCountries,
    getContinentDesc,
  } = slideShowActions();

  const makePath = continent => continent.replace(' ', '_').toLowerCase();

  useEffect(() => {
    getContinents();
  }, []);

  return (
    <FormContainer>
      <AutoSizer>
        {({ width, height }) => (
          <div className="d-flex">
            <Table
              data={continents}
              height={height}
              width={200}
              headerHeight={35}
              rowHeight={35}
              rowGetter={({ index }) => continents[index]}
              rowCount={continents.length}
              rowClassName="virtualized_row"
              onRowClick={({ rowData }) =>
                dispatch(push(makePath(rowData.continent)))
              }
            >
              <Column
                width={200}
                label="Continent"
                dataKey="continent"
                className="colum_row"
                cellRenderer={({ rowData }) => (
                  <span
                    className="d-block w100"
                    onMouseOver={() => {
                      setIsHover(true);
                      getContinentDesc(rowData.continent);
                    }}
                    onMouseLeave={() => setIsHover(false)}
                  >
                    {rowData.continent}
                  </span>
                )}
              />
            </Table>
            {isHover && (
              <div className="continent_description">
                <span className="d-block">{currentContinent}</span>
                <p>{description}</p>
              </div>
            )}
          </div>
        )}
      </AutoSizer>
    </FormContainer>
  );
};

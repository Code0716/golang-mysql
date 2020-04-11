import * as React from 'react';
import { useEffect } from 'react';
import { useSelector } from 'react-redux';
import { AutoSizer, Column, Table } from 'react-virtualized';
import { FormContainer } from '../components/FormContainer';
import { slideShowActions } from '../actions/slideShowActions';

export const SlideShow = () => {
  const slideShow = useSelector(({ slideShow }) => slideShow);

  const { continents, countries, cities } = slideShow;
  const { getContinents, getCities, getCountries } = slideShowActions();

  useEffect(() => {
    getContinents();
  }, []);

  return (
    <FormContainer>
      <AutoSizer>
        {({ width, height }) => (
          <React.Fragment>
            <Table
              data={continents}
              height={height}
              width={width}
              rowHeight={30}
              rowGetter={({ index }) => continents[index]}
              rowCount={continents.length}
              rowClassName="virtualized_row"
            >
              <Column width={200} label="continent" dataKey="continent" />
            </Table>
          </React.Fragment>
        )}
      </AutoSizer>
    </FormContainer>
  );
};

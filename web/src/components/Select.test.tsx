import * as React from 'react';
import * as Adapter from 'enzyme-adapter-react-16';
import * as Enzyme from 'enzyme';
import { shallow } from 'enzyme';
import { Select } from './Select';

Enzyme.configure({ adapter: new Adapter() });

describe('Select.tsx test', () => {
  const fn = jest.fn();
  const value = '/test';
  const option = [{ value: 'test' }, { value: 'test2' }, { value: 'test3' }];
  const wrapper = shallow(
    <Select value={value} option={option} onChange={fn} />,
  );

  test('Select value', () => {
    expect(wrapper.find('select').props().value).toBe(value);
    wrapper.setProps({ value: 'test2' });
    expect(wrapper.props().value).toBe('test2');
  });

  test('Select onChange', () => {
    expect(wrapper.find('select').props().onChange).toEqual(fn);
    wrapper.find('select').simulate('change');
    expect(fn).toHaveBeenCalled();
  });
});

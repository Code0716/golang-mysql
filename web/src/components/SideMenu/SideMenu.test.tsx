import * as React from 'react';
import * as Adapter from 'enzyme-adapter-react-16';
import * as Enzyme from 'enzyme';
import { shallow } from 'enzyme';
import { SideMenu } from './SideMenu';

Enzyme.configure({ adapter: new Adapter() });

describe('SideMenu.tsx test', () => {
  const wrapper = shallow(<SideMenu />);

  test('SideMenu onChange', () => {
    expect(wrapper.find('span').first().hasClass('is-active')).toBeFalsy();
    expect(wrapper.find('input').simulate('change'));
    expect(wrapper.find('span').first().hasClass('is-active')).toBeTruthy();
  });
});

import React from 'react'
import Adapter from 'enzyme-adapter-react-16'
import Enzyme, { shallow } from 'enzyme'
import { Hello } from './Hello'

Enzyme.configure({ adapter: new Adapter() })

describe('Hello.tsx', () => {
    const makeComponent = () => <Hello name="sugai" />
    const wrapper = shallow(makeComponent())
    test('show Hello,', () => {
        expect(wrapper.find('div').text()).toBe('HELLO,sugai')
    })
})

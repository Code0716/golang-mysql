import React from 'react';
import { Link } from 'react-router-dom';
import { FormContainer } from '../components/FormContainer';

class NotFound extends React.Component {
  render() {
    return (
      <FormContainer>
        <h1>ページが見つかりません。</h1>
        <Link to="/">Top</Link>
      </FormContainer>
    );
  }
}
export default NotFound;

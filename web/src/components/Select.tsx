import React from 'react';

export const Select = ({ value, option, onChange }) => {
  const options = option.map(opt => (
    <option key={opt.value} value={opt.value}>
      {opt.value}
    </option>
  ));

  return (
    <select value={value} onChange={onChange}>
      {options}
    </select>
  );
};

import * as React from 'react';
import { useState, useEffect } from 'react';
import { Portal } from '../Portal';
import './style.scss';

type Props = {
  message: string;
};

export const ShotMessage: React.FC = () => {
  return (
    <Portal
      domId="shot_message"
      className={'shot_message_box fade-in show_mess'}
    >
      <div className="massage_box">{`Message : shot message`}</div>
    </Portal>
  );
};

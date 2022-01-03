import React from 'react'

export const MessageList = ({ children }: { children: React.ReactNode }) => (
  <ul id='messages-list' className='list-unstyled'>
    {children}
  </ul>
)

import React from 'react'

export const ClapButton = () => (
  <form className='ms-auto' data-test='message-clap-form'>
    <button className='btn'>
      👏🏻 <span data-test='clap-count'>Claps</span>
    </button>
  </form>
)

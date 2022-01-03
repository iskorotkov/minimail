import React from 'react'
import { Link } from 'react-router-dom'
import { Message } from '../models/message'

export const MessageCard = ({
  message,
  showLink = true,
  buttonEnabled = true,
  onClap = () => {}
}: {
  message: Message
  showLink?: boolean
  buttonEnabled?: boolean
  onClap?: () => void
}) => (
  <article className='card' data-test='message'>
    <div className='card-body'>
      <header className='card-title d-flex'>
        <div className='text-muted' data-test='message-author'>
          {message.author}
        </div>
        {showLink && (
          <Link to={`/messages/${message.id}`} className='card-link ms-auto' data-test='message-open'>
            Открыть ↗️
          </Link>
        )}
      </header>

      <div className='card-text' data-test='message-text'>
        {message.message}
      </div>

      <div className='d-flex'>
        <form className='ms-auto' data-test='message-clap-form'>
          <button
            className='btn'
            disabled={!buttonEnabled}
            onClick={e => {
              e.preventDefault()
              onClap()
            }}
          >
            👏🏻 <span data-test='clap-count'>{message.claps}</span>
          </button>
        </form>
      </div>
    </div>
  </article>
)

import { useState } from 'react'
import { addMessage } from '../services/messages'

type Alert = {
  message: string
  className: string
}

const noAlert = () => ({ message: '', className: 'alert alert d-none' })
const sendingAlert = () => ({ message: 'Загрузка...', className: 'alert alert-warning' })
const sentAlert = () => ({ message: 'Сообщение отправлено', className: 'alert alert-success' })
const errorAlert = (error: string) => ({ message: error, className: 'alert alert-danger' })

export const MessageForm = ({ onAdded }: { onAdded?: () => void }) => {
  const [author, setAuthor] = useState('')
  const [message, setMessage] = useState('')
  const [buttonEnabled, setButtonEnabled] = useState(true)
  const [alert, setAlert] = useState<Alert>(noAlert())

  const onSendClicked = () => {
    setButtonEnabled(false)
    setAlert(sendingAlert())

    addMessage(author, message)
      .then(data => {
        if ('id' in data) {
          setAlert(sentAlert())
          onAdded?.call(undefined)
        } else {
          setAlert(errorAlert(data.message.replace(';', ';\n')))
        }
      })
      .finally(() => setButtonEnabled(true))
  }

  return (
    <form className='card border-secondary mb-3' data-test='send-form'>
      <fieldset>
        <legend className='card-header h5 border-secondary bg-dark text-light'>📩 Отправить письмо</legend>
        <div className='card-body'>
          <div className='mb-3' data-test='send-alert'>
            <div className={alert.className}>{alert.message}</div>
          </div>

          <div className='mb-3'>
            <label htmlFor='sender' className='form-label'>
              От кого:
            </label>
            <input
              id='sender'
              name='sender'
              className='form-control'
              type='text'
              placeholder='Имя отправителя'
              value={author}
              onChange={e => setAuthor(e.target.value)}
            />
          </div>

          <div className='mb-3'>
            <label htmlFor='message' className='form-label'>
              Сообщение:
            </label>
            <textarea
              id='message'
              name='message'
              className='form-control'
              placeholder='Текст сообщения'
              value={message}
              onChange={e => setMessage(e.target.value)}
            ></textarea>
          </div>

          <div className='d-flex'>
            <button
              className='btn btn-outline-success ms-auto'
              disabled={!buttonEnabled}
              onClick={e => {
                e.preventDefault()
                onSendClicked()
              }}
            >
              ✏️ Отправить
            </button>
          </div>
        </div>
      </fieldset>
    </form>
  )
}

import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { Footer } from '../components/Footer'
import { Header } from '../components/Header'
import { MessageCard } from '../components/MessageCard'
import { Message } from '../models/message'
import { addClap, getMessage } from '../services/messages'

export const MessagePage = () => {
  const { messageId } = useParams()
  const [message, setMessage] = useState<Message | null>(null)
  const [buttonEnabled, setButtonEnabled] = useState(true)

  useEffect(() => {
    getMessage(parseInt(messageId!)).then(data => setMessage(data))
  }, [messageId])

  const onClap = () => {
    setButtonEnabled(false)
    setMessage({
      ...message!,
      claps: message!.claps + 1
    })

    addClap(message!.id)
      .then(data =>
        setMessage({
          ...message!,
          claps: data.count
        })
      )
      .finally(() => setButtonEnabled(true))
  }

  return (
    <div className='d-flex flex-column bg-light min-vh-100'>
      <Header />

      <main className='container flex-fill my-3'>
        {message ? (
          <MessageCard message={message} showLink={false} buttonEnabled={buttonEnabled} onClap={onClap} />
        ) : (
          <p>Loading...</p>
        )}
      </main>

      <Footer />
    </div>
  )
}

import React, { useEffect, useState } from 'react'
import { Footer } from '../components/Footer'
import { Header } from '../components/Header'
import { MessageCard } from '../components/MessageCard'
import { MessageForm } from '../components/MessageForm'
import { MessageList } from '../components/MessageList'
import { MessageListItem } from '../components/MessageListItem'
import { Message } from '../models/message'
import { addClap, getAllMessages } from '../services/messages'

export const IndexPage = () => {
  const [messages, setMessages] = useState<Message[] | null>(null)
  const [disabledButtons, setDisabledButtons] = useState([] as number[])

  const updateList = () => getAllMessages().then(data => setMessages(data))

  useEffect(() => {
    updateList()
  }, [])

  const onClap = (message: Message) => {
    setDisabledButtons(_ => [..._, message.id])
    setMessages(
      messages!.map(_ =>
        _ !== message
          ? _
          : {
              ..._,
              claps: _.claps + 1
            }
      )
    )

    addClap(message.id)
      .then(() => updateList())
      .finally(() => setDisabledButtons(_ => _.filter(_ => _ !== message.id)))
  }

  return (
    <div className='d-flex flex-column bg-light min-vh-100'>
      <Header />

      <main className='container flex-fill my-3'>
        <MessageForm
          onAdded={() => {
            // TODO: Do something.
          }}
        />

        <MessageList>
          {messages?.map(_ => (
            <MessageListItem key={_.id}>
              <MessageCard message={_} onClap={() => onClap(_)} buttonEnabled={disabledButtons.indexOf(_.id) === -1} />
            </MessageListItem>
          ))}
        </MessageList>
      </main>

      <Footer />
    </div>
  )
}

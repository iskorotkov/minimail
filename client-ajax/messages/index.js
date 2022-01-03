import { addClap, fetchMessage } from '../api.js'

const findElements = () => {
  /** @type {HTMLDivElement | null} */
  const messageAuthor = document.querySelector('[data-test="message-author"]')
  if (!messageAuthor) {
    throw new Error('message author is null')
  }

  /** @type {HTMLDivElement | null} */
  const messageText = document.querySelector('[data-test="message-text"]')
  if (!messageText) {
    throw new Error('message text is null')
  }

  /** @type {HTMLSpanElement | null} */
  const clapCount = document.querySelector('[data-test="clap-count"]')
  if (!clapCount) {
    throw new Error('clap count is null')
  }

  /** @type {HTMLButtonElement | null} */
  const clapButton = document.querySelector('button')
  if (!clapButton) {
    throw new Error('clap button is null')
  }

  return { messageAuthor, messageText, clapCount, clapButton }
}

const extractMessageId = () => {
  const messageId = document.location.pathname.split('/').reverse().slice(0, 1).map(parseInt)[0]
  if (!messageId) {
    throw new Error('error getting message id from path')
  }

  return messageId
}

const messageId = extractMessageId()
const elements = findElements()

const recreateCard = (/** @type {() => void | undefined} */ onCreated) =>
  fetchMessage(messageId)
    .then(async resp => {
      if (resp.ok) {
        /** @type {{id: number, author: string, message: string, claps: number}} */
        const message = await resp.json()

        elements.messageAuthor.innerText = message.author
        elements.messageText.innerText = message.message
        elements.clapCount.innerText = message.claps.toString()

        elements.clapButton.addEventListener('click', e => {
          e.preventDefault()
          e.stopPropagation()

          elements.clapButton.disabled = true

          const currentClapsCount = parseInt(elements.clapCount.innerText)
          elements.clapCount.innerText = (currentClapsCount + 1).toString()

          addClap(message.id)
            .then(() => onCreated?.call(undefined))
            .catch(e => console.error(e))
            .finally(() => (elements.clapButton.disabled = false))
        })

        return
      }

      console.error(`Error occurred: ${resp.statusText}`)
    })
    .catch(e => console.error(e))

recreateCard()

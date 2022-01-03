import { addClap, addMessage, fetchMessages } from './api.js'

const findElements = () => {
  /** @type {HTMLUListElement | null} */
  const messagesList = document.querySelector('#messages-list')
  if (!messagesList) {
    throw new Error('messages list is null')
  }

  /** @type {HTMLTemplateElement | null} */
  const messageTemplate = document.querySelector('#message-template')
  if (!messageTemplate) {
    throw new Error('message template is null')
  }

  /** @type {HTMLFormElement | null} */
  const sendForm = document.querySelector('[data-test="send-form"]')
  if (!sendForm) {
    throw new Error('send form is null')
  }

  /** @type {HTMLDivElement | null} */
  const alertNode = sendForm.querySelector('[data-test="send-alert"]')
  if (!alertNode) {
    throw new Error('alert node is null')
  }

  /** @type {HTMLDivElement | null} */
  const alertTextNode = alertNode.querySelector('.alert')
  if (!alertTextNode) {
    throw new Error('alert text node is null')
  }

  /** @type {HTMLInputElement | null} */
  const authorNode = sendForm.querySelector('[name="sender"]')
  if (!authorNode) {
    throw new Error('author node is null')
  }

  /** @type {HTMLInputElement | null} */
  const messageNode = sendForm.querySelector('[name="message"]')
  if (!messageNode) {
    throw new Error('message node is null')
  }

  /** @type {HTMLButtonElement | null} */
  const formButton = sendForm.querySelector('button')
  if (!formButton) {
    throw new Error('form button is null')
  }

  return {
    messagesList,
    messageTemplate,
    alertNode,
    form: { sendForm, alertNode, alertTextNode, authorNode, messageNode, formButton }
  }
}

const renderMessages = (
  /** @type {HTMLUListElement} */ messagesList,
  /** @type {HTMLTemplateElement} */ messageTemplate,
  /** @type {Array.<{id: number, author: string, message: string, claps: number}>} */ messages,
  /** @type {() => void | undefined} */ onCreated
) => {
  const messageNodes = []
  for (const message of messages) {
    /** @type {HTMLLIElement} */
    // @ts-expect-error
    const messageNode = messageTemplate.content.cloneNode(true)

    /** @type {HTMLDivElement | null} */
    const authorNode = messageNode.querySelector('[data-test="message-author"]')
    if (authorNode) {
      authorNode.innerText = message.author
    }

    /** @type {HTMLDivElement | null} */
    const textNode = messageNode.querySelector('[data-test="message-text"]')
    if (textNode) {
      textNode.innerText = message.message
    }

    /** @type {HTMLSpanElement | null} */
    const clapsNode = messageNode.querySelector('[data-test="clap-count"]')
    if (clapsNode) {
      clapsNode.innerText = message.claps.toString()

      /** @type {HTMLButtonElement | null} */
      const clapButton = messageNode.querySelector('button')
      if (!clapButton) {
        throw new Error('clap button is null')
      }

      clapButton.addEventListener('click', (/** @type {MouseEvent} */ e) => {
        e.preventDefault()
        e.stopPropagation()

        clapButton.disabled = true

        const currentClapsCount = parseInt(clapsNode.innerText)
        clapsNode.innerText = (currentClapsCount + 1).toString()

        addClap(message.id)
          .then(() => onCreated?.call(undefined))
          .catch(e => console.error(e))
          .finally(() => (clapButton.disabled = false))
      })
    }

    /** @type {HTMLLinkElement | null} */
    const openLink = messageNode.querySelector('[data-test="message-open"]')
    if (openLink) {
      openLink.href = `./messages/${message.id}`
    }

    messageNodes.push(messageNode)
  }

  messagesList.replaceChildren(...messageNodes)
}

const setClasses = (/** @type {HTMLElement} */ element, /** @type {Array.<string>} */ classes) =>
  (element.className = classes.join(' '))

const elements = findElements()

const setupCreateForm = (/** @type {(() => void) | undefined} */ onCreated) => {
  elements.form.sendForm.addEventListener('submit', (/** @type {SubmitEvent} */ e) => {
    e.preventDefault()
    e.stopPropagation()

    setClasses(elements.form.alertTextNode, ['alert', 'alert-warning'])
    elements.form.alertTextNode.innerText = texts.loading
    elements.form.formButton.disabled = true

    addMessage(elements.form.authorNode.value, elements.form.messageNode.value)
      .then(async resp => {
        if (resp.ok) {
          await resp.json()
          setClasses(elements.form.alertTextNode, ['alert', 'alert-success'])
          elements.form.alertTextNode.innerText = texts.messageSent
          onCreated?.call(undefined)
          return
        }

        const { message } = await resp.json()
        setClasses(elements.form.alertTextNode, ['alert', 'alert-danger'])
        elements.form.alertTextNode.innerText = message.replace(';', ';\n')
      })
      .catch(() => {
        setClasses(elements.form.alertTextNode, ['alert', 'alert-danger'])
        elements.form.alertTextNode.innerText = 'Ошибка отправки запроса'
      })
      .finally(() => {
        elements.form.formButton.disabled = false
      })
  })
}

const texts = {
  messageSent: 'Сообщение отправлено',
  loading: 'Загрузка...'
}

const recreateList = () =>
  fetchMessages()
    .then(resp => {
      if (resp.ok) {
        return resp.json()
      }

      throw new Error(resp.statusText)
    })
    .then(data => renderMessages(elements.messagesList, elements.messageTemplate, data, recreateList))
    .catch((/** @type {Error} */ e) => console.error(e))

setupCreateForm(recreateList)

recreateList()

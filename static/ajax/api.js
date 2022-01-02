const apiUrl = 'http://localhost:8080/api'

export const fetchMessages = () => fetch(`${apiUrl}/messages`)

export const fetchMessage = (/** @type {number} */ id) => fetch(`${apiUrl}/messages/${id}`)

export const addClap = (/** @type {number} */ id) => fetch(`${apiUrl}/messages/${id}/claps`, { method: 'POST' })

export const addMessage = (/** @type {string} */ author, /** @type {string} */ message) =>
  fetch(`${apiUrl}/messages`, {
    method: 'POST',
    body: JSON.stringify({ author, message }),
    headers: { 'Content-Type': 'application/json' }
  })

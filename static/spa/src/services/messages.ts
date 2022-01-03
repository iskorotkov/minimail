import axios from 'axios'
import { Message } from '../models/message'

const apiUrl = process.env.REACT_APP_API_URL

export const getAllMessages = async () => {
  try {
    const resp = await axios.get(`${apiUrl}/messages`)
    return resp.data as Message[]
  } catch (e) {
    console.error(e)
    throw e
  }
}

export const getMessage = async (messageId: number) => {
  try {
    const resp = await axios.get(`${apiUrl}/messages/${messageId}`)
    return resp.data as Message
  } catch (e) {
    console.error(e)
    throw e
  }
}

export const addMessage = async (author: string, message: string) => {
  try {
    const resp = await axios.post(`${apiUrl}/messages`, { author, message }, { validateStatus: () => true })
    return resp.data as Message | { message: string }
  } catch (e) {
    console.error(e)
    throw e
  }
}

export const addClap = async (messageId: number) => {
  try {
    const resp = await axios.post(`${apiUrl}/messages/${messageId}/claps`)
    return resp.data as { count: number }
  } catch (e) {
    console.error(e)
    throw e
  }
}

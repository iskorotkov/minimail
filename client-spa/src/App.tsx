import React from 'react'
import { Routes, Route } from 'react-router-dom'
import { IndexPage } from './pages/IndexPage'
import { MessagePage } from './pages/MessagePage'

function App() {
  return (
    <Routes>
      <Route path='/' element={<IndexPage />} />
      <Route path='/messages/:messageId' element={<MessagePage />} />
    </Routes>
  )
}

export default App

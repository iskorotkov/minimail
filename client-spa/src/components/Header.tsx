import React from 'react'
import { Link } from 'react-router-dom'

export const Header = () => (
  <nav className='navbar navbar-expand-lg navbar-dark bg-dark'>
    <div className='container'>
      <Link className='navbar-brand mx-auto' to='/'>
        📨 Мини-почта 📨
      </Link>
    </div>
  </nav>
)

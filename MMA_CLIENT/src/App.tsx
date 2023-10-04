// import { useState, useEffect } from 'react'
import './styles/App.scss'
import {BrowserRouter, Routes, Route, Link} from 'react-router-dom'
import Home from './routes/Home'
import Sidebar from './Sidebar'
import OpenChat from './OpenChat'
import Navbar from './Navbar'


const App : React.FC = () => {
  return (
      <main>
            <div className="app">
              <Navbar/>
              <div className="app__body">  
                <Sidebar/>
                <OpenChat/>
              </div>
            </div>
      </main>
  )
}

export default App

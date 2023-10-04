import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import ErrorPage from './routes/ErrorPage.tsx'
import Register from './routes/Register.tsx'
import './styles/index.scss'

import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Home from './routes/Home.tsx'
import Login from './routes/Login.tsx'
import VerifyUser from './routes/VerifyUser.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/app",
    element: <App/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/homepage",
    element: <Home/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/register",
    element: <Register/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/login",
    element: <Login/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/verifyUser",
    element: <VerifyUser/>,
    errorElement: <ErrorPage/>
  },
]);


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
   <RouterProvider router={router}></RouterProvider>
  </React.StrictMode>,
)

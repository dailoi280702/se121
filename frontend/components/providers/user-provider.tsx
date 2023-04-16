'use client'

import { useEffect } from 'react'

interface Props {
  children: React.ReactNode
}

const FetchUser = async () => {
  try {
    const response = await fetch('http://localhost:8000/v1/auth', {
      method: 'GET',
      credentials: 'include',
    })

    if (!response.ok) {
      const data = await response.text()
      return String(data)
    }

    const data = await response.json()
    return String(data)
  } catch (err) {
    return String(err)
  }
}

export default function UserProvider({ children }: Props) {
  // const user = await FetchUser()
  useEffect(() => {
    const fetchUser = async () => {
      const data = await FetchUser()
      window.alert(String(data))
    }

    fetchUser()
  }, [])

  return <div>{children}</div>
}

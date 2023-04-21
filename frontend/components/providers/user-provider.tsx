'use client'

import { atom, useAtom, useSetAtom } from 'jotai'
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
      return null
    }

    const data = await response.json()
    return data
  } catch (err) {
    return null
  }
}

export type User = {
  id: string
  name: string
  email?: string
  createAt: Date
  isAdmin?: boolean
} | null

export const UserAtom = atom<User>(null)

export default function UserProvider({ children }: Props) {
  const setUser = useSetAtom(UserAtom)

  useEffect(() => {
    const fetchUser = async () => {
      const user = (await FetchUser()) as User
      if (user) {
        user.createAt = new Date(user.createAt)
      }
      setUser(user)
    }

    fetchUser()
  }, [setUser])

  return <div>{children}</div>
}

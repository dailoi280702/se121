'use client'

import { useAtomValue } from 'jotai'
import { useEffect, useState } from 'react'
import { UserAtom } from './providers/user-provider'

const AdminOnlyWrapper = ({ children }: { children: React.ReactNode }) => {
  const user = useAtomValue(UserAtom)
  const [isAdmin, setIsAdmin] = useState(false)

  useEffect(() => {
    setIsAdmin(!!(user && user.isAdmin))
  }, [user])

  return <>{isAdmin && children}</>
}

export default AdminOnlyWrapper

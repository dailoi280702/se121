'use client'

import { useAtomValue } from 'jotai'
import { ReactNode, useEffect, useState } from 'react'
import { UserAtom } from './providers/user-provider'

const UserOnlyWarpper = ({
  children,
  fallbackElement,
}: {
  children: React.ReactNode
  fallbackElement: ReactNode
}) => {
  const user = useAtomValue(UserAtom)
  const [isUser, setIsUser] = useState(false)

  useEffect(() => {
    setIsUser(!!user)
  }, [user])

  return <>{isUser ? children : fallbackElement}</>
}

export default UserOnlyWarpper

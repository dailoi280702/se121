'use client'

import { useAtomValue } from 'jotai'
import { useEffect, useState } from 'react'
import { UserAtom } from './providers/user-provider'

const AdminOnlyPage = ({ children }: { children: React.ReactNode }) => {
  const user = useAtomValue(UserAtom)
  const [loaded, setLoaded] = useState(false)

  useEffect(() => {
    setLoaded(true)
  }, [user])

  return (
    <>
      {loaded && user && user.isAdmin ? (
        children
      ) : (
        <div className="m-auto mt-4 text-center">
          This page could not be found.
        </div>
      )}
    </>
  )
}

export default AdminOnlyPage

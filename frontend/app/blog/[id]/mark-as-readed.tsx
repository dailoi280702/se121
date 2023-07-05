'use client'

import { UserAtom } from '@/components/providers/user-provider'
import { useAtomValue } from 'jotai'
import { useEffect } from 'react'

export default function MarkAsReaded({ blogId }: { blogId: Number }) {
  const user = useAtomValue(UserAtom)

  useEffect(() => {
    if (user) {
      const cleanup = async () => {
        try {
          const res = await fetch(`http://localhost:8000/v1/user/readed-blog`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ userId: user.id, blogId: blogId }),
          })
          if (!res.ok) {
            console.log('failed to mark blog as read', await res.text())
          }
        } catch (e) {
          console.log('failed to mark blog as read', e)
        }
      }

      cleanup()
    }
  }, [user, blogId])

  return <></>
}

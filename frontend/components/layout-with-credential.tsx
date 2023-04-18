import { UserAtom } from './providers/user-provider'
import { useAtomValue } from 'jotai'
import { useEffect } from 'react'
import { useRouter } from 'next/navigation'

type credentialOption = 'not required' | 'required' | 'admin'

interface Props {
  children: React.ReactNode
  option?: credentialOption
  callbackUrl?: string
  toBeDisplayed?: boolean
}

export default function LayoutWithCredential({
  children,
  option = 'not required',
  callbackUrl = '/',
  toBeDisplayed = true,
}: Props) {
  const user = useAtomValue(UserAtom)
  const router = useRouter()

  useEffect(() => {
    if (!toBeDisplayed && user && option != 'not required') {
      router.push(callbackUrl)
    }
  }, [user, option, callbackUrl, router, toBeDisplayed])

  return <>{children}</>
}

'use client'
import SignInForm from '@/components/forms/sign-in'
import LayoutWithCredential from '@/components/layout-with-credential'
import Link from 'next/link'
import { usePathname } from 'next/navigation'

export default function Page({ params }: { params: { callbackUrl: string } }) {
  const path = usePathname()

  return (
    <LayoutWithCredential
      option="required"
      toBeDisplayed={false}
      callbackUrl={
        params.callbackUrl == '' ? '/' : decodeURIComponent(params.callbackUrl)
      }
    >
      <div className="text-center">
        <h2 className="text-xl font-bold text-neutral-800 p-14">SIGN IN</h2>
        <SignInForm callbackUrl={decodeURIComponent(params.callbackUrl)} />
        <div className="px-6 py-14 text-sm sm:py-6 sm:border-neutral-200 sm:bg-neutral-100 sm:border-0 sm:border-t">
          don&apos;t have an account?{' '}
          <Link
            className="text-teal-600"
            href={`/auth/signup/${encodeURIComponent(path)}`}
          >
            Sign up
          </Link>{' '}
          for free
        </div>
      </div>
    </LayoutWithCredential>
  )
}

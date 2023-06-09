'use client'
import SignUpForm from '@/components/forms/sign-up'
import LayoutWithCredential from '@/components/layout-with-credential'
import Link from 'next/link'
import { usePathname } from 'next/navigation'

export default function Page({ params }: { params: { callbackUrl: string } }) {
  const path = usePathname()

  return (
    <div className="text-center">
      <h2 className="p-14 text-xl font-bold text-neutral-800">SIGN UP</h2>
      <LayoutWithCredential
        option="required"
        toBeDisplayed={false}
        callbackUrl={
          params.callbackUrl == ''
            ? '/signin/%2F'
            : decodeURIComponent(params.callbackUrl)
        }
      >
        <SignUpForm callbackUrl={decodeURIComponent(params.callbackUrl)} />
      </LayoutWithCredential>
      <div className="px-6 py-14 text-sm sm:border-0 sm:border-t sm:border-neutral-200 sm:bg-neutral-100 sm:py-6">
        already have an account?{' '}
        <Link
          className="text-teal-600"
          href={`/auth/singup/${encodeURIComponent(path)}`}
        >
          Sign in
        </Link>{' '}
        to countinue
      </div>
    </div>
  )
}

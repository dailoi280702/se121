'use client'
import SignInForm from '@/components/forms/sign-in'
import LayoutWithCredential from '@/components/layout-with-credential'

export default function Page({ params }: { params: { callbackUrl: string } }) {
  return (
    <LayoutWithCredential
      option="required"
      toBeDisplayed={false}
      callbackUrl={params.callbackUrl == '' ? '/' : params.callbackUrl}
    >
      <SignInForm callbackUrl={decodeURIComponent(params.callbackUrl)} />
    </LayoutWithCredential>
  )
}

'use client'
import SignUpForm from '@/components/forms/sign-up'
import LayoutWithCredential from '@/components/layout-with-credential'

export default function Page({ params }: { params: { callbackUrl: string } }) {
  return (
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
  )
}

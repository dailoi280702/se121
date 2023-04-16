import SignInForm from '@/components/forms/sign-in'

export default function Page({ params }: { params: { callbackUrl: string } }) {
  return <SignInForm callbackUrl={decodeURIComponent(params.callbackUrl)} />
}

import Link from 'next/link'

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="text-center">
      <h2 className="text-xl font-bold text-neutral-800 p-14">SIGN UP</h2>
      {children}
      <div className="px-6 py-14 text-sm sm:py-6 sm:border-neutral-200 sm:bg-neutral-100 sm:border-0 sm:border-t">
        already have an account?{' '}
        <Link className="text-teal-600" href="/auth/signin">
          Sign in
        </Link>{' '}
        to countinue
      </div>
    </div>
  )
}

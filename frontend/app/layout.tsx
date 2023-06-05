import Header from '@/components/Header'
import './globals.css'
import NavDrawer from '@/components/nav-drawer'
import UserProvider from '@/components/providers/user-provider'

export const metadata = {
  title: 'Carz',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <div className="bg-neutral-50 text-neutral-900 flex h-screen overflow-auto">
          <NavDrawer />
          <div className="flex-grow">
            <Header />
            <UserProvider>{children}</UserProvider>
          </div>
        </div>
      </body>
    </html>
  )
}

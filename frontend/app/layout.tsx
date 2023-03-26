import Header from '@/components/Header'
import './globals.css'
import NavDrawer from '@/components/nav-drawer'

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
        <div className="bg-neutral-100 text-neutral-900 flex">
          <NavDrawer />
          <div className="flex-grow">
            <Header />
            {children}
          </div>
        </div>
      </body>
    </html>
  )
}

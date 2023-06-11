'use client'

import { atom, useAtom } from 'jotai'
import { Logo } from '@/components/Header'
import useCloseShade from '@/components/hooks/use-close-shade'
import { Shade } from '@/components/shade'
import DrawerCloseButton from './buttons/drawer-close-button'

export const navDrawerVisisibilyAtom = atom<boolean>(false)

const DrawerHeader = ({ onClose }: { onClose: () => void }) => {
  return (
    <div className="flex h-16 w-full items-center justify-between pl-4 pr-2">
      <Logo />
      <DrawerCloseButton onClose={onClose} />
    </div>
  )
}

export const NavMenu = ({ label }: { label?: string }) => {
  return (
    <div className="min-h-min w-full space-y-2">
      <nav className="flex w-full flex-col">
        <a
          href="#"
          className="mb-2 rounded-md px-6 py-3 text-sm font-medium text-gray-600 hover:bg-gray-200"
        >
          {label ? label : 'required*'}
        </a>
      </nav>
    </div>
  )
}

const DrawerBody = () => {
  return (
    <div className="flex h-screen flex-col">
      <NavMenu label="Home" />
      <NavMenu label="Blog" />
      <NavMenu label="Car" />
      <NavMenu label="Brand" />
    </div>
  )
}

export default function NavDrawer() {
  const [navDrawerVisisibily] = useAtom(navDrawerVisisibilyAtom)
  const closeDrawer = useCloseShade(navDrawerVisisibilyAtom)

  return (
    <>
      {navDrawerVisisibily && (
        <>
          <nav
            className="fixed left-0 top-0 z-[6] flex h-full flex-col 
            bg-neutral-100 shadow shadow-neutral-400 md:static md:z-auto"
          >
            <DrawerHeader onClose={closeDrawer} />
            <DrawerBody />
            <span className="w-56 grow" />
            <div>Nav footer</div>
          </nav>
          <Shade onClose={closeDrawer} className="md:hidden" />
        </>
      )}
    </>
  )
}

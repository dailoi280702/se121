'use client'

import { atom, useAtom } from 'jotai'
import { Logo } from '@/components/Header'
import useCloseShade from '@/components/hooks/use-close-shade'
import { Shade } from '@/components/shade'
import DrawerCloseButton from './buttons/drawer-close-button'
import Link from 'next/link'

export const navDrawerVisisibilyAtom = atom<boolean>(false)

const DrawerHeader = ({ onClose }: { onClose: () => void }) => {
  return (
    <div className="flex h-16 w-full items-center justify-between pl-4 pr-2">
      <Logo />
      <DrawerCloseButton onClose={onClose} />
    </div>
  )
}

export const NavMenuItem = ({
  label,
  link,
}: {
  label: string
  link: string
}) => {
  return (
    <div className="flex w-full flex-col">
      <Link
        href={link}
        className="px-6 py-3 text-sm font-medium hover:bg-neutral-200"
      >
        {label ? label : 'required*'}
      </Link>
    </div>
  )
}

const DrawerBody = () => {
  return (
    <div className="flex h-screen flex-col">
      <NavMenuItem label="Home" link="/" />
      <NavMenuItem label="Blog" link="/blog" />
      <NavMenuItem label="Car" link="/car" />
      <NavMenuItem label="Brand" link="/brand" />
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
            border bg-white shadow shadow-neutral-400 md:static
            md:z-auto"
          >
            <DrawerHeader onClose={closeDrawer} />
            <DrawerBody />
            <span className="w-56 grow" />
            <div></div>
          </nav>
          <Shade onClose={closeDrawer} className="md:hidden" />
        </>
      )}
    </>
  )
}

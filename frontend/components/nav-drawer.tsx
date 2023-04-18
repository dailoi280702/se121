'use client'

import { atom, useAtom } from 'jotai'
import { Logo } from '@/components/Header'
import useCloseShade from '@/components/hooks/use-close-shade'
import { Shade } from '@/components/shade'
import DrawerCloseButton from './buttons/drawer-close-button'

export const navDrawerVisisibilyAtom = atom<boolean>(false)

const DrawerHeader = ({ onClose }: { onClose: () => void }) => {
  return (
    <div className="flex items-center justify-between h-16 w-full pl-4 pr-2">
      <Logo />
      <DrawerCloseButton onClose={onClose} />
    </div>
  )
}

const DrawerBody = () => {
  return <div>nav body</div>
}

export default function NavDrawer() {
  const [navDrawerVisisibily] = useAtom(navDrawerVisisibilyAtom)
  const closeDrawer = useCloseShade(navDrawerVisisibilyAtom)

  return (
    <>
      {navDrawerVisisibily && (
        <>
          <Shade onClose={closeDrawer} className="md:hidden" />
          <div
            className="absolute left-0 top-0 h-full w-56 flex flex-col z-[8] md:static
            bg-neutral-50 shadow shadow-neutral-400"
          >
            <DrawerHeader onClose={closeDrawer} />
            <DrawerBody />
            <div className="flex-grow" />
            <div>Nav footer</div>
          </div>
        </>
      )}
    </>
  )
}

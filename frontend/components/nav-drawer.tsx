'use client'

import { XMarkIcon } from '@heroicons/react/24/outline'
import { Logo } from '@/components/Header'
import { atom, useAtom, useSetAtom } from 'jotai'

export const navDrawerVisisibilyAtom = atom<boolean>(false)

const Shade = ({ onClose }: { onClose: () => void }) => {
  return (
    <div
      className="fixed left-0 top-0 h-full w-full z-[7] md:hidden bg-black/40"
      onClick={onClose}
    />
  )
}

const DrawerHeader = ({ onClose }: { onClose: () => void }) => {
  return (
    <div className="flex items-center justify-between h-16 w-full pl-4 pr-2">
      <Logo />
      <button
        className="h-10 w-10 flex items-center justify-center rounded-full font-bold
        hover:bg-neutral-700 hover:bg-opacity-[0.08]"
        onClick={onClose}
      >
        <XMarkIcon className="h-6 w-6 stroke-2" />
      </button>
    </div>
  )
}

const DrawerBody = () => {
  return <div>nav body</div>
}

export default function NavDrawer() {
  const [navDrawerVisisibily, setNavDrawerVisiblity] = useAtom(
    navDrawerVisisibilyAtom
  )
  const closeDrawer = () => setNavDrawerVisiblity(false)

  return (
    <>
      {navDrawerVisisibily && (
        <>
          <Shade onClose={closeDrawer} />
          <div
            className="absolute left-0 top-0 h-full w-56 flex flex-col z-[8] md:static
            bg-white shadow shadow-neutral-200
            transition-transform duration-75"
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

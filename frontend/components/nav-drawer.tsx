'use client'

import { XMarkIcon } from '@heroicons/react/24/outline'
import { Logo } from '@/components/Header'

const Shade = () => {
  const closeDrawer = () => {
    window.alert('close drawer')
  }

  return (
    <div
      className="fixed left-0 top-0 h-full w-full z-[7] md:hidden bg-black/40"
      onClick={closeDrawer}
    />
  )
}

const DrawerHeader = () => {
  return (
    <div className="flex items-center justify-between h-16 w-full pl-4 pr-2">
      <Logo />
      <button
        className="h-10 w-10 flex items-center justify-center rounded-full font-bold
        hover:bg-neutral-700 hover:bg-opacity-[0.08]"
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
  return (
    <>
      <Shade />
      <div className="absolute left-0 top-0 h-full w-56 flex flex-col bg-neutral-200 z-[8] md:static">
        <DrawerHeader />
        <DrawerBody />
        <div className="flex-grow" />
        <div>Nav footer</div>
      </div>
    </>
  )
}

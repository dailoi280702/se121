'use client'

import { ChangeEvent, FormEvent, HTMLInputTypeAttribute, useState } from 'react'

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

export const NavMenu = ({ label }: { label?: string }) => {
  return (
    <div className="w-full space-y-2 min-h-min">
      <nav className="flex flex-col w-full">
        <a
          href="#"
          className="px-6 py-3 mb-2 text-sm font-medium text-gray-600 rounded-md hover:bg-gray-200"
        >
          {label ? label : 'required*'}
        </a>
        {/* <a
          href="#"
          className="px-6 py-3 mb-2 text-sm font-medium text-gray-600 rounded-md hover:bg-gray-200"
        >
          {label ? label : 'required*'}
        </a> */}
      </nav>
    </div>
  )
}

const DrawerBody = () => {
  return (
    <div className="flex flex-col h-screen">
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

'use client'

import {
  Bars3Icon,
  UserCircleIcon,
  MagnifyingGlassIcon,
} from '@heroicons/react/24/outline'
import { useAtom, useSetAtom } from 'jotai'
import { useRouter } from 'next/navigation'
import { navDrawerVisisibilyAtom } from '@/components/nav-drawer'
import ProfileContainer, {
  profileContainerVisisibilyAtom,
} from '../profile-container'
import { useState } from 'react'

const NavButton = ({ onClick }: { onClick: () => void }) => {
  return (
    <button
      className="h-10 w-10 flex items-center justify-center mr-2 rounded-full font-bold
      hover:bg-neutral-700 hover:bg-opacity-[0.08]"
      onClick={onClick}
    >
      <Bars3Icon className="h-6 w-6 stroke-2" />
    </button>
  )
}

export const Logo = () => {
  const router = useRouter()

  return (
    <p
      className="text-lg font-semibold cursor-pointer"
      onClick={() => {
        router.push('/')
      }}
    >
      CARZ
    </p>
  )
}

const Search = () => (
  <div className="">
    <form
      className="z-[6] relative flex items-center
      bg-neutral-100 px-4 h-9 rounded-lg
      focus-within:ring-2 ring-neutral-600"
    >
      <input
        placeholder="Search"
        className="bg-transparent outline-none placeholder:text-neutral-600"
      />
      <MagnifyingGlassIcon className="h-5 w-5 ml-1 stroke-2" />
    </form>

    {/* shader */}
    {/* <div className='fixed bg-black/40 z-[5] left-0 right-0 top-0 bottom-0' /> */}
  </div>
)

const User = () => {
  const [isMenuOpen, setMenuOpen] = useState(false)

  const closeMenu = () => {
    setMenuOpen(false)
    console.log(isMenuOpen)
  }

  const openMenu = () => {
    setMenuOpen(true)
  }

  return (
    <div className="relative grid-flow-col-dense grid">
      {isMenuOpen && (
        <ProfileContainer isOpen={isMenuOpen} onClose={closeMenu} />
      )}
      <button onClick={openMenu}>
        <UserCircleIcon className="h-10 w-10 stroke-1"></UserCircleIcon>
      </button>
    </div>
  )
}

const Header = () => {
  const [navDrawerVisisibily, setNavDrawerVisiblity] = useAtom(
    navDrawerVisisibilyAtom
  )
  const [isProfileOpen, setProfileOpen] = useState(false)
  const openDrawer = () => setNavDrawerVisiblity(true)

  return (
    <div className="h-16 bg-neutral-50 shadow-neutral-200 shadow flex items-center justify-center">
      <div className="max-w-6xl w-full flex items-center px-4">
        {!navDrawerVisisibily && (
          <>
            <NavButton onClick={openDrawer} />
            <Logo />
          </>
        )}
        <div className="flex-grow" />
        <Search />
        <p className="ml-2" />
        <User />
      </div>
    </div>
  )
}

export default Header

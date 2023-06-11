'use client'

import {
  Bars3Icon,
  UserCircleIcon,
  MagnifyingGlassIcon,
} from '@heroicons/react/24/outline'
import { atom, useAtom, useSetAtom } from 'jotai'
import { useRouter } from 'next/navigation'
import { navDrawerVisisibilyAtom } from '@/components/nav-drawer'
import ProfileContainer from '../profile-container'
import GlobalSearch from '../global-search'
import { Shade } from '../shade'
import useCloseShade from '../hooks/use-close-shade'

export const profileVisisibilyAtom = atom<boolean>(false)

const NavButton = ({ onClick }: { onClick: () => void }) => {
  return (
    <button
      className="mr-2 flex h-10 w-10 items-center justify-center rounded-full font-bold
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
      className="mr-4 hidden cursor-pointer text-lg font-semibold sm:inline-block"
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
      className="relative z-[6] flex h-9 w-9
      items-center justify-center rounded-lg ring-neutral-600 focus-within:ring-2 sm:w-min
      sm:bg-neutral-100 sm:px-2"
    >
      <input
        placeholder="Search"
        className="hidden bg-transparent outline-none placeholder:text-neutral-600 sm:inline-block"
      />
      <MagnifyingGlassIcon className="ml-1 h-5 w-5 stroke-2" />
    </form>

    {/* shader */}
    {/* <div className='fixed bg-black/40 z-[5] left-0 right-0 top-0 bottom-0' /> */}
  </div>
)

const User = () => {
  const setProfileVisibility = useSetAtom(profileVisisibilyAtom)

  const openMenu = () => {
    setProfileVisibility(true)
  }

  return (
    <div className="z-[9] grid grid-flow-col-dense sm:relative">
      <ProfileContainer />
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
  const closeProfile = useCloseShade(profileVisisibilyAtom, true)
  const openDrawer = () => setNavDrawerVisiblity(true)
  const [profileVisisibily] = useAtom(profileVisisibilyAtom)

  return (
    <>
      <div className="sticky top-0 z-[1] flex h-16 w-full items-center justify-center bg-neutral-50 shadow shadow-neutral-300">
        <div className="flex w-full max-w-6xl items-center px-4">
          {!navDrawerVisisibily && (
            <>
              <NavButton onClick={openDrawer} />
              <Logo />
            </>
          )}
          <GlobalSearch />
          <p className="ml-2" />
          {profileVisisibily && (
            <Shade onClose={closeProfile} className="z-[8]" />
          )}
          <User />
        </div>
      </div>
    </>
  )
}

export default Header

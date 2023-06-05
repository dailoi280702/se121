import DrawerCloseButton from './buttons/drawer-close-button'
import {
  ArrowLeftOnRectangleIcon,
  ArrowRightOnRectangleIcon,
  Cog6ToothIcon,
  UserCircleIcon,
  UserIcon,
} from '@heroicons/react/24/outline'
import { cloneElement, ReactElement } from 'react'
import { usePathname, useRouter } from 'next/navigation'
import { useAtom, useAtomValue } from 'jotai'
import { profileVisisibilyAtom } from '@/components/Header'
import { UserAtom } from './providers/user-provider'

const MenuButton = ({
  text,
  children,
  onClick,
  className = '',
}: {
  onClick: () => void
  text?: string
  children?: ReactElement
  className?: string
}) => {
  return (
    <button
      className={
        `flex items-center text-sm h-8 rounded-md gap-x-2
      text-black-600 hover:bg-neutral-600/[0.08] px-2 ` + className
      }
      onClick={onClick}
    >
      {children && cloneElement(children, { className: 'h-6 w-6 stroke-2' })}
      {text}
    </button>
  )
}

const ProfileMenu = () => {
  const router = useRouter()
  const pathName = usePathname()
  const [user, setUser] = useAtom(UserAtom)

  const buttons = [
    {
      name: 'My Profile',
      icon: <UserIcon />,
      url: '/',
    },
    {
      name: 'Settings',
      icon: <Cog6ToothIcon />,
      url: '/',
    },
    {
      name: 'Sign In',
      icon: <ArrowLeftOnRectangleIcon />,
      url: '/auth/signin/',
      onClick: () => router.push('auth/signin/' + encodeURIComponent(pathName)),
      className: ' hover:bg-teal-600 hover:text-teal-50',
      displayCondition: { authenticated: false },
    },
    {
      name: 'Sign Up',
      icon: <ArrowLeftOnRectangleIcon />,
      url: '/auth/singup/%2f',
      onClick: () => router.push('auth/signup/' + encodeURIComponent(pathName)),
      displayCondition: { authenticated: false },
    },
    {
      name: 'Sign Out',
      icon: <ArrowRightOnRectangleIcon />,
      url: '/',
      onClick: () => signOut(),
      displayCondition: { authenticated: true },
    },
  ]

  const signOut = async () => {
    const response = await fetch('http://localhost:8000/v1/auth', {
      method: 'DELETE',
      credentials: 'include',
    })

    if (!response.ok) {
      window.alert(await response.text())
    }

    window.location.reload()
    setUser(null)
  }

  return (
    <ul className="flex flex-col space-y-1">
      {buttons.map(
        (button) =>
          (!button.displayCondition ||
            button.displayCondition.authenticated == !!user) && (
            <MenuButton
              key={button.name}
              text={button.name}
              onClick={() => {
                button.onClick ? button.onClick() : router.push(button.url)
              }}
              className={button.className}
            >
              {button.icon}
            </MenuButton>
          )
      )}
    </ul>
  )
}

const Info = () => {
  const user = useAtomValue(UserAtom)

  return (
    <button
      className="rounded-md
      text-black-600 hover:bg-neutral-600/[0.08] px-4 py-1 sm:!mt-0"
    >
      <div className="w-28 sm:w-20 rouned-full mx-auto">
        <UserCircleIcon className="stroke-[0.6]" />
      </div>
      <p className="text-ellipsis overflow-hidden hover:break-words">
        {user ? user.name : 'guest'}
      </p>
    </button>
  )
}

const Divider = ({ className = '' }: { className?: string }) => {
  return <hr className={'border-neutral-200 border-0 border-t ' + className} />
}

export default function ProfileContainer() {
  const [profileVisisibily, setProfileVisibility] = useAtom(
    profileVisisibilyAtom
  )
  const closeProfile = () => {
    if (profileVisisibily) {
      setProfileVisibility(false)
    }
  }

  return (
    <>
      {profileVisisibily && (
        <>
          <div
            className="fixed z-[6] right-0 top-0 h-full w-56 flex flex-col space-y-4 px-2
            sm:absolute sm:h-fit sm:top-full sm:rounded-lg sm:py-4
            bg-neutral-50 shadow shadow-neutral-200"
          >
            <div className="flex items-center justify-between sm:hidden h-16 w-full pr-4 pl-2">
              <DrawerCloseButton onClose={closeProfile} />
            </div>
            <Divider className="!mt-0 sm:hidden" />
            <Info />
            <Divider />
            <ProfileMenu />
          </div>
        </>
      )}
    </>
  )
}

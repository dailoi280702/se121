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
        `text-black-600 flex h-8 items-center gap-x-2 rounded-md
      px-2 text-sm hover:bg-neutral-600/[0.08] ` + className
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
      onClick: () =>
        router.push('/auth/signin/' + encodeURIComponent(pathName)),
      className: ' hover:bg-teal-600 hover:text-teal-50',
      displayCondition: { authenticated: false },
    },
    {
      name: 'Sign Up',
      icon: <ArrowLeftOnRectangleIcon />,
      url: '/auth/singup/%2f',
      onClick: () =>
        router.push('/auth/signup/' + encodeURIComponent(pathName)),
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
    <button className="rounded-md px-4 py-1 text-neutral-600 hover:bg-neutral-600/[0.08] sm:!mt-0">
      <div className="mx-auto w-28 rounded-full sm:w-20">
        <UserCircleIcon className="stroke-[0.6]" />
      </div>
      <p className="overflow-hidden text-ellipsis hover:break-words">
        {user ? user.name : 'guest'}
      </p>
    </button>
  )
}

const Divider = ({ className = '' }: { className?: string }) => {
  return <hr className={'border-0 border-t border-neutral-200 ' + className} />
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
            className="fixed right-0 top-0 z-[6] flex h-full w-56 flex-col space-y-4 bg-neutral-50
            px-2 shadow shadow-neutral-200 sm:absolute sm:top-full
            sm:h-fit sm:rounded-lg sm:py-4"
          >
            <div className="flex h-16 w-full items-center justify-between pl-2 pr-4 sm:hidden">
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

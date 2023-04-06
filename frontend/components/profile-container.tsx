import { Shade } from '@/components/shade'
import DrawerCloseButton from './buttons/drawer-close-button'
import {
  ArrowLeftOnRectangleIcon,
  Cog6ToothIcon,
  UserCircleIcon,
  UserIcon,
} from '@heroicons/react/24/outline'
import { cloneElement, ReactElement } from 'react'
import { useRouter } from 'next/navigation'
import { useAtom } from 'jotai'
import useCloseShade from './hooks/use-close-shade'
import { profileVisisibilyAtom } from '@/components/Header'

const MenuButton = ({
  text,
  children,
  onClick,
}: {
  onClick: () => void
  text?: string
  children?: ReactElement
}) => {
  return (
    <button
      className="flex items-center text-sm h-8 rounded-md gap-x-2
      text-black-600 hover:bg-neutral-600/[0.08] px-2"
      onClick={onClick}
    >
      {children && cloneElement(children, { className: 'h-6 w-6 stroke-2' })}
      {text}
    </button>
  )
}

const ProfileMenu = () => {
  const router = useRouter()

  const buttons = [
    { name: 'My Profile', icon: <UserIcon />, url: '/' },
    { name: 'Settings', icon: <Cog6ToothIcon />, url: '/' },
    {
      name: 'Sign In',
      icon: <ArrowLeftOnRectangleIcon />,
      url: '/auth/signin',
    },
  ]

  return (
    <ul className="flex flex-col space-y-1">
      {buttons.map((button) => (
        <MenuButton
          key={button.name}
          text={button.name}
          onClick={() => {
            router.push(button.url)
          }}
        >
          {button.icon}
        </MenuButton>
      ))}
    </ul>
  )
}

const Info = () => {
  return (
    <button
      className="rounded-md
      text-black-600 hover:bg-neutral-600/[0.08] px-4 py-1 sm:!mt-0"
    >
      <div className="w-28 sm:w-20 rouned-full mx-auto">
        <UserCircleIcon className="stroke-[0.6]" />
      </div>
      <p className="text-ellipsis overflow-hidden hover:break-words">
        a very long long ccs
      </p>
    </button>
  )
}

const Divider = ({ className = '' }: { className?: string }) => {
  return <hr className={'border-neutral-200 border-0 border-t ' + className} />
}

export default function ProfileContainer() {
  const [profileVisisibily] = useAtom(profileVisisibilyAtom)
  const closeProfile = useCloseShade(profileVisisibilyAtom, true)

  return (
    <>
      {profileVisisibily && (
        <>
          <Shade onClose={closeProfile} className="z-[8]" />
          <div
            className="fixed right-0 top-0 h-full w-56 flex flex-col z-[8] space-y-4 px-2
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

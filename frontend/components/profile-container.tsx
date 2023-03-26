import { Shade } from '@/components/shade'
import DrawerCloseButton from './buttons/drawer-close-button'
import {
  ArrowLeftOnRectangleIcon,
  UserCircleIcon,
} from '@heroicons/react/24/outline'

const LoginButton = ({
  text,
  onClick,
}: {
  text?: string
  onClick: () => void
}) => {
  return (
    <button
      className="flex items-center text-md h-8 rounded-md gap-x-2
      text-black-600 hover:bg-neutral-600/[0.08] px-2"
      onClick={onClick}
    >
      <ArrowLeftOnRectangleIcon className="h-6 w-6 stroke-2" />
      {text}
    </button>
  )
}

const ProfileMenu = () => {
  return (
    <ul className="flex flex-col space-y-1">
      <LoginButton text="Sign In" onClick={() => {}} />
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
        a very long long
        nameeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee
      </p>
    </button>
  )
}

const Devider = ({ className = '' }: { className?: string }) => {
  return <hr className={'border-neutral-200 border-0 border-t ' + className} />
}

export default function ProfileContainer({
  isOpen,
  onClose,
}: {
  isOpen: boolean
  onClose: () => void
}) {
  return (
    <>
      {isOpen && (
        <>
          <Shade onClose={onClose} className="z-[8]" />
          <div
            className="absolute right-0 top-0 h-full w-56 flex flex-col z-[8] space-y-4 px-2
            sm:h-fit sm:top-full sm:rounded-lg sm:py-4
            bg-neutral-50 shadow shadow-neutral-200"
          >
            <div className="flex items-center justify-between sm:hidden h-16 w-full pr-4 pl-2">
              <DrawerCloseButton onClose={onClose} />
            </div>
            <Devider className="!mt-0 sm:hidden" />
            <Info />
            <Devider />
            <ProfileMenu />
          </div>
        </>
      )}
    </>
  )
}

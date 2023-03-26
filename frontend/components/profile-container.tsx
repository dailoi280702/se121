import { Shade } from '@/components/shade'
import DrawerCloseButton from './buttons/drawer-close-button'
import { ArrowLeftOnRectangleIcon } from '@heroicons/react/24/outline'

const LoginButton = () => {
  return (
    <button
      className="flex items-center text-lg h-8 rounded-md gap-x-2
      hover:bg-neutral-600/[0.08] px-3"
    >
      <ArrowLeftOnRectangleIcon className="h-6 w-6 stroke-2" />
      Sign In
    </button>
  )
}

const ProfileMenu = () => {
  return (
    <>
      <LoginButton />
      <LoginButton />
      <LoginButton />
      <LoginButton />
    </>
  )
}

const Devider = () => {
  return
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
            className="absolute right-0 top-0 h-full w-56 flex flex-col z-[8] space-y-2
            sm:h-fit sm:top-full sm:rounded-lg sm:p-2
            bg-neutral-50 shadow shadow-neutral-200"
          >
            <div className="flex items-center justify-between sm:hidden h-16 w-full pr-4 pl-2">
              <DrawerCloseButton onClose={onClose} />
            </div>
            <ProfileMenu />
          </div>
        </>
      )}
    </>
  )
}

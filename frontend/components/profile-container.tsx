import { Shade } from '@/components/shade'
import DrawerCloseButton from './buttons/drawer-close-button'

const ProfileMenu = () => {
  return <div>menu</div>
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
            className="absolute right-0 top-0 h-full w-56 flex flex-col z-[8]
            sm:absolute sm:h-fit sm:top-full sm:rounded-lg
            bg-neutral-50 shadow shadow-neutral-200"
          >
            <div className="flex items-center justify-between h-16 w-full pr-4 pl-2">
              <DrawerCloseButton onClose={onClose} />
            </div>
            <ProfileMenu />
          </div>
        </>
      )}
    </>
  )
}

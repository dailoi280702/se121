'use client'

import { atom, useAtom } from 'jotai'
import { Shade } from '@/components/shade'
import useCloseShade from './hooks/use-close-shade'

export const profileContainerVisisibilyAtom = atom<boolean>(false)

const ProfileMenu = () => {
  return <div>menu</div>
}

export default function ProfileContainer() {
  const [profileContainerVisisibily] = useAtom(profileContainerVisisibilyAtom)
  const onClose = useCloseShade(profileContainerVisisibilyAtom, true)

  return (
    <>
      {profileContainerVisisibily && (
        <>
          <Shade onClose={onClose} />
          <div
            className="absolute right-0 top-0 h-full w-56 flex flex-col z-[8]
            bg-neutral-50 shadow shadow-neutral-200"
          >
            <ProfileMenu />
          </div>
        </>
      )}
    </>
  )
}

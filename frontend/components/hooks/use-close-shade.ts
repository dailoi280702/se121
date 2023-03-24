import { PrimitiveAtom, useSetAtom } from 'jotai'
import { useCallback, useEffect } from 'react'
import { usePathname } from 'next/navigation'

const smallScreenLimit = 640

export default function useCloseShade(
  shadeAtom: PrimitiveAtom<boolean>,
  always: boolean = false,
  limitWidth: number = smallScreenLimit
) {
  const setShadeVisibility = useSetAtom(shadeAtom)
  const closeShade = useCallback(
    () => setShadeVisibility(false),
    [setShadeVisibility]
  )
  const path = usePathname()

  useEffect(() => {
    if (always || window.innerWidth <= limitWidth) {
      closeShade()
    }
  }, [path, closeShade, limitWidth, always])

  return closeShade
}

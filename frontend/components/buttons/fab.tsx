'use client'

import { atom, useAtom } from 'jotai'

export const isFabOpenAtom = atom(false)

const Fab = ({
  child,
  icon,
}: {
  child: (closeFab: () => void) => React.ReactNode
  icon: React.ReactNode
}) => {
  const [isFabOpen, setIsFabOpen] = useAtom(isFabOpenAtom)

  const handleFabClick = () => {
    setIsFabOpen(!isFabOpen)
  }

  const closeFab = () => setIsFabOpen(false)

  return (
    <>
      {isFabOpen && (
        <div
          id="fab-shad"
          className="absolute z-[2] top-0 bottom-0 left-0 right-0 bg-black/40 h-screen flex items-center"
          onClick={(e) => {
            if (e.currentTarget === e.target) {
              closeFab()
            }
          }}
        >
          {child(closeFab)}
        </div>
      )}
      <div className="fixed z-[2] bottom-4 right-4">
        <button
          className="bg-teal-600 w-14 h-14 hover:bg-teal-700 text-white font-bold py-2 px-4 rounded-2xl shadow shadow-neutral-900"
          onClick={handleFabClick}
        >
          {icon}
        </button>
      </div>
    </>
  )
}

export default Fab

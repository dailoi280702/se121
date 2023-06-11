export default function Fab({
  children,
  icon,
  isFabOpen,
  setIsFabOpen,
}: {
  children: React.ReactNode
  icon: React.ReactNode
  isFabOpen: boolean
  setIsFabOpen: (value: boolean) => void
}) {
  const handleFabClick = () => {
    setIsFabOpen(!isFabOpen)
  }

  const closeFab = () => setIsFabOpen(false)

  return (
    <>
      {isFabOpen && (
        <div
          id="fab-shad"
          className="absolute inset-0 z-[2] flex h-screen items-center overflow-y-scroll bg-black/40"
          onClick={(e) => {
            if (e.currentTarget === e.target) {
              closeFab()
            }
          }}
        >
          {children}
        </div>
      )}
      <div className="fixed bottom-4 right-4 z-[2]">
        <button
          className="h-14 w-14 rounded-2xl bg-teal-600 px-4 py-2 font-bold text-white shadow shadow-neutral-900 hover:bg-teal-700"
          onClick={handleFabClick}
        >
          {icon}
        </button>
      </div>
    </>
  )
}

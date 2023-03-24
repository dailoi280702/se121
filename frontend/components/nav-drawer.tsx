'use client'

const Shade = () => {
  const closeDrawer = () => {
    window.alert('close drawer')
  }

  return (
    <div
      className="fixed left-0 top-0 h-full w-full z-[7] md:hidden bg-black/40"
      onClick={closeDrawer}
    />
  )
}

export default function NavDrawer() {
  return (
    <>
      <Shade />
      <div className="absolute left-0 top-0 h-full w-56 flex flex-col bg-neutral-200 z-[8] md:static">
        <div>Nav header</div>
        <div className="flex-grow" />
        <div>Nav footer</div>
      </div>
    </>
  )
}

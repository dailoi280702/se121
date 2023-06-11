export function Shade({
  onClose,
  className = '',
}: {
  onClose: () => void
  className?: string
}) {
  return (
    <div
      className={
        'fixed left-0 right-0 top-0 z-[5] h-full w-full bg-black/40 ' +
        className
      }
      onClick={onClose}
    />
  )
}

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
        'fixed left-0 right-0 top-0 h-full w-full z-[5] bg-black/40 ' +
        className
      }
      onClick={onClose}
    />
  )
}

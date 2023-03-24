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
        'fixed left-0 top-0 h-full w-full z-[7] bg-black/40 ' + className
      }
      onClick={onClose}
    />
  )
}

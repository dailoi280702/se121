import { XMarkIcon } from '@heroicons/react/24/outline'

export default function Tag({
  name,
  className = '',
  onIconClick,
}: {
  name: string
  className?: string
  onIconClick?: () => void
}) {
  return (
    <div
      className={[
        `flex h-8 items-center rounded-lg border border-neutral-200 
        px-3 text-center text-sm text-gray-900 focus:outline-none`,
        className,
      ].join(' ')}
    >
      {name}
      {onIconClick && (
        <XMarkIcon
          className="ml-2 h-5 w-5 cursor-pointer stroke-2"
          onClick={onIconClick}
        />
      )}
    </div>
  )
}

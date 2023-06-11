import { XMarkIcon } from '@heroicons/react/24/outline'

export default function DialogFormLayout({
  children,
  disabled,
  title,
  buttonLabel,
  onClose,
  onDone,
}: {
  children: React.ReactNode
  disabled?: boolean
  title: string
  buttonLabel: string
  onClose: () => void
  onDone: () => void
}) {
  return (
    <div
      className={[
        'm-auto min-h-screen w-full pb-20',
        'bg-neutral-50 sm:min-h-max',
        'sm:my-6 sm:h-fit sm:max-w-sm sm:rounded-3xl sm:pb-0',
      ].join(' ')}
    >
      <div className="flex h-14 items-center">
        <XMarkIcon className="mx-4 h-6 w-6 stroke-2" onClick={onClose} />
        <p className="text-[1.375rem]">{title}</p>
        <button
          className="ml-auto mr-6 h-10 rounded-full px-3 text-sm font-medium text-teal-600 
          outline-none enabled:hover:bg-teal-600/10"
          disabled={disabled}
          onClick={onDone}
        >
          {buttonLabel}
        </button>
      </div>
      <div className="p-6">{children}</div>
    </div>
  )
}

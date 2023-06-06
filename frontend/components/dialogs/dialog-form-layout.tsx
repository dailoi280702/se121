'use client'

import { XMarkIcon } from '@heroicons/react/24/outline'

const DialogFormLayout = ({
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
}) => {
  return (
    <div
      className={[
        'w-full h-full m-auto my-auto',
        'bg-neutral-50',
        'sm:rounded-3xl sm:max-w-sm sm:h-fit',
      ].join(' ')}
    >
      <div className="h-14 className = flex items-center sticky">
        <XMarkIcon className="mx-4 w-6 h-6 stroke-2" onClick={onClose} />
        <p className="text-[1.375rem]">{title}</p>
        <button
          className="ml-auto mr-6 rounded-[2.5rem] text-sm font-medium h-10 px-3 outline-none 
          text-teal-600 enabled:hover:bg-teal-600/10"
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

export default DialogFormLayout

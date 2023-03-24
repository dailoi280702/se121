import { XMarkIcon } from '@heroicons/react/24/outline'

export default function DrawerCloseButton({
  onClose,
}: {
  onClose: () => void
}) {
  return (
    <button
      className="h-10 w-10 flex items-center justify-center rounded-full font-bold
        hover:bg-neutral-700 hover:bg-opacity-[0.08]"
      onClick={onClose}
    >
      <XMarkIcon className="h-6 w-6 stroke-2" />
    </button>
  )
}

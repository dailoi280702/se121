import AdminOnlyWrapper from '@/components/admin-only-wrapper'
import DialogFormLayout from '@/components/dialogs/dialog-form-layout'
import OutLineInput from '@/components/inputs/outline-input'
import { triggerFormUsingRef } from '@/utils'
import { useState } from 'react'
import useAddUpdateCar from '../hooks/use-add-update-car'

export default function AddUpdateCar({
  brand,
  series,
  type,
  children,
  isOpen,
  setIsOpen,
}: {
  brand?: Brand
  series?: Series
  type: 'update' | 'create'
  children: React.ReactNode
  isOpen?: boolean
  setIsOpen?: (isOpen: boolean) => void
}) {
  const [_isOpen, _setIsOpen] = useState(false)
  const { car, resetState, errors, onSubmit, onChange, formRef } =
    useAddUpdateCar({
      initData: series
        ? series
        : {
            id: 0,
            name: '',
            brand: brand,
          },
      type: type ? type : 'create',
      onSuccess: () => {
        close()
      },
    })

  const close = () => {
    _setIsOpen(false)
    if (setIsOpen) {
      setIsOpen(false)
    }
  }

  return (
    <AdminOnlyWrapper>
      {(_isOpen || isOpen) && (
        <dialog
          className="fixed inset-0 z-[2] flex h-screen w-full items-center 
          overflow-y-scroll bg-black/40 p-0"
          onClick={(e) => {
            if (e.target === e.currentTarget) {
              close()
            }
          }}
        >
          <DialogFormLayout
            title={type === 'create' ? 'Add Car Model' : 'Update Car Model'}
            buttonLabel="Done"
            disabled={false}
            onClose={close}
            onDone={() => triggerFormUsingRef(formRef)}
          >
            <form
              className="flex flex-col"
              ref={formRef}
              onSubmit={(e) => {
                e.preventDefault()
                onSubmit()
              }}
            >
              <OutLineInput
                placeholder="Model name"
                label="name"
                name="name"
                value={car.name}
                onChange={onChange}
                errorMessage={errors.name}
                required
              />
              <div className="flex items-center justify-end space-x-2">
                <button
                  className="ml-auto mr-3 h-10 rounded-full px-3 text-sm 
                  font-medium text-red-600 outline-none enabled:hover:bg-red-600/10"
                  type="button"
                  onClick={() => resetState()}
                >
                  Reset
                </button>
                <button
                  className="mb-2 h-10 rounded-full bg-teal-600 px-5 py-2.5
                  text-sm font-medium text-white outline-none hover:bg-teal-700 
                  focus:outline-none"
                  type="submit"
                >
                  {type === 'create' ? 'Create Series' : 'Update'}
                </button>
              </div>
            </form>
          </DialogFormLayout>
        </dialog>
      )}
      <div onClick={() => _setIsOpen(true)}>{children}</div>
    </AdminOnlyWrapper>
  )
}

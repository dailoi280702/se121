'use client'

import DialogFormLayout from '@/components/dialogs/dialog-form-layout'
import OutLineInput from '@/components/inputs/outline-input'
import { triggerFormUsingRef } from '@/utils'
import { PlusIcon } from '@heroicons/react/24/outline'
import { useState } from 'react'
import useAddUpdateSeries from './use-add-update-series'

export default function CreateSeries({ brand }: { brand: Brand }) {
  const [isInteracting, setIsInteracting] = useState(false)
  const { name, resetState, errors, onSubmit, onChange, formRef } =
    useAddUpdateSeries({
      initData: {
        id: 0,
        name: '',
        brand: brand,
      },
      type: 'create',
      onSuccess: () => {
        setIsInteracting(false)
      },
    })

  return (
    <>
      {isInteracting && (
        <div
          className="absolute inset-0 z-[2] flex h-screen items-center 
          overflow-y-scroll bg-black/40"
          onClick={(e) => {
            if (e.target === e.currentTarget) {
              setIsInteracting(false)
            }
          }}
        >
          <DialogFormLayout
            title="Add seires"
            buttonLabel="Done"
            disabled={false}
            onClose={() => setIsInteracting(false)}
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
                placeholder="Serreis name"
                label="name"
                name="name"
                value={name}
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
                  Clear
                </button>
                <button
                  className="mb-2 h-10 rounded-full bg-teal-600 px-5 py-2.5
                  text-sm font-medium text-white outline-none hover:bg-teal-700 
                  focus:outline-none"
                  type="submit"
                >
                  Create Seires
                </button>
              </div>
            </form>
          </DialogFormLayout>
        </div>
      )}
      <PlusIcon
        className="h-6 w-6 stroke-2"
        onClick={() => setIsInteracting(true)}
      />
    </>
  )
}

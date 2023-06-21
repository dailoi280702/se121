'use client'

import DialogFormLayout from '@/components/dialogs/dialog-form-layout'
import { triggerFormUsingRef } from '@/utils'
import { PencilIcon, PlusIcon } from '@heroicons/react/24/outline'
import { useState } from 'react'
import useAddUpdateSeries from './use-add-update-series'

export default function CreateSeries({ brand }: { brand: Brand }) {
  const [isInteracting, setIsInteracting] = useState(false)
  const { name, onSubmit, onChange, formRef } = useAddUpdateSeries({
    initData: {
      id: 0,
      name: '',
      brand: brand,
    },
    type: 'create',
  })

  return (
    <>
      {isInteracting && (
        <div
          className="absolute inset-0 z-[2] flex h-screen items-center 
          overflow-y-scroll bg-black/40"
          onClick={() => setIsInteracting(false)}
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
              <input value={name} onChange={onChange} />
              <div>
                <button type="submit">Done</button>
                <button type="button" onClick={() => setIsInteracting(false)}>
                  Cancel
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

'use client'

// import { PencilIcon } from '@heroicons/react/24/outline'
// import AdminOnlyWrapper from '../admin-only-wrapper'
// import Fab from '../buttons/fab'
// import DialogFormLayout from '../dialogs/dialog-form-layout'
// import AddBrandFrom from './add-brand-forms'

import AdminOnlyWrapper from '../admin-only-wrapper'
import { PencilIcon } from '@heroicons/react/24/outline'
import dynamic from 'next/dynamic'
import Loading from '@/app/loading'
import useAddEditBrand from '../hooks/use-add-brand'
import { triggerFormUsingRef } from '@/utils'
import { useState } from 'react'

const Fab = dynamic(() => import('../buttons/fab'))
const DialogFormLayout = dynamic(
  () => import('../dialogs/dialog-form-layout'),
  { loading: Loading, ssr: false }
)
const AddBrandForm = dynamic(() => import('./add-brand-forms'))

const AddBrandFromWithFab = () => {
  const [isFabOpen, setIsFabOpen] = useState(false)
  const hook = useAddEditBrand({
    type: 'update',
    onSuccess: () => setIsFabOpen(false),
  })

  return (
    <AdminOnlyWrapper>
      <Fab
        icon={<PencilIcon />}
        isFabOpen={isFabOpen}
        setIsFabOpen={setIsFabOpen}
      >
        <DialogFormLayout
          title="Add brand"
          buttonLabel="Done"
          disabled={false}
          onDone={() => triggerFormUsingRef(hook.formRef)}
          onClose={() => setIsFabOpen(false)}
        >
          <AddBrandForm hook={hook} />
        </DialogFormLayout>
      </Fab>
    </AdminOnlyWrapper>
  )
}
export default AddBrandFromWithFab

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

const Fab = dynamic(() => import('../buttons/fab'))
const DialogFormLayout = dynamic(
  () => import('../dialogs/dialog-form-layout'),
  { loading: Loading }
)
const AddBrandFrom = dynamic(() => import('./add-brand-forms'))

const AddBrandFromWithFab = () => {
  return (
    <AdminOnlyWrapper>
      <Fab
        icon={<PencilIcon />}
        child={(closeFab) => (
          <DialogFormLayout
            title="Add brand"
            buttonLabel="Done"
            disabled={false}
            onDone={() => {}}
            onClose={closeFab}
          >
            <AddBrandFrom />
          </DialogFormLayout>
        )}
      />
    </AdminOnlyWrapper>
  )
}
export default AddBrandFromWithFab

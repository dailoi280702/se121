'use client'

import AdminOnlyWrapper from '@/components/admin-only-wrapper'
import Fab from '@/components/buttons/fab'
import AddUpdateCar from '@/components/forms/add-update-car'
import { PencilIcon } from '@heroicons/react/24/outline'
import { useState } from 'react'

export default function UpdateCarFAB({ car }: { car: Car }) {
  const [isOpen, setIsOpen] = useState(false)

  return (
    <AdminOnlyWrapper>
      <Fab isFabOpen={isOpen} setIsFabOpen={setIsOpen} icon={<PencilIcon />}>
        <AddUpdateCar
          type="update"
          isOpen={isOpen}
          setIsOpen={setIsOpen}
          brand={car.brand}
          series={car.series}
          car={car}
        >
          <></>
        </AddUpdateCar>
      </Fab>
    </AdminOnlyWrapper>
  )
}

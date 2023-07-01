'use client'

import Fab from '@/components/buttons/fab'
import AddUpdateCar from '@/components/forms/add-update-car'
import { PencilIcon } from '@heroicons/react/24/outline'
import { useState } from 'react'

export default function UpdateCarFAB({ car }: { car: Car }) {
  const [isOpen, setIsOpen] = useState(false)
  return (
    <AddUpdateCar
      type="update"
      isOpen={isOpen}
      setIsOpen={setIsOpen}
      brand={car.brand}
      series={car.series}
      car={car}
    >
      <Fab isFabOpen={false} setIsFabOpen={() => {}} icon={<PencilIcon />}>
        <></>
      </Fab>
    </AddUpdateCar>
  )
}

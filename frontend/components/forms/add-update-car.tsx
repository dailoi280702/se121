/* eslint-disable react/jsx-no-undef */
import AdminOnlyWrapper from '@/components/admin-only-wrapper'
import DialogFormLayout from '@/components/dialogs/dialog-form-layout'
import OutLineInput from '@/components/inputs/outline-input'
import { triggerFormUsingRef } from '@/utils'
import { useState } from 'react'
import useAddUpdateCar from '../hooks/use-add-update-car'
import OutLineOptionMenu from '../menu/OutlineOptionMenu'

export default function AddUpdateCar({
  car,
  brand,
  series,
  type,
  children,
  isOpen,
  setIsOpen,
}: {
  car?: Car
  brand?: Brand
  series?: Series
  type: 'update' | 'create'
  children: React.ReactNode
  isOpen?: boolean
  setIsOpen?: (isOpen: boolean) => void
}) {
  const [_isOpen, _setIsOpen] = useState(false)
  const currentYear = new Date().getFullYear()
  const {
    car: _car,
    fuelTypes,
    transmissions,
    resetState,
    errors,
    onSubmit,
    onChange,
    onFuelTypeChange,
    onTransmissionChange,
    formRef,
  } = useAddUpdateCar({
    initData: car
      ? car
      : {
          id: 0,
          brand: brand,
          series: series,
          name: '',
          year: currentYear,
          horsePower: undefined,
          torque: undefined,
          transmission: undefined,
          fuelType: undefined,
          imageUrl: undefined,
          review: undefined,
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

  const years = Array.from(
    { length: currentYear - 1900 + 1 },
    (_, index) => currentYear - index
  )

  return (
    <AdminOnlyWrapper>
      {(_isOpen || isOpen) && (
        <dialog
          className="fixed inset-0 z-[2] flex h-screen w-full items-center 
          overflow-y-hidden bg-black/40 p-0"
          onClick={(e) => {
            if (e.target === e.currentTarget) {
              close()
            }
          }}
        >
          <DialogFormLayout
            className="sm:max-w-prose"
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
              <div className="flex">
                <p />
                <div className="w-1/2 space-y-2">
                  <OutLineInput
                    placeholder="name"
                    label="Model name"
                    name="name"
                    value={_car.name}
                    onChange={onChange}
                    errorMessage={errors.name}
                    required
                  />
                  <OutLineInput
                    label="Torque"
                    placeholder="torque"
                    name="torue"
                    value={_car.torque}
                    onChange={onChange}
                    errorMessage={errors.torque}
                    type="number"
                    min={0}
                    step={1}
                  />
                  <OutLineOptionMenu
                    label="Transmission type"
                    name="transmission"
                    options={transmissions}
                    value={
                      _car.transmission?.name ? _car.transmission.name : ''
                    }
                    onChange={onTransmissionChange}
                    errorMessage={errors.transmission}
                  />
                </div>
                <div className="ml-4 w-1/2 space-y-2">
                  <OutLineOptionMenu
                    label="Release year"
                    name="year"
                    options={years.map((y) => y.toString())}
                    value={_car.year}
                    onChange={onChange}
                    errorMessage={errors.year}
                  />
                  <OutLineInput
                    label="Power"
                    placeholder="horse power"
                    name="horsePower"
                    value={_car.horsePower}
                    onChange={onChange}
                    errorMessage={errors.horsePower}
                    type="number"
                    min={0}
                    step={1}
                  />
                  <OutLineOptionMenu
                    label="Fuel type"
                    name="fuelType"
                    placeholder=""
                    options={fuelTypes}
                    value={_car.fuelType?.name ? _car.fuelType.name : ''}
                    onChange={onFuelTypeChange}
                    errorMessage={errors.fuelType}
                  />
                </div>
              </div>
              <div className="flex items-center justify-end space-x-2">
                <button
                  className="ml-auto mr-3 h-10 rounded-full px-3 text-sm 
                  font-medium text-red-600 outline-none 
                  enabled:hover:bg-red-600/10"
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
                  {type === 'create' ? 'Create Model' : 'Update'}
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

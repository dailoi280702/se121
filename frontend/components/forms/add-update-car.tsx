import AdminOnlyWrapper from '@/components/admin-only-wrapper'
import DialogFormLayout from '@/components/dialogs/dialog-form-layout'
import OutLineInput from '@/components/inputs/outline-input'
import { triggerFormUsingRef } from '@/utils'
import { CameraIcon } from '@heroicons/react/24/outline'
import { useRef, useState } from 'react'
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
  const imageUpLoadRef = useRef<HTMLInputElement>(null)
  const currentYear = new Date().getFullYear()
  const {
    car: _car,
    fuelTypes,
    transmissions,
    errors,
    formRef,
    isGenerateingReview,
    isSubmitting,
    selectedImage,
    resetState,
    onSubmit,
    onChange,
    onImageChange,
    onFuelTypeChange,
    onTransmissionChange,
    generateReview,
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
    series: series,
    onSuccess: () => {
      close()
    },
  })

  const close = () => {
    _setIsOpen(false)
    if (setIsOpen) {
      setIsOpen(false)
    }
    resetState()
  }

  const years = Array.from(
    { length: currentYear - 1900 + 1 },
    (_, index) => currentYear - index
  )

  return (
    <AdminOnlyWrapper>
      {(_isOpen || isOpen) && (
        <dialog
          className="fixed inset-0 z-[2] flex h-screen w-full flex-col
          overflow-y-auto bg-neutral-50 p-0 sm:bg-black/40"
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
                className="flex flex-col pb-16 sm:pb-0"
                ref={formRef}
                onSubmit={(e) => {
                  e.preventDefault()
                  onSubmit()
                }}
              >
                <div className="flex">
                  <div className="w-1/2 space-y-2">
                    <OutLineInput
                      placeholder="brand"
                      label="Brand"
                      value={car ? car.brand?.name : brand?.name}
                      disabled
                    />
                    <OutLineInput
                      placeholder="name"
                      label="Model name"
                      name="name"
                      value={_car.name ? _car.name : ''}
                      onChange={onChange}
                      errorMessage={errors.name}
                      required
                    />
                    <OutLineInput
                      label="Torque"
                      placeholder="torque"
                      name="torue"
                      value={_car.torque ? _car.torque : ''}
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
                    <OutLineInput
                      placeholder="series"
                      label="Series"
                      value={car ? car.series?.name : series?.name}
                      disabled
                    />
                    <OutLineOptionMenu
                      label="Release year"
                      name="year"
                      options={years.map((y) => y.toString())}
                      value={_car.year ? _car.year : ''}
                      onChange={onChange}
                      errorMessage={errors.year}
                    />
                    <OutLineInput
                      label="Power"
                      placeholder="horse power"
                      name="horsePower"
                      value={_car.horsePower ? _car.horsePower : ''}
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

                <input
                  ref={imageUpLoadRef}
                  name="fileUploader"
                  type="file"
                  hidden
                  onChange={onImageChange}
                />
                {selectedImage || car?.imageUrl ? (
                  <div className="mt-2 space-y-2">
                    <label className="bg-inherit text-xs font-medium focus-within:text-teal-50">
                      Thumbnail
                    </label>
                    {/* eslint-disable-next-line @next/next/no-img-element */}
                    <img
                      className="w-full cursor-pointer rounded-md object-contain"
                      src={
                        selectedImage ? selectedImage.toString() : car?.imageUrl
                      }
                      alt=""
                      onClick={() => imageUpLoadRef!.current?.click()}
                    />
                  </div>
                ) : (
                  <button
                    className="mx-auto flex h-10 items-center rounded-full p-3 text-sm outline-none hover:bg-neutral-600/10"
                    onClick={() => imageUpLoadRef!.current?.click()}
                    type="button"
                  >
                    <CameraIcon className="mr-1 h-6 w-6" />
                    Choose a thumbnail
                  </button>
                )}
                <div
                  className="my-2 text-center text-xs text-red-600"
                  dangerouslySetInnerHTML={{
                    __html: errors.imageUrl ? errors.imageUrl : '\u2000',
                  }}
                />
                <label className="my-2 bg-inherit text-xs font-medium focus-within:text-teal-50">
                  Quick review
                </label>
                <div className="w-full rounded-lg border border-gray-200 bg-gray-50">
                  <div className="rounded-t-lg bg-transparent px-4 py-2">
                    <label className="sr-only">Review</label>
                    <textarea
                      rows={4}
                      className="w-full border-0 bg-transparent px-0 text-sm text-gray-900 outline-none focus:ring-0"
                      placeholder="Write a comment..."
                      onChange={onChange}
                      value={car?.review}
                      name="review"
                      required
                    ></textarea>
                  </div>
                  <div className="flex items-center justify-center border-t px-3 py-2">
                    <button
                      className="inline-flex h-6 items-center px-4 text-center text-xs font-medium text-teal-600 hover:text-teal-700 disabled:text-neutral-400"
                      disabled={isGenerateingReview}
                      onClick={generateReview}
                      type="button"
                    >
                      Generate review
                      {isGenerateingReview && (
                        <div role="status">
                          <svg
                            aria-hidden="true"
                            className="ml-3 h-6 w-6 animate-spin fill-teal-600 text-gray-200"
                            viewBox="0 0 100 101"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                          >
                            <path
                              d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                              fill="currentColor"
                            />
                            <path
                              d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                              fill="currentFill"
                            />
                          </svg>
                          <span className="sr-only">Loading...</span>
                        </div>
                      )}
                    </button>
                  </div>
                </div>
                <div
                  className="my-2 text-center text-xs text-red-600"
                  dangerouslySetInnerHTML={{
                    __html: errors.review ? errors.review : '\u2000',
                  }}
                />
                <div className="mt-2 flex items-center justify-end">
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
                    className="flex h-10 items-center rounded-full bg-teal-600
                  px-5 py-2.5 text-sm font-medium text-white 
                  outline-none hover:bg-teal-700 focus:outline-none"
                    type="submit"
                  >
                    {type === 'create' ? 'Create Model' : 'Update'}
                    {isSubmitting && (
                      <div role="status">
                        <svg
                          aria-hidden="true"
                          className="ml-3 h-5 animate-spin fill-white text-gray-200"
                          viewBox="0 0 100 101"
                          fill="none"
                          xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                            d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                            fill="currentColor"
                          />
                          <path
                            d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                            fill="currentFill"
                          />
                        </svg>
                        <span className="sr-only">Loading...</span>
                      </div>
                    )}
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

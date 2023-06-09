import { CameraIcon } from '@heroicons/react/24/outline'
import { useRef } from 'react'
import useAddEditBrand from '../hooks/use-add-brand'
import OutLineInput from '../inputs/outline-input'
import OutLineOptionMenu from '../menu/OutlineOptionMenu'

const AddBrandForm = ({ hook }: { hook: ReturnType<typeof useAddEditBrand> }) => {
  const {
    brand,
    errors,
    formRef,
    selectedImage,
    addImage,
    resetState,
    onSubmit,
    onChange,
  } = hook
  const imageUpLoadRef = useRef<HTMLInputElement>(null)

  const currentYear = new Date().getFullYear()
  const years = Array.from(
    { length: currentYear - 1900 + 1 },
    (_, index) => currentYear - index
  )

  return (
    <form
      className="w-full max-w-sm mx-auto space-y-2 z-[3]"
      ref={formRef}
      onSubmit={onSubmit}
    >
      <OutLineInput
        label="Name"
        placeholder="brand Name"
        name="name"
        value={brand.name}
        onChange={onChange}
        errorMessage={errors.name}
        required
      />
      <OutLineInput
        label="Country"
        placeholder="country of origin"
        name="countryOfOrigin"
        value={brand.countryOfOrigin}
        onChange={onChange}
        errorMessage={errors.countryOfOrigin}
        required
      />
      <OutLineOptionMenu
        label="Year"
        name="foundedYear"
        options={years.map((y) => y.toString())}
        value={brand.foundedYear}
        onChange={onChange}
        errorMessage={errors.foundedYear}
        required
      />
      <OutLineInput
        placeholder="website URL"
        label="Website URL"
        name="websiteUrl"
        type="url"
        value={brand.websiteUrl}
        onChange={onChange}
        errorMessage={errors.websiteUrl}
        required
      />
      <input
        ref={imageUpLoadRef}
        name="fileUploader"
        type="file"
        hidden
        onChange={addImage}
      />
      {selectedImage ? (
        <div className="space-y-2">
          <label className="text-xs bg-inherit focus-within:text-teal-50 font-medium">
            Logo
          </label>
          {/* eslint-disable-next-line @next/next/no-img-element */}
          <img
            className="w-full object-contain rounded-md cursor-pointer"
            src={selectedImage.toString()}
            alt=""
            onClick={() => imageUpLoadRef!.current?.click()}
          />
        </div>
      ) : (
        <button
          className="flex items-center p-3 mx-auto hover:bg-neutral-600/10 rounded-full h-10 outline-none"
          onClick={() => imageUpLoadRef!.current?.click()}
          type="button"
        >
          <CameraIcon className="w-6 h-6 mr-1" />
          Chose a logo
        </button>
      )}
      <div
        className="text-xs text-red-600 text-center"
        dangerouslySetInnerHTML={{
          __html: errors.logoUrl ? errors.logoUrl : '\u2000',
        }}
      />
      <div className="flex items-center">
        <button
          className="ml-auto mr-3 rounded-full text-sm font-medium h-10 px-3 outline-none 
          text-red-600 enabled:hover:bg-red-600/10"
          type="button"
          onClick={() => resetState()}
        >
          Clear
        </button>
        <button
          className="focus:outline-none text-white bg-teal-600 hover:bg-teal-700 h-10 outline-none
          font-medium text-sm px-5 py-2.5 mb-2 rounded-full"
          type="submit"
        >
          Create Brand
        </button>
      </div>
    </form>
  )
}

export default AddBrandForm

import { CameraIcon } from '@heroicons/react/24/outline'
import { useRef } from 'react'
import useAddEditBrand from '../hooks/use-add-brand'
import OutLineInput from '../inputs/outline-input'
import OutLineOptionMenu from '../menu/OutlineOptionMenu'

const AddBrandForm = ({
  hook,
}: {
  hook: ReturnType<typeof useAddEditBrand>
}) => {
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
      className="z-[3] mx-auto w-full max-w-sm space-y-2"
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
          <label className="bg-inherit text-xs font-medium focus-within:text-teal-50">
            Logo
          </label>
          {/* eslint-disable-next-line @next/next/no-img-element */}
          <img
            className="w-full cursor-pointer rounded-md object-contain"
            src={selectedImage.toString()}
            alt=""
            onClick={() => imageUpLoadRef!.current?.click()}
          />
        </div>
      ) : (
        <button
          className="mx-auto flex h-10 items-center rounded-full p-3 outline-none hover:bg-neutral-600/10"
          onClick={() => imageUpLoadRef!.current?.click()}
          type="button"
        >
          <CameraIcon className="mr-1 h-6 w-6" />
          Chose a logo
        </button>
      )}
      <div
        className="text-center text-xs text-red-600"
        dangerouslySetInnerHTML={{
          __html: errors.logoUrl ? errors.logoUrl : '\u2000',
        }}
      />
      <div className="flex items-center">
        <button
          className="ml-auto mr-3 h-10 rounded-full px-3 text-sm font-medium text-red-600 
          outline-none enabled:hover:bg-red-600/10"
          type="button"
          onClick={() => resetState()}
        >
          Clear
        </button>
        <button
          className="mb-2 h-10 rounded-full bg-teal-600 px-5 py-2.5
          text-sm font-medium text-white outline-none hover:bg-teal-700 focus:outline-none"
          type="submit"
        >
          Create Brand
        </button>
      </div>
    </form>
  )
}

export default AddBrandForm

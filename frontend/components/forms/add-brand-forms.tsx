'use client'

import useAddBrand from '../hooks/use-add-brand'
import OutLineInput from '../inputs/outline-input'
import OutLineOptionMenu from '../menu/OutlineOptionMenu'

const AddBrandForm = ({ hook }: { hook: ReturnType<typeof useAddBrand> }) => {
  const { brand, errors, formRef, resetState, onSubmit, onChange } = hook

  const currentYear = new Date().getFullYear()
  const years = Array.from(
    { length: currentYear - 1900 + 1 },
    (_, index) => currentYear - index
  )

  return (
    <form className="w-full max-w-sm mx-auto" ref={formRef} onSubmit={onSubmit}>
      <OutLineInput
        label="Name"
        placeholder="brand Name"
        name="name"
        value={brand.name}
        onChange={onChange}
        errorMessage={errors.name}
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
      <OutLineInput
        label="Logo URL"
        name="logoUrl"
        type="url"
        value={brand.logoUrl}
        onChange={onChange}
        required
        errorMessage={errors.logoUrl}
      />
      <div className="flex items-center justify-between">
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          type="submit"
        >
          Create Brand
        </button>
        <button
          className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          type="button"
          onClick={() => resetState()}
        >
          Reset
        </button>
      </div>
    </form>
  )
}

export default AddBrandForm

import { HTMLProps } from 'react'

interface ExtendsInputProps extends HTMLProps<HTMLInputElement> {
  name?: string
  label?: string
  errorMessage?: string
}

const OutLineInput = ({
  name,
  label,
  errorMessage,
  ...props
}: ExtendsInputProps) => {
  return (
    <div className="min-h-min w-full space-y-2">
      <label className="bg-inherit text-xs font-medium focus-within:text-teal-50">
        {label}
      </label>
      <input
        className="h-10 w-full rounded-md bg-transparent indent-4 text-base outline
        outline-1 outline-neutral-200 placeholder:text-neutral-700
        focus:bg-neutral-100 focus:outline-2 focus:outline-teal-500"
        name={name}
        {...props}
      />
      <div
        className="text-xs text-red-600"
        dangerouslySetInnerHTML={{
          __html: errorMessage ? errorMessage : '\u2000',
        }}
      />
    </div>
  )
}

export default OutLineInput

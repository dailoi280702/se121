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
    <div className="w-full space-y-2 min-h-min">
      <label className="text-xs bg-inherit focus-within:text-teal-50 font-medium">
        {label}
      </label>
      <input
        className="w-full h-10 rounded-md indent-4 outline outline-1 text-base
        bg-transparent outline-neutral-200 placeholder-neutral-700
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

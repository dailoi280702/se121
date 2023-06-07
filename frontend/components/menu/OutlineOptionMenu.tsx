import React, { useState } from 'react'
import { HTMLProps } from 'react'

interface OutLineOptionMenuProps extends HTMLProps<HTMLSelectElement> {
  name?: string
  label?: string
  errorMessage?: string
  options: string[]
}

const OutLineOptionMenu = ({
  name,
  label,
  errorMessage,
  options,
  ...props
}: OutLineOptionMenuProps) => {
  const [selectedOption, setSelectedOption] = useState('')

  const handleOptionChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedOption(event.target.value)
    if (props.onChange) {
      props.onChange(event)
    }
  }

  return (
    <div className="w-full space-y-2 min-h-min">
      <label className="text-xs bg-inherit focus-within:text-teal-50 font-medium">
        {label}
      </label>
      <select
        className="w-full h-10 rounded-md indent-3 text-base bg-transparent 
        outline outline-1 outline-neutral-200 placeholder-neutral-700 
        focus:bg-neutral-100 focus:outline-2 focus:outline-teal-500"
        name={name}
        value={selectedOption}
        onChange={handleOptionChange}
        {...props}
      >
        {options.map((option) => (
          <option key={option} value={option}>
            {option}
          </option>
        ))}
      </select>
      <div
        className="text-xs text-red-600"
        dangerouslySetInnerHTML={{
          __html: errorMessage ? errorMessage : '\u2000',
        }}
      />
    </div>
  )
}

export default OutLineOptionMenu

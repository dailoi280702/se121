import { ChangeEvent, FormEvent, useState } from 'react'

export const useForm = <T>(callback: any, initialstate: T) => {
  const [values, setValues] = useState(initialstate)

  const onChange = <
    T extends HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
  >(
    e: ChangeEvent<T>
  ) => {
    setValues({ ...values, [e.target.name]: e.target.value })
  }

  const onSubmit = async (e?: FormEvent<HTMLFormElement>) => {
    if (e) {
      e.preventDefault()
    }
    await callback()
  }

  return {
    onChange,
    onSubmit,
    values,
    setValues,
  }
}

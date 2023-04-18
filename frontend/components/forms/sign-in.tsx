'use client'

import { useSetAtom } from 'jotai'
import { useRouter } from 'next/navigation'
import { ChangeEvent, FormEvent, HTMLInputTypeAttribute, useState } from 'react'
import ReactHtmlParser from 'react-html-parser'
import { UserAtom } from '../providers/user-provider'

const Input = ({
  name,
  label,
  errorMessage,
  placeHolder,
  required,
  type,
  onChange,
}: {
  name?: string
  label?: string
  errorMessage?: string
  placeHolder?: string
  required?: boolean
  type?: HTMLInputTypeAttribute
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void
}) => {
  return (
    <div className="w-full space-y-2 min-h-min">
      <h3 className="text-xs bg-inherit focus-within:text-teal-50 font-medium">
        {label}
      </h3>
      <input
        className="w-full h-10 rounded-md indent-4 outline outline-1 text-base
        bg-transparent outline-neutral-200 placeholder-neutral-700
        focus:bg-neutral-100 focus:outline-2 focus:outline-teal-500"
        name={name}
        placeholder={placeHolder ? placeHolder : 'required*'}
        required={required}
        type={type}
        onChange={onChange}
      />
      <h4 className="text-xs text-red-600">
        {errorMessage ? errorMessage : '\u2000'}
      </h4>
    </div>
  )
}

const useForm = <T,>(callback: any, initialstate: T) => {
  const [values, setValues] = useState(initialstate)

  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [e.target.name]: e.target.value })
  }

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    await callback()
  }

  return {
    onChange,
    onSubmit,
    values,
    setValues,
  }
}

interface Props {
  callbackUrl?: string
}

export default function SignInForm({ callbackUrl = '/' }: Props) {
  const router = useRouter()
  const setUser = useSetAtom(UserAtom)

  const initialstate = {
    nameOrEmail: '',
    password: '',
    messages: [],
    details: {
      nameOrEmail: '',
      password: '',
    },
  }

  const signInCallBack = async () => {
    console.log(values.nameOrEmail)
    console.log(values.password)

    const formData = {
      nameOrEmail: values.nameOrEmail.trim(),
      password: values.password.trim(),
    }

    const response = await fetch('http://localhost:8000/v1/auth', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
      credentials: 'include',
    })

    if (!response.ok) {
      const jsonData = await response.json()
      setValues({ ...values, ...jsonData })
      return
    }

    const user = await response.json()
    setUser(user)
    router.push(callbackUrl)
  }

  const { onChange, onSubmit, values, setValues } = useForm(
    signInCallBack,
    initialstate
  )

  return (
    <form
      className="flex flex-col px-6 space-y-2 text-left text-sm sm:mb-14"
      onSubmit={onSubmit}
    >
      <ul className="">
        {values.messages.map((message, idx) => (
          <h4 className="text-xs text-red-600" key={idx}>
            {message ? message : '\u2000'}
          </h4>
        ))}
      </ul>
      <Input
        name="nameOrEmail"
        label="name or email"
        errorMessage={values.details.nameOrEmail}
        onChange={onChange}
      />
      <Input
        name="password"
        label="password"
        errorMessage={values.details.password}
        onChange={onChange}
        type="password"
      />
      <button
        className="w-full h-10 rounded-md bg-teal-600 text-teal-50 !mt-8 font-medium"
        type="submit"
      >
        Sign in
      </button>
    </form>
  )
}

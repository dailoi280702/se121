'use client'

import { useSetAtom } from 'jotai'
import { useRouter } from 'next/navigation'
import { ChangeEvent, FormEvent, HTMLInputTypeAttribute, useState } from 'react'
import { User, UserAtom } from '../providers/user-provider'

export const Input = ({
  name,
  label,
  errorMessage,
  placeHolder,
  required,
  type,
  onChange,
  onBlur,
}: {
  name?: string
  label?: string
  errorMessage?: string
  placeHolder?: string
  required?: boolean
  type?: HTMLInputTypeAttribute
  onBlur?: () => void
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void
}) => {
  return (
    <div className="min-h-min w-full space-y-2">
      <h3 className="bg-inherit text-xs font-medium focus-within:text-teal-50">
        {label}
      </h3>
      <input
        className="h-10 w-full rounded-md bg-transparent indent-4 text-base outline
        outline-1 outline-neutral-200 placeholder:text-neutral-700
        focus:bg-neutral-100 focus:outline-2 focus:outline-teal-500"
        name={name}
        placeholder={placeHolder ? placeHolder : 'required*'}
        required={required}
        type={type}
        onChange={onChange}
        onBlur={() => {
          if (onBlur) onBlur()
        }}
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

export const useForm = <T,>(callback: any, initialstate: T) => {
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

    const user = (await response.json()) as User
    if (user) {
      user.createAt = new Date(user.createAt)
    }
    setUser(user)
    router.push(callbackUrl)
  }

  const { onChange, onSubmit, values, setValues } = useForm(
    signInCallBack,
    initialstate
  )

  return (
    <form
      className="flex flex-col space-y-2 px-6 text-left text-sm sm:mb-14"
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
        className="!mt-8 h-10 w-full rounded-md bg-teal-600 font-medium text-teal-50"
        type="submit"
      >
        Sign in
      </button>
    </form>
  )
}

'use client'
import { useRouter } from 'next/navigation'
import { useForm, Input } from './sign-in'

const USER_REGEX = /^[A-z][A-z0-9-_]{3,23}$/
const PWD_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/
const REGISTER_URL = '/register'

interface Props {
  callbackUrl?: string
}

export default function RegisterForm({ callbackUrl = '/' }: Props) {
  const router = useRouter()

  const validInputFields = () => {
    if (values.name == '') return false

    // check if there are no error message in value.details
    return Object.values(values.details).every(
      (errorMessage) => errorMessage == ''
    )
  }

  const signupCallback = async () => {
    if (!validInputFields()) {
      return
    }

    const formData = {
      name: values.name.trim(),
      email: values.email.trim(),
      password: values.password.trim(),
      rePassword: values.rePassword.trim(),
    }

    const response = await fetch('http://localhost:8000/v1/auth', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    })

    if (!response.ok) {
      const jsonData = await response.json()
      setValues({ ...values, ...jsonData })
      return
    }

    router.push(callbackUrl)
  }

  const initialstate = {
    name: '',
    email: '',
    password: '',
    rePassword: '',
    message: [],
    details: {
      name: '',
      email: '',
      password: '',
      rePassword: '',
    },
  }

  const { onChange, onSubmit, values, setValues } = useForm(
    signupCallback,
    initialstate
  )

  const onFocusOut = (field: string, value: string) => {
    const actions = new Map([
      [
        'name',
        {
          conditions: [
            {
              value: /^[A-z][A-z0-9-_]{3,23}$/,
              message:
                '4 to 24 characters. <br />Must begin with a letter.<br />Letters, numbers, underscores, hyphens allowed.',
            },
          ],
        },
      ],
      [
        'email',
        {
          conditions: [
            {
              value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
              message: 'Email invalid',
            },
          ],
        },
      ],
      [
        'password',
        {
          conditions: [
            {
              value: /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/,
              message:
                '8 to 24 characters.<br />Must include uppercase and lowercase letters, a number and a special character.<br />Allowed special characters.',
            },
          ],
        },
      ],
      [
        'rePassword',
        {
          conditions: [
            {
              value: values.password == values.rePassword,
              message: 'Password does not match',
            },
          ],
        },
      ],
    ])

    let message: string | undefined = undefined

    if (actions.has(field)) {
      actions.get(field)?.conditions.forEach((condition) => {
        if (typeof condition.value === 'boolean') {
          if (!condition.value) {
            message = condition.message
          }
        } else {
          if (!condition.value.test(value)) {
            message = condition.message
          }
        }
      })
    }

    setValues({
      ...values,
      details: {
        ...values.details,
        [field]: message || '',
      },
    })
  }

  const inputFields = [
    {
      name: 'name',
      label: 'Enter your username',
      placeHolder: 'Username ...',
      errorMessage: values.details.name,
      value: values.name,
    },
    {
      name: 'email',
      label: 'Enter your email',
      placeHolder: 'Email ...',
      errorMessage: values.details.email,
      value: values.email,
    },
    {
      name: 'password',
      label: 'Enter your password',
      placeHolder: 'Password ...',
      errorMessage: values.details.password,
      value: values.password,
      type: 'password',
    },
    {
      name: 'rePassword',
      label: 'Reenter your password',
      placeHolder: 'Reenter password ...',
      errorMessage: values.details.rePassword,
      value: values.rePassword,
      type: 'password',
    },
  ]

  return (
    <form
      className="flex flex-col space-y-2 px-6 text-left text-sm sm:mb-14"
      onSubmit={onSubmit}
    >
      {inputFields.map((input) => (
        <Input
          key={input.name}
          name={input.name}
          label={input.label}
          placeHolder={input.placeHolder}
          errorMessage={input.errorMessage}
          onChange={onChange}
          onBlur={() => onFocusOut(input.name, input.value)}
          type={input.type}
        />
      ))}
      <button
        className="!mt-8 h-10 w-full rounded-md bg-teal-600 font-medium text-teal-50"
        type="submit"
      >
        Sign up
      </button>
    </form>
  )
}

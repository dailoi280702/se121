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
              value: '',
              message:
                '4 to 24 characters. <br />Must begin with a letter.<br />Letters, numbers, underscores, hyphens allowed.',
            },
            {
              value: 'khoa',
              message: 'khoacc',
            },
          ],
        },
      ],
      [
        'email',
        {
          conditions: [
            {
              value: '',
              message: 'cc',
            },
          ],
        },
      ],
    ])

    var message = undefined
    if (actions.get(field)) {
      actions.get(field)?.conditions.forEach((condition) => {
        if (value == condition.value) {
          message = condition.message
        }
      })
    }
    setValues({
      ...values,
      details: {
        ...values.details,
        [field]: message ? message : '',
      },
    })
  }

  return (
    <form
      className="flex flex-col px-6 space-y-2 text-left text-sm sm:mb-14"
      onSubmit={onSubmit}
    >
      <Input
        name="name"
        label="Enter your username"
        placeHolder="Username ..."
        errorMessage={values.details.name}
        onChange={onChange}
        onBlur={() => onFocusOut('name', values.name)}
      />
      <Input
        name="email"
        label="Enter your email"
        placeHolder="Email ..."
        errorMessage={values.details.email}
        onChange={onChange}
        onBlur={() => onFocusOut('email', values.email)}
      />
      <Input
        name="password"
        label="Enter your password"
        placeHolder="Password ..."
        errorMessage={values.details.password}
        onChange={onChange}
        onBlur={() => onFocusOut('password', values.email)}
      />
      <Input
        name="rePassword"
        label="Reenter your password"
        placeHolder="Reenter password ..."
        errorMessage={values.details.rePassword}
        onChange={onChange}
        onBlur={() => onFocusOut('rePassword', values.email)}
      />
      <button
        className="w-full h-10 rounded-md bg-teal-600 text-teal-50 !mt-8 font-medium"
        type="submit"
      >
        Sign up
      </button>
    </form>
  )
}

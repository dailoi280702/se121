import { usePathname, useRouter } from 'next/navigation'
import { useRef, useState } from 'react'
import { useForm } from './use-form'

export default function useAddUpdateCar({
  initData,
  onSuccess,
  type,
}: {
  initData?: Car
  onSuccess?: () => void
  type: 'create' | 'update'
}) {
  const currentYear = new Date().getFullYear()
  const _initData = {
    id: 0,
    brand: undefined,
    series: undefined,
    name: '',
    year: currentYear,
    horsePower: undefined,
    torque: undefined,
    transmission: undefined,
    fuelType: undefined,
    imageUrl: undefined,
    review: undefined,
  } as Car

  const initError = {
    id: '',
    brand: '',
    series: '',
    name: '',
    year: '',
    horsePower: '',
    torque: '',
    transmission: '',
    fuelType: '',
    imageUrl: '',
    review: '',
  }

  const [errors, setErrors] = useState(initError)
  const router = useRouter()
  const path = usePathname()
  const formRef = useRef<HTMLFormElement>(null)

  const resetState = () => {
    setValues(initData ? initData : _initData)
    setErrors(initError)
  }

  const validate = () => {
    return true
  }

  const handleFailure = async (response: Response) => {
    if (response.status === 400 || response.status === 403) {
      const contentType = response.headers.get('content-type')
      if (contentType && contentType.indexOf('application/json') !== -1) {
        const data = await response.json()
        if (data.details) {
          setErrors(data.details)
        }
        return
      }
    }
    console.log(await response.text())
  }

  const handleResponse = async (response: Response) => {
    if (!response.ok) {
      await handleFailure(response)
      return
    }

    if (onSuccess) {
      onSuccess()
    }
    resetState()
    router.replace(path)
  }

  const update = async () => {
    if (!validate() && !initData) return

    try {
      const data = {}
      return

      const response = await fetch('http://localhost:8000/v1/car', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      await handleResponse(response)
    } catch (err) {
      console.log(err)
    }
  }

  const add = async () => {
    if (!validate()) return

    try {
      const data = {}
      return

      const response = await fetch('http://localhost:8000/v1/car', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      await handleResponse(response)
    } catch (err) {
      console.log(err)
    }
  }

  const {
    onChange,
    onSubmit,
    values: car,
    setValues,
  } = useForm(type === 'create' ? add : update, initData ? initData : _initData)

  return {
    car,
    errors,
    formRef,
    resetState,
    onChange,
    onSubmit,
  }
}

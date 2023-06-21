import { useForm } from '@/components/hooks/use-form'
import { usePathname, useRouter } from 'next/navigation'
import { useRef, useState } from 'react'

type actionType = 'update' | 'create'
export default function useAddUpdateSeries({
  initData,
  onSuccess,
  type,
}: {
  initData?: SeriesDetail
  onSuccess?: () => void
  type: actionType
}) {
  const [errors, setErrors] = useState({ name: '' })
  const router = useRouter()
  const path = usePathname()
  const formRef = useRef<HTMLFormElement>(null)

  const resetState = () => {
    setValues({
      name: initData ? initData.name : '',
    })
    setErrors({ name: '' })
  }

  const validate = () => {
    if (name.trim() === '') {
      setErrors({ name: 'Name cannot be emtpy' })
      return false
    }
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
    if (!validate()) return
    // :Todo
  }

  const add = async () => {
    if (!validate()) return

    try {
      const data = {
        brandId: initData?.brand.id,
        name: name.trim(),
      }

      const response = await fetch('http.localhotst://series/', {
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
    values: { name },
    setValues,
  } = useForm(type === 'create' ? add : update, {
    name: initData ? initData.name : '',
  })

  return {
    name,
    errors,
    formRef,
    resetState,
    onChange,
    onSubmit,
  }
}

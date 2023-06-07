'use client'

import { useRouter } from 'next/navigation'
import { useRef, useState } from 'react'
import { useForm } from './use-form'

export default function useAddBrand() {
  const router = useRouter()
  const formRef = useRef<HTMLFormElement>(null)

  const initBrand = {
    name: '',
    countryOfOrigin: '',
    foundedYear: new Date().getFullYear(),
    websiteUrl: '',
    logoUrl: '',
  } as Brand

  const initErrors = {
    name: '',
    countryOfOrigin: '',
    foundedYear: '',
    websiteUrl: '',
    logoUrl: '',
  }
  const [errors, setErrors] = useState(initErrors)

  const resetState = () => {
    setValues(initBrand)
  }

  const validate = () => {
    setErrors((pre) => ({
      ...pre,
      name: brand.name.trim() === '' ? 'Brand name can not be empty' : '',
      countryOfOrigin:
        brand.countryOfOrigin!.trim() === ''
          ? 'countryOfOrigin name can not be empty'
          : '',
      websiteUrl:
        brand.websiteUrl!.trim() === ''
          ? 'countryOfOrigin name can not be empty'
          : '',
    }))

    for (const [_, v] of Object.entries(initErrors)) {
      if (v) {
        return false
      }
    }
    return true
  }

  const submit = async () => {
    if (validate()) {
      // :TODO upload image

      const formData = {
        name: brand.name.trim(),
        // email: values.email.trim(),
        // password: values.password.trim(),
        // rePassword: values.rePassword.trim(),
      }

      // const response = await fetch('http://localhost:8000/v1/brand', {
      //   body: JSON.stringify(formData),
      // })
      //
      // if (!response.ok) {
      //   const data = await response.json()
      //   if (data.details) {
      //     setValues((values) => ({ ...values, errors: data.details }))
      //   }
      //   return
      // }

      router.refresh()
    }
  }

  const {
    values: brand,
    setValues,
    onChange,
    onSubmit,
  } = useForm(submit, initBrand)

  return {
    brand,
    errors,
    onSubmit,
    onChange,
    resetState,
    formRef,
  }
}

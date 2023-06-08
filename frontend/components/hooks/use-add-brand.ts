'use client'

import { usePathname, useRouter } from 'next/navigation'
import { useRef, useState } from 'react'
import { useForm } from './use-form'
import { v4 as uuidv4 } from 'uuid'
import { getDownloadURL, ref, uploadString } from 'firebase/storage'
import { storage } from '@/firebase/config'

export default function useAddBrand() {
  const router = useRouter()
  const path = usePathname()
  const formRef = useRef<HTMLFormElement>(null)
  const [selectedImage, setSelectedImage] = useState<
    string | null | ArrayBuffer
  >()

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
    setSelectedImage(null)
  }

  const validate = () => {
    setErrors((pre) => ({
      ...pre,
      name: brand.name.trim() === '' ? 'Brand name can not be empty' : '',
      countryOfOrigin:
        brand.countryOfOrigin!.trim() === '' ? 'Country can not be empty' : '',
      websiteUrl:
        brand.websiteUrl!.trim() === '' ? 'Website URL can not be empty' : '',
      logoUrl: !selectedImage ? 'Please chose a logo' : '',
    }))

    for (const [_, v] of Object.entries(initErrors)) {
      if (v) {
        return false
      }
    }
    return true
  }

  const submit = async () => {
    // Validte form data
    if (validate()) {
      // Upload image
      const imgID = uuidv4()
      const imageRef = ref(storage, `brand/${imgID}/image`)
      await uploadString(imageRef, selectedImage!.toString(), 'data_url')
        .then(async (_) => {
          // Retrive image URL
          const downloadURL = await getDownloadURL(imageRef)

          // Preprare data
          const data = {
            name: brand.name.trim(),
            countryOfOrigin: brand.countryOfOrigin?.trim(),
            foundedYear: Number(brand.foundedYear!),
            websiteUrl: brand.websiteUrl?.trim(),
            logoUrl: downloadURL,
          }

          // Update image URL
          const response = await fetch('http://localhost:8000/v1/brand', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
          })

          // Handle failure
          if (response.status === 400 || response.status === 403) {
            const data = await response.json()
            console.log(data.details)
            if (data.details) {
              setErrors(data.details)
            }
            console.log(errors)
            return
          } else if (!response.ok) {
            window.alert(data)
            return
          }

          // Refresh page, reset state
          resetState()
          router.replace(path)
        })
        .catch((err) => console.log(err))
    }
  }

  const addImage = (e: React.ChangeEvent<HTMLInputElement>) => {
    const reader = new FileReader()
    if (e.target.files && e.target.files[0]) {
      reader.readAsDataURL(e.target.files[0])
    }
    reader.onload = (readerEvent) => {
      if (readerEvent.target) {
        setSelectedImage(readerEvent.target.result)
      }
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
    selectedImage,
    addImage,
  }
}

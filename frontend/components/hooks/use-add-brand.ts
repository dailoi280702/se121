'use client'

import { usePathname, useRouter } from 'next/navigation'
import { useRef, useState } from 'react'
import { useForm } from './use-form'
import { v4 as uuidv4 } from 'uuid'
import { getDownloadURL, ref, uploadString } from 'firebase/storage'
import { storage } from '@/firebase/config'

type actionType = 'update' | 'create'

export default function useAddEditBrand({
  initData,
  onSuccess,
  type,
}: {
  initData?: Brand
  onSuccess?: () => void
  type: actionType
}) {
  const router = useRouter()
  const path = usePathname()
  const formRef = useRef<HTMLFormElement>(null)
  const [selectedImage, setSelectedImage] = useState<
    string | null | ArrayBuffer
  >()

  const emptyInitData = {
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
    setValues(initData ? initData : emptyInitData)
    setSelectedImage(null)
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

  const create = async () => {
    if (validate()) {
      const data = {
        name: brand.name.trim(),
        countryOfOrigin: brand.countryOfOrigin?.trim(),
        foundedYear: Number(brand.foundedYear!),
        websiteUrl: brand.websiteUrl?.trim(),
      }

      // Post data
      const response = await fetch('http://localhost:8000/v1/brand', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      if (!response.ok) {
        await handleFailure(response)
        return
      }

      // Retrive id
      const { id } = await response.json()
      const imgID = uuidv4()
      const imageRef = ref(storage, `brand/${imgID}/image`)

      // Upload image
      await uploadString(imageRef, selectedImage!.toString(), 'data_url')
        .then(async (_) => {
          // Retrive image URL
          const imageUrl = await getDownloadURL(imageRef)

          // Update image URL
          const response = await fetch(`http://localhost:8000/v1/brand/`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ id: id as Number, logoUrl: imageUrl }),
          })

          if (!response.ok) {
            await handleFailure(response)
            return
          }

          if (onSuccess) {
            onSuccess()
          }
          resetState()
          router.replace(path)
        })
        .catch((err) => console.log(err))
    }
  }

  const update = async () => {
    // Update brand
    if (onSuccess) {
      onSuccess()
    }
  }

  const {
    values: brand,
    setValues,
    onChange,
    onSubmit,
  } = useForm(
    type === 'create' ? create : update,
    initData ? initData : emptyInitData
  )

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

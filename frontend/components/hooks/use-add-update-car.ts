import { storage } from '@/firebase/config'
import { objectToQuery } from '@/utils'
import { getDownloadURL, ref, uploadString } from 'firebase/storage'
import { usePathname, useRouter } from 'next/navigation'
import { ChangeEvent, useEffect, useRef, useState } from 'react'
import { v4 } from 'uuid'
import { useForm } from './use-form'

export default function useAddUpdateCar({
  initData,
  series,
  onSuccess,
  type,
}: {
  initData?: Car
  series?: Series
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
    ...initData,
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
  const [fuelTypes, setFuelTypes] = useState<FuelType[]>([])
  const [transmissions, setTransmissions] = useState<Transmission[]>([])
  const [isGenerateingReview, setIsGeneratingReview] = useState(false)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [selectedImage, setSelectedImage] = useState<
    string | null | ArrayBuffer
  >()

  useEffect(() => {
    const fetchMetadata = async () => {
      try {
        const fetchURL = 'http://localhost:8000/v1/car/index'
        const res = await fetch(fetchURL)
        if (!res.ok) {
          console.log(res.text())
          return
        }

        const contentType = res.headers.get('content-type')
        if (contentType && contentType.indexOf('application/json') !== -1) {
          const data = await res.json()
          if (data.fuelType) {
            setFuelTypes(data.fuelType)
          }
          if (data.transmission) {
            setTransmissions(data.transmission)
          }
          return
        }
      } catch (e) {
        console.log(e)
      }
    }

    fetchMetadata()
  }, [])

  const onFuelTypeChange = (e: ChangeEvent<HTMLSelectElement>) => {
    fuelTypes.forEach((f) => {
      if (f.name === e.target.value) {
        setValues((car) => {
          return { ...car, fuelType: f }
        })
      }
    })
  }

  const onTransmissionChange = (e: ChangeEvent<HTMLSelectElement>) => {
    transmissions.forEach((t) => {
      if (t.name === e.target.value) {
        setValues((car) => {
          return { ...car, transmission: t }
        })
      }
    })
  }

  const onImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
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

  const generateReview = async () => {
    if (!validate() && !_initData) return

    setIsGeneratingReview(true)

    try {
      const data = {
        name: car.name,
        brand: _initData?.brand?.name,
        series: series,
        horsePower: car.horsePower,
        torque: car.torque,
        transmissions: car.transmission?.name,
        fuelTypes: car.fuelType?.name,
      }

      const res = await fetch(
        `http://localhost:8000/v1/text-generate/car-review?${objectToQuery(
          data
        )}`
      )

      if (!res.ok) {
        console.log('generate review failed:: ', res.text())
        setErrors((prev) => {
          return {
            ...prev,
            review: 'something happend, please try again later :/',
          }
        })
        setIsGeneratingReview(false)
        return
      }

      const contentType = res.headers.get('content-type')
      if (contentType && contentType.indexOf('application/json') !== -1) {
        const data = await res.json()
        if (data.text) {
          setValues((prev) => {
            return {
              ...prev,
              review: data.text,
            }
          })
        }
      }
    } catch (e) {
      setErrors((prev) => {
        return {
          ...prev,
          review: 'something happend, please try again later :/',
        }
      })
      console.log('generate review failed: ', e)
    }
    setIsGeneratingReview(false)
  }

  const resetState = () => {
    setValues(_initData)
    setErrors(initError)
    setSelectedImage(null)
    setIsSubmitting(false)
    setIsGeneratingReview(false)
  }

  const validate = () => {
    setErrors((pre) => ({
      ...pre,
      name: car.name.trim() === '' ? 'Car name can not be empty' : '',
      review:
        !car.review || car.review.trim() === '' ? 'Please write a review' : '',
      imageUrl:
        !selectedImage && !car.imageUrl ? 'Please choose a thumbnail' : '',
    }))

    for (const [_, v] of Object.entries(errors)) {
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
      } else {
        console.log(await response.text())
      }
    }
    setIsSubmitting(false)
    setIsGeneratingReview(false)
  }

  const handleSuccess = () => {
    if (onSuccess) {
      onSuccess()
    }
    resetState()
    router.replace(path)
  }

  const update = async () => {
    if (!validate() || !initData) return

    setIsSubmitting(true)
    try {
      const data: { [key: string]: any } = {}

      for (const key of ['year', 'torque', 'horsePower']) {
        const val = Number(car[key])
        if (initData[key] !== Number(car[key]) && val > 0 && val % 1 === 0) {
          data[key] = val
        }
      }
      for (const key of ['name', 'review']) {
        if (initData[key] !== car[key]) {
          data[key] = car[key]
        }
      }

      if (car.fuelType?.id !== initData.fuelType?.id) {
        data['fuelTypeId'] = car.fuelType?.id
      }
      if (car.transmission?.id !== initData.transmission?.id) {
        data['transmissionId'] = car.transmission?.id
      }

      if (Object.keys(data).length === 0) {
        setIsSubmitting(false)
        return
      }

      data['id'] = car.id
      console.log(JSON.stringify(data))

      const response = await fetch('http://localhost:8000/v1/car', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      if (!response.ok) {
        await handleFailure(response)
        return
      }

      if (selectedImage) {
        // Retrive id
        const imgID = v4()
        const imageRef = ref(storage, `car/${imgID}/image`)

        // Upload image
        await uploadString(imageRef, selectedImage!.toString(), 'data_url')
          .then(async (_) => {
            // Retrive image URL
            const imageUrl = await getDownloadURL(imageRef)

            // Update image URL
            const response = await fetch(`http://localhost:8000/v1/car/`, {
              method: 'PUT',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({
                id: car.id as Number,
                imageUrl: imageUrl,
              }),
            })

            if (!response.ok) {
              await handleFailure(response)
              return
            }
          })
          .catch((err) => console.log('err while uploading car image: ', err))
      }
      handleSuccess()
    } catch (err) {
      console.log('err while adding car model: ', err)
    }
    setIsSubmitting(false)
  }

  const add = async () => {
    if (!validate()) return

    setIsSubmitting(true)
    try {
      const data = {
        name: car.name,
        brandId: Number(car.brand?.id),
        seriesId: Number(series?.id),
        year: Number(car.year),
        horsePower: car.horsePower ? Number(car.horsePower) : undefined,
        torque: car.torque ? Number(car.torque) : undefined,
        transmissionId: car.transmission
          ? Number(car.transmission?.id)
          : transmissions.length > 0
          ? transmissions[0].id
          : undefined,
        fuelTypeId: car.fuelType
          ? Number(car.fuelType?.id)
          : fuelTypes.length > 0
          ? fuelTypes[0].id
          : undefined,
        review: car.review,
      }

      const response = await fetch('http://localhost:8000/v1/car', {
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
      const imgID = v4()
      const imageRef = ref(storage, `car/${imgID}/image`)

      // Upload image
      await uploadString(imageRef, selectedImage!.toString(), 'data_url')
        .then(async (_) => {
          // Retrive image URL
          const imageUrl = await getDownloadURL(imageRef)

          // Update image URL
          const response = await fetch(`http://localhost:8000/v1/car/`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ id: id as Number, imageUrl: imageUrl }),
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
        .catch((err) => console.log('err while uploading car image: ', err))
    } catch (err) {
      console.log('err while adding car model: ', err)
    }
    setIsSubmitting(false)
  }

  const {
    onChange,
    onSubmit,
    values: car,
    setValues,
  } = useForm(type === 'create' ? add : update, _initData)

  return {
    car,
    fuelTypes: fuelTypes.map((f) => f.name),
    transmissions: transmissions.map((t) => t.name),
    errors,
    isGenerateingReview,
    isSubmitting,
    selectedImage,
    formRef,
    resetState,
    onChange,
    onFuelTypeChange,
    onTransmissionChange,
    onImageChange,
    generateReview,
    onSubmit,
  }
}

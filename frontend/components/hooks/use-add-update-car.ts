import { usePathname, useRouter } from 'next/navigation'
import {
  ChangeEvent,
  ChangeEventHandler,
  useEffect,
  useRef,
  useState,
} from 'react'
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
  const [fuelTypes, setFuelTypes] = useState<FuelType[]>([])
  const [transmissions, setTransmissions] = useState<Transmission[]>([])
  const [isGenerateingReview, setIsGeneratingReview] = useState(false)

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

  const generateReview = async () => {
    if (!validate() && !initData) return

    setIsGeneratingReview(true)

    // message GenerateReviewReq {
    //   string name = 1;
    //   optional string brand = 2;
    //   optional string series = 3;
    //   optional int32 horsePower = 4;
    //   optional int32 torque = 5;
    //   optional string transmission = 6;
    //   optional string fuelType = 7;
    // }

    try {
      const data = {
        name: car.name,
        brand: car.brand?.name,
        series: car.series?.name,
        horsePower: car.horsePower,
        torque: car.torque,
        transmissions: car.transmission?.name,
        fuelTypes: car.fuelType?.name,
      }

      // :TODO ??????
    } catch (e) {
      console.log(e)
    }
    setIsGeneratingReview(false)
  }

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
    fuelTypes: fuelTypes.map((f) => f.name),
    transmissions: transmissions.map((t) => t.name),
    errors,
    isGenerateingReview,
    formRef,
    resetState,
    onChange,
    onFuelTypeChange,
    onTransmissionChange,
    generateReview,
    onSubmit,
  }
}

'use client'

import {
  ArrowTopRightOnSquareIcon,
  ExclamationCircleIcon,
} from '@heroicons/react/24/outline'
import Image from 'next/image'
import { useRouter } from 'next/navigation'

const CarCard = ({ car }: { car: Car }) => {
  const { id, brand, series, name, imageUrl } = car

  const router = useRouter()

  const handleClick = () => {
    router.push(`/car/${id}`)
  }

  return (
    <div
      className="relative flex h-40 min-w-max cursor-pointer overflow-hidden rounded-lg bg-white shadow-lg"
      onClick={handleClick}
    >
      {imageUrl && (
        <Image
          className="h-40 w-full object-cover object-center"
          width={300}
          height={160}
          sizes=""
          src={imageUrl}
          alt={name}
        />
      )}
      <div className="absolute inset-x-0 bottom-0 w-full bg-gradient-to-t from-black to-black/0 px-4 py-2 pt-6 text-white drop-shadow-md">
        <h3 className="text-lg font-medium">{name}</h3>
        <p>
          {brand?.name} - {series?.name}
        </p>
      </div>
    </div>
  )
}

export const SmallCarCard = ({
  car,
  className = '',
}: {
  car: Car
  className?: string
}) => {
  const { id, name, imageUrl } = car

  const router = useRouter()

  const handleClick = () => {
    router.push(`/car/${id}`)
  }

  return (
    <div
      className={[
        'relative flex h-32 min-w-max cursor-pointer overflow-hidden rounded-lg bg-white shadow-lg',
        className,
      ].join(' ')}
      onClick={handleClick}
    >
      {imageUrl ? (
        <Image
          className="object-cover"
          height={200}
          width={240}
          src={imageUrl}
          alt={name}
        />
      ) : (
        <div className="flex w-60 items-center justify-center stroke-black/40">
          Detail
          <ArrowTopRightOnSquareIcon className="ml-2 h-6 w-6" />
        </div>
      )}
      <div className="absolute inset-x-0 bottom-0 w-full bg-gradient-to-t from-black to-black/0 px-4 py-2 pt-6 text-white drop-shadow-md">
        <h3 className="text-sm font-medium">{name}</h3>
      </div>
    </div>
  )
}

export default CarCard

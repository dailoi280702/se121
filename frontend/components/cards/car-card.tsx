'use client'

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

export default CarCard

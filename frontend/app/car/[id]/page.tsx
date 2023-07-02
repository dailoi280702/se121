'use server'

import Link from 'next/link'
import Image from 'next/image'
import UpdateCarFAB from './udpate-car-fab'
import { notFound } from 'next/navigation'
import { SmallCarCard } from '@/components/cards/car-card'

async function fetchCar(id: number): Promise<Car | undefined> {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/car/${id}`, {
      cache: 'no-cache',
    })
    if (!res.ok) return undefined
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

async function fetchRelatedCars(id: number): Promise<{ cars: Car[] }> {
  try {
    const res = await fetch(
      `http://api-gateway:8000/v1/car/${id}/related?limit=5`,
      {
        cache: 'no-cache',
      }
    )
    if (res.ok) {
      return res.json()
    }
  } catch (err) {
    console.log(err)
  }
  return { cars: [] }
}

export default async function page({ params }: { params: { id: number } }) {
  // const car: Car = await fetchCar(params.id)
  // const { relatedCars }: { relatedCars: Car[] } = await fetchCar(params.id)
  const [car, { cars: relatedCars }] = await Promise.all([
    fetchCar(params.id),
    fetchRelatedCars(params.id),
  ])

  if (!car) {
    notFound()
  }
  return <CarPage car={car} relatedCars={relatedCars} />
}

const CarPage = ({ car, relatedCars }: { car: Car; relatedCars: Car[] }) => {
  const {
    brand,
    series,
    name,
    year,
    horsePower,
    torque,
    transmission,
    fuelType,
    imageUrl,
    review,
  } = car

  return (
    <>
      <div className="container mx-auto py-8">
        <div className="flex flex-col md:flex-row">
          {imageUrl && (
            <div className="relative w-full overflow-hidden md:w-1/2 md:rounded-md">
              <Image
                src={imageUrl}
                alt={name}
                width={900}
                height={600}
                className="w-full object-contain"
              />
            </div>
          )}
          <div className="mt-4 px-4 md:mt-0 md:w-1/2">
            <h1 className="text-2xl font-bold">{name}</h1>
            <div className="mt-2 flex items-center">
              {brand && (
                <Link
                  className="text-gray-600 hover:text-teal-600"
                  href={`/brand/${brand.id}`}
                >
                  {brand.name}
                </Link>
              )}
              {series && <p className="ml-4 text-gray-600">{series.name}</p>}
            </div>
            <hr className="my-4" />
            <div>
              {year && <p className="text-gray-600">Year: {year}</p>}
              {horsePower && (
                <p className="text-gray-600">Horsepower: {horsePower}</p>
              )}
              {torque && <p className="text-gray-600">Torque: {torque}</p>}
              {transmission && (
                <div className="text-gray-600">
                  Transmission:{' '}
                  <Link
                    className="hover:text-teal-600"
                    href={`/car?search=${transmission.name}`}
                  >
                    {transmission?.name}
                  </Link>
                </div>
              )}
              {fuelType && (
                <div className="text-gray-600">
                  Fuel Type:{' '}
                  <Link
                    className="hover:text-teal-600"
                    href={`/car?search=${fuelType.name}`}
                  >
                    {fuelType?.name}
                  </Link>
                </div>
              )}
            </div>
          </div>
        </div>
        {review && (
          <>
            <div className="px-4">
              <hr className="my-4" />
              <h2 className="text-lg font-bold">Review</h2>
              <p className="mt-2 text-gray-600">{review}</p>
            </div>
          </>
        )}
        <RelatedCars relatedCars={relatedCars} />
      </div>
      <UpdateCarFAB car={car} />
    </>
  )
}

const RelatedCars = ({ relatedCars }: { relatedCars: Car[] }) => {
  return (
    <>
      {relatedCars && relatedCars.length > 0 && (
        <div className="mt-6 px-4">
          <h2 className="text-lg font-bold">Related models</h2>
          <ul className="mt-2 flex w-full space-x-2 overflow-x-auto">
            {relatedCars.map((c) => (
              <SmallCarCard key={c.id} car={c} />
            ))}
          </ul>
        </div>
      )}
    </>
  )
}

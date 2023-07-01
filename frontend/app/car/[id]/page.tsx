'use server'

import Link from 'next/link'
import Image from 'next/image'
import UpdateCarFAB from './udpate-car-fab'

async function fetchCar(id: number) {
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

export default async function page({ params }: { params: { id: number } }) {
  const car: Car = await fetchCar(params.id)
  // return <div>{JSON.stringify(car)}</div>
  return <CarPage car={car} />
}

const CarPage = ({ car }: { car: Car }) => {
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
            <div className="relative overflow-hidden md:w-1/2 md:rounded-md">
              <Image
                src={imageUrl}
                alt={name}
                fill
                className="w-full object-contain"
              />
            </div>
          )}
          <div className="mt-4 px-4 md:mt-0 md:w-1/2">
            <h1 className="text-2xl font-bold">{name}</h1>
            <div className="mt-2 flex items-center">
              {brand && (
                <Link className="text-gray-600" href={`/brand/${brand.id}`}>
                  {brand.name}
                </Link>
              )}
              {series && <p className="ml-4 text-gray-600">{series.name}</p>}
            </div>
            <hr />
            <div className="mt-4">
              {year && <p className="text-gray-600">Year: {year}</p>}
              {horsePower && (
                <p className="text-gray-600">Horsepower: {horsePower}</p>
              )}
              {torque && <p className="text-gray-600">Torque: {torque}</p>}
              {transmission && (
                <p className="text-gray-600">
                  Transmission: {transmission?.name}
                </p>
              )}
              {fuelType && (
                <p className="text-gray-600">Fuel Type: {fuelType?.name}</p>
              )}
            </div>
          </div>
        </div>
        {review && (
          <>
            <div className="px-4">
              <hr className="my-4" />
              <h2 className="text-lg font-bold">Review</h2>
              <p className="text-gray-600">{review}</p>
            </div>
          </>
        )}
      </div>
      <UpdateCarFAB car={car} />
    </>
  )
}

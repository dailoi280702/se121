'use server'

import PageProgressBar from '@/components/page-progress-bar'
import PageSearch from '@/components/page-search'
import { objectToQuery } from '@/utils'
import { ArrowTopRightOnSquareIcon } from '@heroicons/react/24/outline'
import Image from 'next/image'
import Link from 'next/link'

const SEARCH_LIMIT = 20

async function fetchCars(req: SearchReq) {
  try {
    const fetchURL =
      'http://api-gateway:8000/v1/car/search?' + objectToQuery(req)
    const res = await fetch(fetchURL, { cache: 'no-store' })
    if (!res.ok) {
      console.log(res.text())
      return
    }
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

async function Cars({ promise }: { promise: Promise<SearchCarRes> }) {
  const carList: SearchCarRes = await promise

  return (
    <>
      <ul className="my-4 flex flex-wrap justify-center">
        {carList && carList.cars ? (
          <>
            {carList.cars.map((car) => (
              <div className="w-96 p-2 sm:w-1/2 md:w-1/3" key={car.id}>
                <CarCard car={car} />
              </div>
            ))}
          </>
        ) : (
          <div>No Result Found</div>
        )}
      </ul>
      {carList && carList.total && carList.total > 0 && (
        <PageProgressBar total={Math.ceil(carList.total / SEARCH_LIMIT)} />
      )}
    </>
  )
}

const CarCard = ({ car }: { car: Car }) => {
  const { id, brand, series, name, imageUrl } = car

  return (
    <div className="relative h-full cursor-pointer overflow-hidden rounded-lg bg-white shadow-lg sm:w-full">
      {imageUrl ? (
        <Image
          className="object-cover"
          height={600}
          width={600}
          src={imageUrl}
          alt={name}
        />
      ) : (
        <div className="flex h-40 min-h-full items-center justify-center stroke-black/40 stroke-2 drop-shadow-lg">
          Detail
          <ArrowTopRightOnSquareIcon className="ml-2 h-5 w-5" />
        </div>
      )}
      <div className="absolute inset-x-0 bottom-0 w-full bg-gradient-to-t from-black to-black/0 px-4 py-2 pt-6 text-white drop-shadow-md">
        <Link href={`/car/${id}`}>
          <p className="h-4" />
          <h3 className="text-lg font-medium">{name}</h3>
          <p className="text-sm">
            {brand?.name} - {series?.name}
          </p>
        </Link>
      </div>
    </div>
  )
}

export default async function Page({
  searchParams,
}: {
  searchParams: SearchQuery
}) {
  const { search, orderby, page } = searchParams
  const searchRequest = {
    query: search ? decodeURIComponent(search) : '',
    orderby: orderby ? decodeURIComponent(orderby) : 'year',
    limit: SEARCH_LIMIT,
    startAt: page ? SEARCH_LIMIT * (page - 1) : 1,
  }
  const filterOptions = new Map([
    ['Date', 'date'],
    ['Name', 'name'],
    ['Year', 'year'],
    ['Horse Power', 'horsePower'],
    ['Torque', 'torque'],
  ])

  const cars = fetchCars(searchRequest)

  return (
    <div className="mx-auto p-4 sm:max-w-6xl">
      <h1 className="mx-4 mb-4 text-xl font-medium">All Car Models</h1>
      <PageSearch filterOptions={filterOptions} defaultOption={'Year'} />
      <Cars promise={cars} />
    </div>
  )
}

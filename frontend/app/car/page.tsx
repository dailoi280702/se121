'use server'

import PageProgressBar from '@/components/page-progress-bar'
import PageSearch from '@/components/page-search'
import { objectToQuery } from '@/utils'

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
      <ul className="flex flex-col items-center space-y-4">
        {carList && carList.cars ? (
          <>
            {carList.cars.map((car) => (
              <div className="text-ellipsis overflow-hidden" key={car.id}>
                {JSON.stringify(car)}
              </div>
            ))}
          </>
        ) : (
          <div>No Result Found</div>
        )}
      </ul>
      {carList && carList.total && carList.total > 0 && (
        <PageProgressBar total={carList.total} />
      )}
    </>
  )
}

const SEARCH_LIMIT = 20

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
    startAt: page ? SEARCH_LIMIT * (page - 1) + 1 : 1,
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
    <div className="mx-auto sm:max-w-6xl p-4">
      <PageSearch filterOptions={filterOptions} defaultOption={'Year'} />
      <Cars promise={cars} />
    </div>
  )
}

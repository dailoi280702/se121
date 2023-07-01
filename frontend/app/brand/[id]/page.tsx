'use server'

import { PlusIcon } from '@heroicons/react/24/outline'
import { notFound } from 'next/navigation'
import BrandDetail from './brand-detail'
import AddUpdateSeries from './create-series'
import SeriesList from './series-list'

async function fetchBrand(id: number) {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/brand/${id}`, {
      cache: 'no-cache',
    })
    if (!res.ok) return undefined
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

async function fetchBrandSeries(id: number) {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/series?brandId=${id}`, {
      cache: 'no-cache',
    })
    if (!res.ok) return undefined
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

async function fetchBrandCars(id: number) {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/car?brandID=${id}`, {
      cache: 'no-cache',
    })
    if (!res.ok) return undefined
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

export default async function Page({
  params: { id },
}: {
  params: { id: number }
}) {
  const brand: Brand = await fetchBrand(id)
  const { series }: { series: Series[] } = await fetchBrandSeries(id)
  const { cars }: { cars: Car[] } = await fetchBrandCars(id)

  if (!brand) {
    notFound()
  }

  return (
    <>
      <BrandDetail brand={brand} />
      {series && (
        <>
          <div
            className="mb-4 mt-8 flex items-center justify-between
            space-x-2 text-2xl"
          >
            {`${brand.name}'s Series`}
            <AddUpdateSeries brand={brand} type="create">
              <button
                className="flex h-10 items-center rounded-full px-3 text-sm
                font-medium text-teal-600 outline-none 
                enabled:hover:bg-teal-600/10"
              >
                <PlusIcon className="mr-2 h-5 w-5 stroke-2" />
                New Series
              </button>
            </AddUpdateSeries>
          </div>
          <SeriesList series={series} cars={cars} brand={brand} />
        </>
      )}
    </>
  )
}

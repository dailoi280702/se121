'use client'

import AdminOnlyWrapper from '@/components/admin-only-wrapper'
import AddUpdateCar from '@/components/forms/add-update-car'
import { PencilIcon } from '@heroicons/react/24/outline'
import Image from 'next/image'
import { useState } from 'react'
import AddUpdateSeries from './create-series'

export default function SeriesList({
  series,
  cars,
  brand,
}: {
  series: Series[]
  cars: Car[]
  brand: Brand
}) {
  const [modalVisibility, setModalVisibility] = useState({
    updateSeries: false,
    createCarModel: false,
  })
  const [seriesToBeUpdate, setseriesToBeUpdate] = useState<SeriesDetail | null>(
    null
  )

  const openModal = (
    series: Series,
    type: 'updateSeries' | 'createCarModel'
  ) => {
    if (document.activeElement instanceof HTMLElement) {
      document.activeElement.blur()
    }
    setseriesToBeUpdate({ ...series, brand: brand })

    setModalVisibility((prev) => {
      if (type === 'updateSeries') {
        return { ...prev, updateSeries: true }
      } else {
        return { ...prev, createCarModel: true }
      }
    })
  }

  const seriesMap: Map<number, Car[]> = cars
    ? cars.reduce((map, car) => {
        const seriesId = car.series!.id
        map.set(seriesId, map.get(seriesId) || [])
        map.get(seriesId).push(car)
        return map
      }, new Map())
    : new Map()

  return (
    <>
      <div className="flex flex-col">
        {series.length &&
          series.map((s) => (
            <div key={s.id}>
              <hr />
              <div className="group my-2 flex h-10 w-full items-center">
                <h3>{s.name}</h3>
                <SeriesMenu
                  onUpdateSeriesClick={() => openModal(s, 'updateSeries')}
                  onCreateCarModelClick={() => openModal(s, 'createCarModel')}
                />
              </div>
              {seriesMap.get(s.id) &&
                seriesMap.get(s.id)!.map((c) => <div key={c.id}>{c.name}</div>)}
            </div>
          ))}
      </div>
      {modalVisibility.updateSeries && seriesToBeUpdate && (
        <AddUpdateSeries
          type="update"
          brand={brand}
          series={seriesToBeUpdate}
          isOpen={modalVisibility.updateSeries}
          setIsOpen={(isOpen: boolean) =>
            setModalVisibility((prev) => {
              return { ...prev, updateSeries: isOpen }
            })
          }
        >
          <div hidden />
        </AddUpdateSeries>
      )}
      <AddUpdateCar
        type="create"
        brand={brand}
        series={
          { ...seriesToBeUpdate, brandId: seriesToBeUpdate?.brand.id } as Series
        }
        isOpen={modalVisibility.createCarModel}
        setIsOpen={(isOpen: boolean) =>
          setModalVisibility((prev) => {
            return { ...prev, createCarModel: isOpen }
          })
        }
      >
        <div hidden />
      </AddUpdateCar>
    </>
  )
}

const SeriesMenu = ({
  onUpdateSeriesClick,
  onCreateCarModelClick,
}: {
  onUpdateSeriesClick: () => void
  onCreateCarModelClick: () => void
}) => {
  return (
    <AdminOnlyWrapper>
      <button
        className="group relative ml-auto hidden focus-within:block
        group-hover:block"
      >
        <div
          className="ml-auto flex h-10 w-10 items-center 
          justify-center rounded-full font-medium text-teal-600
          hover:bg-teal-700 hover:bg-opacity-[0.08]"
        >
          <PencilIcon className="h-6 w-6" />
        </div>
        <div
          className="absolute right-0 top-full z-10 hidden w-max divide-y
          divide-gray-100 overflow-hidden rounded-lg bg-white text-left 
          shadow group-focus:block"
        >
          <ul className="text-sm text-gray-700">
            <li
              className="border-b p-2 hover:bg-teal-600/10"
              onClick={onUpdateSeriesClick}
            >
              Update series
            </li>
          </ul>
          <ul className="text-sm text-gray-700">
            <li
              className="border-b p-2 hover:bg-teal-600/10"
              onClick={onCreateCarModelClick}
            >
              Add car model
            </li>
          </ul>
        </div>
      </button>
    </AdminOnlyWrapper>
  )
}

const CarCard = ({ car }: { car: Car }) => {
  const { name, brand, series, imageUrl } = car
  return (
    <div className="relative rounded-md">
      {imageUrl && (
        <Image
          className="absolute z-[1]"
          src={imageUrl!}
          fill
          sizes=""
          alt={name}
        />
      )}
    </div>
  )
}

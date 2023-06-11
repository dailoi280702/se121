'use server'

import { notFound } from 'next/navigation'
import BrandDetail from './brand-detail'

async function fetchBrand(id: number) {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/brand/${id}`)
    if (!res.ok) return undefined
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

async function fetchBrandSeries(id: number) {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/series?brandId=${id}`)
    if (!res.ok) return undefined
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

async function fetchBrandCars(id: number) {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/car?brandID=${id}`)
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
  const series: Series[] = await fetchBrandSeries(id)
  const cars: Car[] = await fetchBrandCars(id)

  if (!brand) {
    notFound()
  }

  return (
    <>
      <BrandDetail brand={brand} />
      Series
    </>
  )
}

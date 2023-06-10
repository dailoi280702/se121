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

export default async function Page({
  params: { id },
}: {
  params: { id: number }
}) {
  const brand: Brand = await fetchBrand(id)

  if (!brand) {
    notFound()
  }

  return (
    <>
      <div>{JSON.stringify(brand)}</div>
      <BrandDetail brand={brand} />
    </>
  )
}

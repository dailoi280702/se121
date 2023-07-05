'use server'

import AddBrandFromWithFab from '@/components/forms/add-brand-form-with-fab'
import PageProgressBar from '@/components/page-progress-bar'
import PageSearch from '@/components/page-search'
import { objectToQuery } from '@/utils'
import Image from 'next/image'
import BrandCardOverlay from './brand-card-overlay'

const SEARCH_LIMIT = 16

export async function fetchBrands(req: SearchReq) {
  try {
    const fetchURL =
      'http://api-gateway:8000/v1/brand/search?' + objectToQuery(req)
    const res = await fetch(fetchURL, { cache: 'no-cache' })
    if (!res.ok) {
      console.log(res.text())
      return
    }
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

const BrandComponent = ({ brand }: { brand: Brand }) => {
  const { id, name, countryOfOrigin, foundedYear, logoUrl } = brand

  return (
    <div className="w-1/2 p-2 sm:w-1/3 md:w-1/4 lg:w-1/5">
      <div
        className="group relative mx-auto h-full max-w-[300px] overflow-hidden 
        rounded-md bg-white p-4 shadow-lg"
      >
        <h2 className="mb-4 text-2xl font-bold">{name}</h2>
        {logoUrl && (
          <Image
            className="mx-auto mb-4 w-auto"
            src={logoUrl}
            alt={`${name} logo`}
            width={200}
            height={200}
          />
        )}
        {countryOfOrigin && (
          <p className="mb-2 text-gray-600">Origin: {countryOfOrigin}</p>
        )}
        {foundedYear && (
          <p className="mb-2 text-gray-600">Founded year: {foundedYear}</p>
        )}
        <BrandCardOverlay id={id} />
      </div>
    </div>
  )
}

async function Brands({ promise }: { promise: Promise<SearchBrandRes> }) {
  const brands: SearchBrandRes = await promise

  return (
    <>
      {/* <ul className="flex flex-col items-center space-y-4"> */}
      <ul className="my-4 flex flex-wrap justify-center">
        {brands && brands.brands ? (
          <>
            {brands.brands.map((brand) => (
              // <div key={brand.id}>{JSON.stringify(brand)}</div>
              <BrandComponent key={brand.id} brand={brand} />
            ))}
          </>
        ) : (
          <div>No Result Found</div>
        )}
      </ul>
      {brands && brands.total && brands.total > 0 && (
        <PageProgressBar total={Math.ceil(brands.total / SEARCH_LIMIT)} />
      )}
    </>
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
    ['Country', 'country'],
  ])

  const brands = fetchBrands(searchRequest)

  return (
    <>
      <h1 className="m-4 text-xl font-medium">All Brands</h1>
      <PageSearch filterOptions={filterOptions} defaultOption={'Year'} />
      <Brands promise={brands} />
      <AddBrandFromWithFab />
    </>
  )
}

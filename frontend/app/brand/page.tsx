'use server'

import AddBrandFromWithFab from '@/components/forms/add-brand-form-with-fab'
import PageProgressBar from '@/components/page-progress-bar'
import PageSearch from '@/components/page-search'
import { objectToQuery } from '@/utils'
import Image from 'next/image'
// import dynamic from 'next/dynamic'
// const AddBrandFromWithFab = dynamic(
//   () => import('@/components/forms/add-brand-form-with-fab')
// )

const SEARCH_LIMIT = 16

async function fetchBrands(req: SearchReq) {
  try {
    const fetchURL =
      'http://api-gateway:8000/v1/brand/search?' + objectToQuery(req)
    const res = await fetch(fetchURL)
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
  const { name, countryOfOrigin, foundedYear, logoUrl } = brand

  return (
    <div className="w-full sm:w-1/2 md:w-1/3 lg:w-1/4 p-2">
      <div className="p-4 bg-white shadow-lg rounded-md h-full max-w-[300px] mx-auto">
        <h2 className="text-2xl font-bold mb-4">{name}</h2>
        {logoUrl && (
          <Image
            className="mx-auto mb-4"
            src={logoUrl}
            alt={`${name} logo`}
            width={200}
            height={200}
          />
        )}
        {countryOfOrigin && (
          <p className="text-gray-600 mb-2">Origin: {countryOfOrigin}</p>
        )}
        {foundedYear && (
          <p className="text-gray-600 mb-2">Founded year: {foundedYear}</p>
        )}
        {/*
          {websiteUrl && (
            <p className="text-blue-500 mb-2">
              <a href={websiteUrl}>Official site</a>
            </p>
          )}
        */}
      </div>
    </div>
  )
}

async function Brands({ promise }: { promise: Promise<SearchBrandRes> }) {
  const brands: SearchBrandRes = await promise

  return (
    <>
      {/* <ul className="flex flex-col items-center space-y-4"> */}
      <ul className="flex flex-wrap justify-center my-4">
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
    <div className="mx-auto sm:max-w-6xl py-8 px-4 h-full">
      <PageSearch filterOptions={filterOptions} defaultOption={'Year'} />
      <Brands promise={brands} />
      <AddBrandFromWithFab />
    </div>
  )
}

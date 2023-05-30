'use server'

import PageProgressBar from '@/components/page-progress-bar'
import PageSearch from '@/components/page-search'
import { objectToQuery } from '@/utils'

const SEARCH_LIMIT = 20

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

async function Brands({ promise }: { promise: Promise<SearchBrandRes> }) {
  const brands: SearchBrandRes = await promise

  return (
    <>
      <ul className="flex flex-col items-center space-y-4">
        {brands && brands.brands ? (
          <>
            {brands.brands.map((brand) => (
              <div key={brand.id}>{JSON.stringify(brand)}</div>
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
    <div className="mx-auto sm:max-w-6xl p-4">
      <PageSearch filterOptions={filterOptions} defaultOption={'Year'} />
      <Brands promise={brands} />
    </div>
  )
}

'use client'

import { MagnifyingGlassIcon } from '@heroicons/react/24/outline'
import { usePathname, useRouter, useSearchParams } from 'next/navigation'
import { objectToQuery } from '@/utils'
import { useEffect, useMemo, useState } from 'react'

export default function PageSearch({
  filterOptions,
  defaultOption,
}: {
  filterOptions: Map<string, string>
  defaultOption: string
}) {
  const [query, setQuery] = useState<SearchQuery>()
  const [option, setOption] = useState<string>('')
  const options = useMemo(() => new Map(filterOptions), [filterOptions])
  const pathName = usePathname()
  const router = useRouter()
  const params = useSearchParams()

  useEffect(() => {
    const search = params.get('search')
    const orderby = params.get('orderby')
    const page = params.get('page')

    setQuery({
      page: parseInt(page ? page : '1'),
      orderby: orderby ? orderby : undefined,
      search: search ? search : undefined,
    })

    const getDefaultOption = (op: string): string => {
      let res = defaultOption
      options.forEach((v, k) => {
        if (op === v) {
          res = k
          return
        }
      })
      return res
    }

    setOption(getDefaultOption(orderby ? orderby : defaultOption))
  }, [params, defaultOption, options])

  const submit = (query: SearchQuery) => {
    if (!query) return

    const url = pathName + '?' + objectToQuery(query)
    router.push(url)
  }

  const onInputChange = (text: string) => {
    setQuery({ ...query, search: !!text ? text : undefined })
  }

  return (
    <form
      className="bg-neutral-100 focus-within:bg-neutral-200 transition-colors
        h-10 flex items-center rounded-md overflow-hidden px-2"
      onSubmit={(e) => {
        if (query) {
          e.preventDefault()
          submit(query)
        }
      }}
    >
      <MagnifyingGlassIcon
        className="h-5 w-5 ml-1 stroke-2 stroke-neutral-600"
        type="submit"
      />
      <input
        className="bg-transparent flex-grow outline-none mx-2 placeholder:text-neutral-700"
        placeholder="Search"
        onChange={(e) => onInputChange(e.target.value)}
        value={query ? query.search : ''}
      />
      <select
        className="bg-transparent outline-none text-right pr-2"
        value={option}
        onChange={(e) => {
          submit({ ...query, orderby: options.get(e.target.value) })
        }}
      >
        {Array.from(options).map(([v, k]) => (
          <option key={k} value={v} selected={v === option}>
            {v}
          </option>
        ))}
      </select>
    </form>
  )
}

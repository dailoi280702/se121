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
      className="mx-2 flex h-10
        items-center overflow-hidden rounded-md border bg-neutral-100 px-2 transition-colors focus-within:bg-neutral-200"
      onSubmit={(e) => {
        if (query) {
          e.preventDefault()
          submit(query)
        }
      }}
    >
      <MagnifyingGlassIcon
        className="ml-1 h-5 w-5 stroke-neutral-600 stroke-2"
        type="submit"
      />
      <input
        className="mx-2 grow bg-transparent outline-none placeholder:text-neutral-700"
        placeholder="Search"
        onChange={(e) => onInputChange(e.target.value)}
        value={query ? (query.search ? query.search : '') : ''}
      />
      <select
        className="bg-transparent pr-2 text-right outline-none"
        value={option}
        onChange={(e) => {
          submit({ ...query, orderby: options.get(e.target.value) })
        }}
      >
        {Array.from(options).map(([v, k]) => (
          <option
            key={k}
            value={v}
            // selected={v === option}
          >
            {v}
          </option>
        ))}
      </select>
    </form>
  )
}

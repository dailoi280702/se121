'use client'

import { objectToQuery } from '@/utils'
import { ArrowLeftIcon, ArrowRightIcon } from '@heroicons/react/24/outline'
import { usePathname, useRouter, useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react'

export default function PageProgressBar({ total: t }: { total: number }) {
  const [query, setQuery] = useState<SearchQuery>()
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
  }, [params])

  const submit = (page: number) => {
    const url =
      pathName +
      '?' +
      objectToQuery({
        ...query,
        page: page,
      } as SearchQuery)
    router.push(url)
  }

  return (
    <ul className="flex items-center justify-center space-x-2 overflow-x-scroll">
      <ProgressButton
        disabled={query ? query.page === 1 : false}
        onClick={() => {
          if (query && query.page && query.page > 1) {
            submit(query.page - 1)
          }
        }}
      >
        <ArrowLeftIcon className="w-5 h-5 mx-auto stroke-2" />
      </ProgressButton>
      {Array.from({ length: t }, (_, i) => i + 1).map((i) => (
        <ProgressButton
          key={i}
          text={i}
          isCurrent={query ? query.page === i : false}
          onClick={() => {
            submit(i)
          }}
        />
      ))}
      <ProgressButton
        disabled={query ? query.page === t : false}
        onClick={() => {
          if (query && query.page && query.page < t) {
            submit(query.page + 1)
          }
        }}
      >
        <ArrowRightIcon className="w-5 h-5 mx-auto stroke-2" />
      </ProgressButton>
    </ul>
  )
}

const ProgressButton = ({
  text,
  isCurrent,
  onClick,
  children,
  disabled,
}: {
  text?: any
  isCurrent?: boolean
  onClick: () => void
  children?: React.ReactNode
  disabled?: boolean
}) => {
  return (
    <button
      className={
        `${
          isCurrent
            ? 'bg-teal-600 hover:bg-teal-700 text-neutral-100 '
            : 'hover:bg-neutral-200 disabled:text-neutral-400 disabled:hover:bg-transparent '
        }` + 'h-10 w-10 min-w-[40] rounded-md font-medium transition-colors'
      }
      disabled={disabled}
      onClick={onClick}
    >
      {children ? children : text}
    </button>
  )
}

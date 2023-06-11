'use client'

import { ArrowUpRightIcon } from '@heroicons/react/24/outline'
import { useRouter } from 'next/navigation'

export default function BrandCardOverlay({ id }: { id: number }) {
  const router = useRouter()

  return (
    <div
      className="absolute inset-0 hidden items-center justify-center
          bg-transparent p-5 font-bold drop-shadow-lg backdrop-blur-md
          hover:block hover:bg-black/10 group-hover:flex"
      onClick={() => {
        router.push(`/brand/${id}`)
      }}
    >
      Read detail
      <ArrowUpRightIcon className="ml-2 h-4 w-4 stroke-[4]" />
    </div>
  )
}

'use client'

import AdminOnlyWrapper from '@/components/admin-only-wrapper'
import Fab from '@/components/buttons/fab'
import { PencilIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'

export default function UpdateBlogfab({ id }: { id: any }) {
  return (
    <AdminOnlyWrapper>
      <Link href={`/blog/${id}/edit`}>
        <Fab isFabOpen={false} setIsFabOpen={() => {}} icon={<PencilIcon />}>
          <></>
        </Fab>
      </Link>
    </AdminOnlyWrapper>
  )
}

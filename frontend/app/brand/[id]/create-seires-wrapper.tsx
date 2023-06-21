'use client'

import AdminOnlyWrapper from '@/components/admin-only-wrapper'
import CreateSeries from './create-series'

export default function CreateSeriesWrapper({ brand }: { brand: Brand }) {
  return (
    <AdminOnlyWrapper>
      <CreateSeries brand={brand} />
    </AdminOnlyWrapper>
  )
}

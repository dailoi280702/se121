'use server'

import AdminOnlyPage from '@/components/admin-only-page'
import AddUpdateBlog from '@/components/forms/add-update-blog'
import { notFound } from 'next/navigation'
import { fetchBlog } from '../page'

const fetchTags: () => Promise<{ tags: Tag[] }> = async () => {
  try {
    const res = await fetch('http://api-gateway:8000/v1/tag', {
      cache: 'no-cache',
    })
    if (!res.ok) {
      console.log(res.text())
      return { tags: [] }
    }
    return res.json()
  } catch (err) {
    console.log(err)
  }
  return { tags: [] }
}

export default async function Page({ params }: { params: { id: number } }) {
  const [blog, { tags }] = await Promise.all([
    fetchBlog(params.id),
    fetchTags(),
  ])

  if (!blog) {
    notFound()
  }

  return (
    <AdminOnlyPage>
      <div className="mx-auto h-full px-4 sm:max-w-3xl md:px-0">
        <h1 className="my-4 text-xl font-medium">Write New Blog</h1>
        <AddUpdateBlog type="update" tags={tags} blog={blog} />
      </div>
    </AdminOnlyPage>
  )
}

'use server'

import AdminOnlyPage from '@/components/admin-only-page'
import AddUpdateBlog from '@/components/forms/add-update-blog'

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

export default async function Page() {
  const { tags } = await fetchTags()

  return (
    <AdminOnlyPage>
      <div className="mx-auto h-full px-4 sm:max-w-3xl md:px-0">
        <h1 className="my-4 text-xl font-medium">Write New Blog</h1>
        <AddUpdateBlog type="create" tags={tags} />
      </div>
    </AdminOnlyPage>
  )
}

'use server'

import { notFound } from 'next/navigation'

export async function fetchBlog(id: number): Promise<Blog | undefined> {
  try {
    const res = await fetch(`http://api-gateway:8000/v1/blog/${id}`, {
      cache: 'no-cache',
    })
    if (!res.ok) return undefined
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

export default async function page({ params }: { params: { id: number } }) {
  const [blog] = await Promise.all([fetchBlog(params.id)])

  if (!blog) {
    notFound()
  }
  return <div>{JSON.stringify(blog)}</div>
}

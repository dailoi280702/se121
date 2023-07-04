'use server'

import { notFound } from 'next/navigation'
import Image from 'next/image'
import Tag from '@/components/tag'
import UpdateBlogfab from './update-blog-fab'
import Link from 'next/link'
import CommentSession from './comment-session'

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
  return (
    <>
      {blog.imageUrl && (
        <Image
          src={blog.imageUrl}
          width={1024}
          height={1024}
          className="w-full"
          alt="thumbnail"
        />
      )}
      <section className="my-4 px-4 md:px-0">
        <h1 className="text-2xl font-medium">{blog.title}</h1>
        <p className="text-xs">{new Date(blog.createdAt).toLocaleString()}</p>
        {blog.tags && (
          <ul className="flex flex-row-reverse flex-wrap items-center">
            {blog.tags.map((tag, index) => (
              <div className="cursor-pointer py-1 pr-2" key={index}>
                <Link href={`/blog?search=${tag.name}`}>
                  <Tag className="bg-white" name={tag.name} />
                </Link>
              </div>
            ))}
          </ul>
        )}
        <pre className="my-10 whitespace-pre-wrap font-sans text-base">
          {blog.body}
        </pre>
        {blog.tldr && (
          <>
            <h1 className="text-lg font-medium">TLDR;</h1>
            <p className="text-base">{blog.tldr}</p>
          </>
        )}
      </section>
      <CommentSession blogId={blog.id} />
      <UpdateBlogfab id={blog.id} />
    </>
  )
}

'use client'

import { UserAtom } from '@/components/providers/user-provider'
import RecommendedBlogs from '@/components/recommended-blogs'
import { ArrowRightIcon } from '@heroicons/react/24/outline'
import { useAtomValue } from 'jotai'
import Link from 'next/link'
import { useEffect, useState } from 'react'

export default function RecommendedBlogsSesssion() {
  const [blogs, setBlogs] = useState<Blog[]>([])
  const user = useAtomValue(UserAtom)

  useEffect(() => {
    if (user) {
      const fetchRecommendBlogs = async () => {
        try {
          const res = await fetch(
            `http://localhost:8000/v1/user/${user.id}/recommended-blogs?limit=10`
          )
          if (!res.ok) {
            console.log(
              'error while fetching user recommended-blogs: ',
              await res.text()
            )
          }

          const contentType = res.headers.get('content-type')
          if (contentType && contentType.indexOf('application/json') !== -1) {
            const data = await res.json()
            setBlogs(data.blogs)
          }
        } catch (e) {
          console.log('error while fetching user recommended-blogs: ', e)
        }
      }

      fetchRecommendBlogs()
    }
  }, [user])

  return (
    <>
      {user && blogs && blogs.length > 0 && (
        <section className="my-8">
          <div className="flex items-center">
            <h3 className="my-2 text-xl font-medium">
              Blogs Recommended For You
            </h3>
            <p className="ml-auto mr-2 text-sm font-normal">See more</p>
            <Link href={'/blog'}>
              <button
                className="flex h-10 w-10 items-center justify-center rounded-full font-bold
        hover:bg-neutral-700 hover:bg-opacity-[0.08]"
              >
                <ArrowRightIcon className="h-6 w-6 stroke-2" />
              </button>
            </Link>
          </div>
          <RecommendedBlogs blogs={blogs} />
        </section>
      )}
    </>
  )
}

'use server'

import PageProgressBar from '@/components/page-progress-bar'
import PageSearch from '@/components/page-search'
import { objectToQuery } from '@/utils'
import { PlusIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'

const SEARCH_LIMIT = 20

async function fetchBlogs(req: SearchReq) {
  try {
    const fetchURL =
      'http://api-gateway:8000/v1/blog/search?' + objectToQuery(req)
    const res = await fetch(fetchURL)
    if (!res.ok) {
      console.log(res.text())
      return
    }
    return res.json()
  } catch (err) {
    console.log(err)
  }
}

async function Blogs({ promise }: { promise: Promise<SearchBlogRes> }) {
  const blogs: SearchBlogRes = await promise

  return (
    <>
      <ul className="flex flex-col items-center space-y-4">
        {blogs && blogs.blogs ? (
          <>
            {blogs.blogs.map((blog) => (
              <div key={blog.id}>{JSON.stringify(blog)}</div>
            ))}
          </>
        ) : (
          <div>No Result Found</div>
        )}
      </ul>
      {blogs && blogs.total && blogs.total > 0 && (
        <PageProgressBar total={Math.ceil(blogs.total / SEARCH_LIMIT)} />
      )}
    </>
  )
}

export default async function Page({
  searchParams,
}: {
  searchParams: SearchQuery
}) {
  const { search, orderby, page } = searchParams
  const searchRequest = {
    query: search ? decodeURIComponent(search) : '',
    orderby: orderby ? decodeURIComponent(orderby) : 'date',
    limit: SEARCH_LIMIT,
    startAt: page ? SEARCH_LIMIT * (page - 1) : 1,
  }
  const filterOptions = new Map([
    ['Date', 'date'],
    ['Title', 'title'],
    ['Content', 'body'],
    ['Summarization', 'tldr'],
  ])

  const blogs = fetchBlogs(searchRequest)

  return (
    <div className="mx-auto p-4 sm:max-w-6xl">
      <div className="mb-4 flex w-full items-center justify-between px-2">
        <h1 className=" text-xl font-medium">All Blogs</h1>
        <Link href="/blog/new">
          <button
            className="flex h-10 items-center rounded-full px-3 text-sm
          font-medium text-teal-600 outline-none 
          enabled:hover:bg-teal-600/10"
          >
            <PlusIcon className="mr-2 h-5 w-5 stroke-2" />
            New Blog
          </button>
        </Link>
      </div>
      <PageSearch filterOptions={filterOptions} defaultOption={'Date'} />
      <Blogs promise={blogs} />
    </div>
  )
}

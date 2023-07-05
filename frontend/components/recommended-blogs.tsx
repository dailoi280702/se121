import Link from 'next/link'

export default function RecommendedBlogs({ blogs }: { blogs: Blog[] }) {
  return (
    <ul className="flex w-full flex-col items-center space-y-2">
      {blogs && (
        <>
          {blogs.map((blog) => (
            <>
              <Link
                href={`/blog/${blog.id}`}
                className={`w-full cursor-pointer`}
                key={blog.id}
              >
                <h2 className="font-medium">{blog.title}</h2>
                <p className="mb-2 text-xs">
                  {new Date(blog.createdAt).toLocaleString()}
                </p>
                <p>{blog.tldr ? blog.tldr : blog.body}</p>
                <hr className="my-4" />
              </Link>
            </>
          ))}
        </>
      )}
    </ul>
  )
}

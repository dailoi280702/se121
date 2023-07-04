import AddComment from '@/components/forms/add-comment'
import UserOnlyWarpper from '@/components/user-only-wrapper'
import Link from 'next/link'

export default function CommentSession({ blogId }: { blogId: Number }) {
  return (
    <section id="comments" className="my-10 px-4 md:px-0">
      <div className="flex items-center">
        <h3 className="text-xl font-medium">Conversation</h3>
        <p className="ml-auto text-sm">comments: {0}</p>
      </div>
      <hr className="my-2" />
      <UserOnlyWarpper
        fallbackElement={
          <Link
            className="hover:text-teal-600"
            href={`/auth/signin/${encodeURIComponent(`/blog/${blogId}`)}`}
          >
            Sign in to comment
          </Link>
        }
      >
        <AddComment blogId={blogId} />
      </UserOnlyWarpper>
    </section>
  )
}

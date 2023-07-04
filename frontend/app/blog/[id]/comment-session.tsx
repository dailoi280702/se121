import AddComment from '@/components/forms/add-comment'
import UserOnlyWarpper from '@/components/user-only-wrapper'
import Link from 'next/link'

export default function CommentSession({
  blogId,
  comments,
}: {
  blogId: Number
  comments: CommentDetail[]
}) {
  return (
    <section id="comments" className="my-10 px-4 md:px-0">
      <div className="flex items-center">
        <h3 className="text-xl font-medium">Conversation</h3>
        <p className="ml-auto text-sm">comments: {comments.length}</p>
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
      <CommentList comments={comments} />
    </section>
  )
}

const CommentList = ({ comments }: { comments: CommentDetail[] }) => {
  return (
    <>
      {comments &&
        comments.map((c) => (
          <div className="my-4" key={c.id}>
            <p className="font-medium">{c.user.name}</p>
            <p className="ml-auto text-xs">
              {new Date(c.createdAt).toLocaleString()}
            </p>
            <pre className="my-2 whitespace-pre-wrap font-sans text-base">
              {c.comment}
            </pre>
            <hr className="my-4" />
          </div>
        ))}
    </>
  )
}

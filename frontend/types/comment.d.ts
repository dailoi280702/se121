interface CommentDetail {
  id: number
  blogId: number
  user: UserProfile
  comment: string
  createdAt: number
  updatedAt?: number
}

interface UserProfile {
  id: string
  name: string
  imageUrl?: string
}

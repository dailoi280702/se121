type Blog = {
  id: number
  title: string
  body: string
  imageUrl?: string
  author: string
  tldr?: string
  createdAt: number
  updatedAt?: number
  tags: Tag[]
} & {
  [key: string]: any
}

type Tag = {
  id: number
  name: string
  description?: string
} & { [key: string]: any }

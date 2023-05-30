interface SearchReq {
  query?: string
  orderby?: string
  startAt?: number
  limit?: number
  isAscending?: boolean
}

interface SearchQuery {
  orderby?: string
  search?: string
  page?: number
}

interface SearchCarRes {
  cars: Car[]
  total: number
}

interface SearchBlogRes {
  cars: Blog[]
  total: number
}

interface SearchSeries {
  cars: Series[]
  total: number
}

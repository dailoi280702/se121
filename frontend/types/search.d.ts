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

interface SearchBrandRes {
  brands: Brand[]
  total: number
}

interface SearchBlogRes {
  blogs: Blog[]
  total: number
}

interface SearchSeries {
  series: Series[]
  total: number
}

interface SearchRes {
  cars?: SearchCarRes
  brands?: SearchBrandRes
  blogs?: SearchBlogRes
}

import { objectToQuery } from '@/utils'
import { useEffect, useState } from 'react'

const useGlobalSearch = () => {
  const [text, setText] = useState<string>('')
  const [data, setData] = useState<SearchRes | null>(null)

  const debounce = <F extends (...args: any[]) => void>(
    func: F,
    delay: number
  ) => {
    let timer: NodeJS.Timeout
    return (...args: Parameters<F>) => {
      clearTimeout(timer)
      timer = setTimeout(() => func(...args), delay)
    }
  }

  const search = debounce((s: string) => {
    setText(s)
  }, 200) // Debounce delay of 500 milliseconds

  useEffect(() => {
    let timeoutId: NodeJS.Timeout | null = null

    if (text.trim() !== '') {
      // Clear any existing timeout
      if (timeoutId) {
        clearTimeout(timeoutId)
      }

      // Start a new timeout to delay the API request
      timeoutId = setTimeout(() => {
        const fetchURL = `http://localhost:8000/v1/search?${objectToQuery({
          query: text,
          limit: 5,
        } as SearchReq)}`
        fetch(fetchURL)
          .then((res) => res.json())
          .then((data: SearchRes) => setData(data))
          .catch((error) => console.error('Error:', error))
      }, 300) // Delay of 0.3 second before making the API request
    }

    return () => {
      if (timeoutId) {
        clearTimeout(timeoutId)
      }
    }
  }, [text])

  return {
    search,
    data,
  }
}

export default useGlobalSearch

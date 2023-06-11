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
      timeoutId = setTimeout(async () => {
        try {
          const fetchURL = `http://localhost:8000/v1/search?${objectToQuery({
            query: text,
            limit: 5,
          } as SearchReq)}`
          const res = await fetch(fetchURL)

          if (res.ok) {
            const contentType = res.headers.get('content-type')
            if (contentType === 'application/json') {
              setData(await res.json())
            } else {
              console.log(await res.text())
            }
          }
        } catch (err) {
          console.log(err)
        }
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

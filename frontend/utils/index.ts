// export function objectToQuery(obj: Object) {
//   const params = new URLSearchParams()
//   Object.entries(obj).forEach(([key, value]) => {
//     if (Array.isArray(value)) {
//       value.forEach((item) => {
//         params.append(key, item.toString())
//       })
//     } else {
//       params.append(key, value.toString())
//     }
//   })
//   return params.toString()
// }

import { RefObject } from 'react'

export function objectToQuery(obj: Object) {
  const params = new URLSearchParams()

  Object.entries(obj).forEach(([key, value]) => {
    if (value !== null && value !== undefined && value !== '') {
      if (Array.isArray(value)) {
        value.forEach((item) => {
          if (item !== null && item !== undefined && item !== '') {
            params.append(key, item.toString())
          }
        })
      } else {
        params.append(key, value.toString())
      }
    }
  })

  return params.toString()
}

export function triggerFormUsingRef(ref?: RefObject<HTMLFormElement>) {
  if (ref && ref.current) {
    if (ref.current.reportValidity()) {
      const submitEvent = new Event('submit', {
        bubbles: true,
        cancelable: true,
      })
      ref.current.dispatchEvent(submitEvent)
    }
  }
}

export function getDomainName(url: string) {
  // Remove protocol (e.g., http:// or https://)
  let domain = url.replace(/(^\w+:|^)\/\//, '')

  // Remove path and query string
  domain = domain.split('/')[0]

  // Remove subdomains if present
  const parts = domain.split('.')
  if (parts.length > 2) {
    domain = parts.slice(1).join('.')
  }

  return domain
}

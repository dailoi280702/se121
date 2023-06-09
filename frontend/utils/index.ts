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

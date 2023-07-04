import { useAtomValue } from 'jotai'
import { useRouter } from 'next/navigation'
import { useRef, useState } from 'react'
import { UserAtom } from '../providers/user-provider'
import { useForm } from './use-form'

const initError = {
  comment: '',
}

const initData = {
  comment: '',
}

export default function useAddComment({ blogId }: { blogId: Number }) {
  const [errors, setErrors] = useState(initError)
  const router = useRouter()
  const formRef = useRef<HTMLFormElement>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const user = useAtomValue(UserAtom)

  const resetState = () => {
    setValues(initData)
    setErrors(initError)
    setIsSubmitting(false)
  }

  const validate = () => {
    setErrors((pre) => ({
      ...pre,
      comment: comment.comment.trim() === '' ? 'Comment can not be empty' : '',
    }))

    for (const [_, v] of Object.entries(errors)) {
      if (v) {
        return false
      }
    }
    return true
  }

  const handleFailure = async (response: Response) => {
    if (response.status === 400 || response.status === 403) {
      const contentType = response.headers.get('content-type')
      if (contentType && contentType.indexOf('application/json') !== -1) {
        const data = await response.json()
        if (data.details) {
          setErrors(data.details)
        }
      } else {
        console.log(await response.text())
      }
    }
    setIsSubmitting(false)
  }

  const add = async () => {
    if (isSubmitting || !user || !validate()) return

    setIsSubmitting(true)
    try {
      const data = {
        comment: comment.comment.trim(),
        blogId: Number(blogId),
        userId: user.id,
      }

      const response = await fetch('http://localhost:8000/v1/comment', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      if (!response.ok) {
        await handleFailure(response)
        return
      }

      resetState()
      router.push(`/blog/${blogId}#comments`)
    } catch (err) {
      console.log('err while adding blog: ', err)
    }
    setIsSubmitting(false)
  }

  const {
    onChange,
    onSubmit,
    values: comment,
    setValues,
  } = useForm(add, initData)

  return {
    states: {
      isSubmitting,
    },
    values: {
      comment,
      errors,
      formRef,
    },
    events: {
      resetState,
      onChange,
      onSubmit,
    },
  }
}

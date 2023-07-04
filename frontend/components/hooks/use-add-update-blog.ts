import { storage } from '@/firebase/config'
import { objectToQuery } from '@/utils'
import { getDownloadURL, ref, uploadString } from 'firebase/storage'
import { useAtomValue } from 'jotai'
import { usePathname, useRouter } from 'next/navigation'
import { useRef, useState } from 'react'
import { UserAtom } from '../providers/user-provider'
import { useForm } from './use-form'

const initError = {
  id: '',
  title: '',
  body: '',
  imageUrl: '',
  author: '',
  tldr: '',
  createdAt: '',
  updatedAt: '',
  tags: '',
}

export default function useAddUpdateBlog({
  initData,
  onSuccess,
  type,
  tags,
}: {
  initData: Blog
  tags: Tag[]
  onSuccess?: () => void
  type: 'create' | 'update'
}) {
  const [errors, setErrors] = useState(initError)
  const router = useRouter()
  const path = usePathname()
  const formRef = useRef<HTMLFormElement>(null)
  const [tagInput, setTagInput] = useState('')
  const [isGenerateingReview, setIsGeneratingReview] = useState(false)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [selectedImage, setSelectedImage] = useState<
    string | null | ArrayBuffer
  >()
  const [suggestedTags, setSuggestedTags] = useState<Tag[]>([])
  const user = useAtomValue(UserAtom)

  const onImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const reader = new FileReader()
    if (e.target.files && e.target.files[0]) {
      reader.readAsDataURL(e.target.files[0])
    }
    reader.onload = (readerEvent) => {
      if (readerEvent.target) {
        setSelectedImage(readerEvent.target.result)
      }
    }
  }

  const onTagInputChange = (value: string) => {
    const name = value.trim()
    if (name !== '') {
      if (tags && tags.length > 0) {
        setSuggestedTags(
          tags
            .filter(
              (tag) =>
                tag.name.includes(name) &&
                (!blog.tags || !blog.tags.includes(tag))
            )
            .slice(0, 10)
        )
      }
    } else {
      setSuggestedTags([])
    }
    setTagInput(value)
  }

  const addTag = (tag?: Tag) => {
    if (!tag) {
      const name = tagInput.trim()
      if (name !== '') {
        const isDuplicated =
          blog.tags && blog.tags.some((tags) => tags.name === name)
        if (!isDuplicated) {
          blog.tags.push({ id: 0, name: name })
          setValues((prev) => ({ ...prev, tags: blog.tags }))
        }
        setTagInput('')
      }
    } else {
      console.log('fukkkkk')
      blog.tags.push(tag)
      setValues((prev) => ({ ...prev, tags: blog.tags }))
      setTagInput('')
    }
  }

  const removeTag = (tag: Tag) => {
    const index = blog.tags.indexOf(tag)
    blog.tags.splice(index, 1)
    setValues((prev) => ({ ...prev, tags: blog.tags }))
  }

  const generateReview = async () => {
    if (!validate() && initData) return

    setIsGeneratingReview(true)

    try {
      const data = {
        title: blog.title,
        body: blog.body,
      }

      const res = await fetch(
        `http://localhost:8000/v1/text-generate/blog-summarize?${objectToQuery(
          data
        )}`
      )

      if (!res.ok) {
        console.log('generate review failed:: ', res.text())
        setErrors((prev) => {
          return {
            ...prev,
            review: 'something happend, please try again later :/',
          }
        })
        setIsGeneratingReview(false)
        return
      }

      const contentType = res.headers.get('content-type')
      if (contentType && contentType.indexOf('application/json') !== -1) {
        const data = await res.json()
        if (data.text) {
          setValues((prev) => {
            return {
              ...prev,
              review: data.text,
            }
          })
        }
      }
    } catch (e) {
      setErrors((prev) => {
        return {
          ...prev,
          review: 'something happend, please try again later :/',
        }
      })
      console.log('generate review failed: ', e)
    }
    setIsGeneratingReview(false)
  }

  const resetState = () => {
    setTagInput('')
    setValues(initData)
    setErrors(initError)
    setSelectedImage(null)
    setIsSubmitting(false)
    setIsGeneratingReview(false)
  }

  const validate = () => {
    setErrors((pre) => ({
      ...pre,
      title: blog.title.trim() === '' ? 'Title can not be empty' : '',
      body: blog.body.trim() === '' ? 'Body can not be empty' : '',
      review:
        !blog.tldr || blog.tldr.trim() === ''
          ? 'Please summarize your blog'
          : '',
      imageUrl:
        !selectedImage && !blog.imageUrl ? 'Please choose a thumbnail' : '',
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
    setIsGeneratingReview(false)
  }

  const handleSuccess = () => {
    if (onSuccess) {
      onSuccess()
    }
    resetState()
    router.replace(path)
  }

  const update = async () => {
    if (isSubmitting || !validate() || !initData) return

    setIsSubmitting(true)
    try {
      const data: { [key: string]: any } = {}

      for (const key of ['title', 'body', 'tldr']) {
        if (initData[key] !== blog[key]) {
          data[key] = blog[key]
        }
      }
      data['tags'] = blog.tags
      data['id'] = blog.id

      const response = await fetch('http://localhost:8000/v1/blog', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })

      if (!response.ok) {
        await handleFailure(response)
        return
      }

      if (selectedImage) {
        const imageRef = ref(storage, `blog/${blog.id}/image`)

        // Upload image
        await uploadString(imageRef, selectedImage!.toString(), 'data_url')
          .then(async (_) => {
            // Retrive image URL
            const imageUrl = await getDownloadURL(imageRef)

            // Update image URL
            const response = await fetch(`http://localhost:8000/v1/blog/`, {
              method: 'PUT',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({
                id: blog.id as Number,
                imageUrl: imageUrl,
                tags: blog.tags,
              }),
            })

            if (!response.ok) {
              await handleFailure(response)
              return
            }
          })
          .catch((err) => console.log('err while uploading blog image: ', err))
      }
      handleSuccess()
      router.push(`/blog/${blog.id}`)
    } catch (err) {
      console.log('err while updating blog: ', err)
    }
    setIsSubmitting(false)
  }

  const add = async () => {
    if (isSubmitting || !user || !validate()) return

    setIsSubmitting(true)
    try {
      const data = {
        title: blog.title,
        body: blog.body,
        tldr: blog.tldr,
        tags: blog.tags,
        author: user.id,
      }

      const response = await fetch('http://localhost:8000/v1/blog', {
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

      // Retrive id
      const { id } = await response.json()
      const imageRef = ref(storage, `blog/${id}/image`)

      // Upload image
      await uploadString(imageRef, selectedImage!.toString(), 'data_url')
        .then(async (_) => {
          // Retrive image URL
          const imageUrl = await getDownloadURL(imageRef)

          // Update image URL
          const response = await fetch(`http://localhost:8000/v1/blog/`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              id: id as Number,
              imageUrl: imageUrl,
              tags: blog.tags,
            }),
          })

          if (!response.ok) {
            await handleFailure(response)
            return
          }

          if (onSuccess) {
            onSuccess()
          }
          resetState()
          router.push(`/blog/${id}`)
        })
        .catch((err) => console.log('err while uploading blog image: ', err))
    } catch (err) {
      console.log('err while adding blog: ', err)
    }
    setIsSubmitting(false)
  }

  const {
    onChange,
    onSubmit,
    values: blog,
    setValues,
  } = useForm(type === 'create' ? add : update, initData)

  return {
    states: {
      isGenerateingReview,
      isSubmitting,
    },
    values: {
      blog,
      errors,
      tagInput,
      selectedImage,
      formRef,
      suggestedTags,
    },
    events: {
      resetState,
      onChange,
      onImageChange,
      onTagInputChange,
      addTag,
      removeTag,
      generateReview,
      onSubmit,
    },
  }
}

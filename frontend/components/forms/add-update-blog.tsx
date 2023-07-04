'use client'

import useAddUpdateBlog from '../hooks/use-add-update-blog'
import OutLineInput from '../inputs/outline-input'
import { useState, useRef } from 'react'
import Tag from '../tag'
import { CameraIcon } from '@heroicons/react/24/outline'

const defaultBlog = {
  id: 0,
  title: '',
  body: '',
  author: '',
  createdAt: 0,
  updatedAt: undefined,
  tldr: undefined,
  imageUrl: undefined,
  tags: [],
}

const TagsIput = ({
  inputValue,
  values,
  suggestion,
  errorMessage,
  onInputChange,
  addTag,
  removeTag,
}: {
  inputValue: string
  values: Tag[]
  suggestion: Tag[]
  errorMessage?: string
  onInputChange: (value: string) => void
  addTag: (tag?: Tag) => void
  removeTag: (tags: Tag) => void
}) => {
  const [isInputFocused, setInputFocused] = useState(false)
  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      event.preventDefault()
      addTag()
    } else if (event.key === 'Tab') {
      event.preventDefault()
      if (suggestion && suggestion.length > 0) {
        onInputChange(suggestion[0].name)
      }
    }
  }

  return (
    <div className="min-h-min w-full space-y-2">
      <label className="group bg-inherit text-xs font-medium focus-within:text-teal-50">
        Tags
      </label>
      <div className="relative">
        <input
          className="h-10 w-full rounded-md bg-transparent indent-4 text-base outline
          outline-1 outline-neutral-200 placeholder:text-neutral-700
          focus:bg-neutral-100 focus:outline-2 focus:outline-teal-500"
          value={inputValue}
          placeholder="tag"
          onChange={(e) => onInputChange(e.target.value)}
          onKeyDown={handleKeyDown}
          onFocus={() => setInputFocused(true)}
          onBlur={() => setInputFocused(false)}
        />
        {inputValue !== '' &&
          isInputFocused &&
          suggestion &&
          suggestion.length > 0 && (
            <ul
              className="absolute left-0 top-12 rounded-md 
          border bg-white py-2 shadow"
            >
              {suggestion.map((tag, index) => (
                <button
                  key={index}
                  className="block w-full px-2 text-left hover:bg-neutral-200"
                  onClick={() => addTag(tag)}
                >
                  {tag.name}
                </button>
              ))}
            </ul>
          )}
      </div>
      {values && values.length > 0 && (
        <ul className="flex flex-wrap items-center">
          {values.map((tag, index) => (
            <div className="py-1 pr-2" key={index}>
              <Tag name={tag.name} onIconClick={() => removeTag(tag)} />
            </div>
          ))}
        </ul>
      )}
      <div
        className="text-xs text-red-600"
        dangerouslySetInnerHTML={{
          __html: errorMessage ? errorMessage : '\u2000',
        }}
      />
    </div>
  )
}

export default function AddUpdateBlog({
  blog = defaultBlog,
  tags,
  type,
}: {
  blog?: Blog
  tags: Tag[]
  type: 'create' | 'update'
}) {
  const { values, states, events } = useAddUpdateBlog({
    type: type,
    tags: tags,
    initData: blog,
  })
  const imageUpLoadRef = useRef<HTMLInputElement>(null)

  return (
    <form
      className="space-y-2"
      onSubmit={(e) => {
        e.preventDefault()
        events.onSubmit()
      }}
    >
      {/* thumbnail */}
      <input
        ref={imageUpLoadRef}
        name="fileUploader"
        type="file"
        hidden
        onChange={events.onImageChange}
      />
      {values.selectedImage || values.blog.imageUrl ? (
        <div className="mt-2 space-y-2">
          <label className="bg-inherit text-xs font-medium focus-within:text-teal-50">
            Thumbnail
          </label>
          {/* eslint-disable-next-line @next/next/no-img-element */}
          <img
            className="w-full cursor-pointer rounded-md object-contain"
            src={
              values.selectedImage
                ? values.selectedImage.toString()
                : values.blog.imageUrl
            }
            alt=""
            onClick={() => imageUpLoadRef!.current?.click()}
          />
        </div>
      ) : (
        <button
          className="mx-auto flex h-10 items-center rounded-full p-3 text-sm outline-none hover:bg-neutral-600/10"
          onClick={() => imageUpLoadRef!.current?.click()}
          type="button"
        >
          <CameraIcon className="mr-1 h-6 w-6" />
          Choose a thumbnail
        </button>
      )}
      <div
        className="my-2 text-center text-xs text-red-600"
        dangerouslySetInnerHTML={{
          __html: values.errors.imageUrl ? values.errors.imageUrl : '\u2000',
        }}
      />
      {/* title */}
      <OutLineInput
        placeholder="title"
        label="Title"
        name="title"
        value={values.blog.title}
        onChange={events.onChange}
        errorMessage={values.errors.title}
        required
      />
      {/* tags */}
      <TagsIput
        suggestion={values.suggestedTags}
        inputValue={values.tagInput}
        values={values.blog.tags}
        errorMessage={values.errors.tags}
        onInputChange={events.onTagInputChange}
        addTag={events.addTag}
        removeTag={events.removeTag}
      />
      {/* content */}
      <div className="space-y-2">
        <label className="bg-inherit text-xs font-medium focus-within:text-teal-50">
          Content
        </label>
        <div className="w-full rounded-lg border border-gray-200 bg-transparent">
          <div className="rounded-t-lg bg-transparent px-4 py-2">
            <label className="sr-only">Content</label>
            <textarea
              rows={20}
              className="w-full border-0 bg-transparent px-0 text-sm text-neutral-800 outline-none focus:ring-0"
              placeholder="write something..."
              onChange={events.onChange}
              value={values.blog.body ? values.blog.body : ''}
              name="body"
              required
            ></textarea>
          </div>
        </div>
        <div
          className="text-center text-xs text-red-600"
          dangerouslySetInnerHTML={{
            __html: values.errors.body ? values.errors.body : '\u2000',
          }}
        />
      </div>
      {/* review */}
      <div className="space-y-2">
        <label className="bg-inherit text-xs font-medium focus-within:text-teal-50">
          Summarization
        </label>
        <div className="w-full rounded-lg border border-gray-200 bg-transparent">
          <div className="rounded-t-lg bg-transparent px-4 py-2">
            <label className="sr-only">Review</label>
            <textarea
              rows={4}
              className="w-full border-0 bg-transparent px-0 text-sm text-neutral-800 outline-none focus:ring-0"
              placeholder="summarize..."
              onChange={events.onChange}
              value={values.blog.tldr ? values.blog.tldr : ''}
              name="tldr"
              required
            ></textarea>
          </div>
          <div className="flex items-center justify-center border-t px-3 py-2">
            <button
              className="inline-flex h-6 items-center px-4 text-center text-xs font-medium text-teal-600 hover:text-teal-700 focus:outline-none disabled:text-neutral-400"
              disabled={states.isGenerateingReview}
              onClick={events.generateReview}
              type="button"
            >
              Generate summarization
              {states.isGenerateingReview && (
                <div role="status">
                  <svg
                    aria-hidden="true"
                    className="ml-3 h-6 w-6 animate-spin fill-teal-600 text-gray-200"
                    viewBox="0 0 100 101"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                      fill="currentColor"
                    />
                    <path
                      d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                      fill="currentFill"
                    />
                  </svg>
                  <span className="sr-only">Loading...</span>
                </div>
              )}
            </button>
          </div>
        </div>
        <div
          className="text-center text-xs text-red-600"
          dangerouslySetInnerHTML={{
            __html: values.errors.tldr ? values.errors.tldr : '\u2000',
          }}
        />
      </div>
      {/* submit buttons */}
      <div className="mt-2 flex items-center justify-end">
        <button
          className="ml-auto mr-3 h-10 rounded-full px-3 text-sm 
                  font-medium text-red-600 outline-none 
                  enabled:hover:bg-red-600/10"
          type="button"
          onClick={() => events.resetState()}
        >
          Reset
        </button>
        <button
          className="flex h-10 items-center rounded-full bg-teal-600
                  px-5 py-2.5 text-sm font-medium text-white 
                  outline-none hover:bg-teal-700 focus:outline-none"
          type="submit"
        >
          {type === 'create' ? 'Post' : 'Update'}
          {states.isSubmitting && (
            <div role="status">
              <svg
                aria-hidden="true"
                className="ml-3 h-5 animate-spin fill-white text-gray-200"
                viewBox="0 0 100 101"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                  fill="currentColor"
                />
                <path
                  d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                  fill="currentFill"
                />
              </svg>
              <span className="sr-only">Loading...</span>
            </div>
          )}
        </button>
      </div>
    </form>
  )
}

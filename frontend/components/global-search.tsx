import { MagnifyingGlassIcon } from '@heroicons/react/24/outline'
import { atom, useAtom } from 'jotai'
import { HTMLProps, ReactNode } from 'react'
import useCloseShade from './hooks/use-close-shade'
import useGlobalSearch from './hooks/use-global-search'
import { Shade } from './shade'

export const globalSearchAtom = atom<boolean>(false)

interface SearchULProps extends HTMLProps<HTMLUListElement> {
  children: ReactNode
}

const SearchUL = ({ children, ...props }: SearchULProps) => {
  return (
    <ul className="space-y-1 px-4" {...props}>
      {children}
    </ul>
  )
}

export default function GlobalSearch() {
  const { search, data } = useGlobalSearch()
  const [showGlobalSearch, setGlobalSearch] = useAtom(globalSearchAtom)
  const closeSearch = useCloseShade(globalSearchAtom, true)
  const openSearch = () => {
    if (!showGlobalSearch) {
      setGlobalSearch(true)
    }
  }

  return (
    <>
      {showGlobalSearch && <Shade onClose={closeSearch} className="z-[7]" />}
      <div
        className={
          `relative z-[7] flex h-9 items-center
          justify-center rounded-lg bg-neutral-100 px-2 ring-neutral-600
          transition-all focus-within:ring-2 sm:w-min ` +
          `${
            showGlobalSearch
              ? 'ml-auto flex-grow'
              : 'ml-auto rounded-full sm:rounded-lg'
          }`
        }
      >
        <input
          placeholder="Search"
          className={
            `${!showGlobalSearch ? 'hidden sm:inline-block' : ''}` +
            ` flex-grow bg-transparent outline-none placeholder:text-neutral-600`
          }
          onFocus={openSearch}
          onChange={(e) => search(e.target.value)}
        />
        <MagnifyingGlassIcon
          className="ml-1 h-5 w-5 stroke-2"
          onClick={openSearch}
        />
        {data && showGlobalSearch && (
          <div
            className="absolute  inset-x-0 top-12 space-y-4 rounded-md
            bg-neutral-100 p-4"
          >
            {data.blogs && data.blogs.blogs && (
              <section>
                Blogs
                <SearchUL>
                  {data.blogs.blogs.map((blog) => (
                    <div key={blog.id}>{blog.title}</div>
                  ))}
                </SearchUL>
              </section>
            )}
            {data?.cars && data.cars.cars && (
              <section>
                Cars
                <SearchUL>
                  {data.cars.cars.map((cars) => (
                    <div key={cars.id}>{cars.name}</div>
                  ))}
                </SearchUL>
              </section>
            )}
            {data?.brands && data.brands.brands && (
              <section>
                Brands
                <SearchUL>
                  {data.brands.brands.map((blog) => (
                    <div key={blog.id}>{blog.name}</div>
                  ))}
                </SearchUL>
              </section>
            )}
          </div>
        )}
      </div>
    </>
  )
}

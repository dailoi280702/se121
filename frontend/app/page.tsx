'use server'

import { ArrowRightIcon } from '@heroicons/react/24/outline'
import Link from 'next/link'
import { fetchCars } from './car/page'
import { fetchBlogs } from './blog/page'
import RecommendedBlogs from '@/components/recommended-blogs'
import CarCard from '@/components/cards/car-card'
import RecommendedBlogsSesssion from './blog/recommeded-blogs-session'

export default async function Home() {
  const [{ cars }, { blogs }] = await Promise.all([
    fetchCars({ orderby: 'date', limit: 10 }) as Promise<SearchCarRes>,
    fetchBlogs({ orderby: 'date', limit: 10 }) as Promise<SearchBlogRes>,
  ])

  return (
    <div className="mx-auto mb-24 h-full py-4 sm:max-w-5xl">
      <p className="text-lg">Hi, what are you interested in?</p>
      <section className="my-8">
        <div className="flex items-center">
          <h3 className="my-2 text-xl font-medium">
            Recently Added Car Models
          </h3>
          <p className="ml-auto mr-2 text-sm font-normal">See more</p>
          <Link href={'/car'}>
            <button
              className="flex h-10 w-10 items-center justify-center rounded-full font-bold
        hover:bg-neutral-700 hover:bg-opacity-[0.08]"
            >
              <ArrowRightIcon className="h-6 w-6 stroke-2" />
            </button>
          </Link>
        </div>
        {cars && (
          <ul className="flex overflow-x-auto">
            {cars.map((car) => (
              <div className="pr-2" key={car.id}>
                <CarCard car={car} />
              </div>
            ))}
          </ul>
        )}
      </section>
      <RecommendedBlogsSesssion />
      <section className="my-8">
        <div className="flex items-center">
          <h3 className="my-2 text-xl font-medium">Recently Added Blogs</h3>
          <p className="ml-auto mr-2 text-sm font-normal">See more</p>
          <Link href={'/blog'}>
            <button
              className="flex h-10 w-10 items-center justify-center rounded-full font-bold
        hover:bg-neutral-700 hover:bg-opacity-[0.08]"
            >
              <ArrowRightIcon className="h-6 w-6 stroke-2" />
            </button>
          </Link>
        </div>
        <RecommendedBlogs blogs={blogs} />
      </section>
    </div>
  )
}

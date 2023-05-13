import {
  ArrowDownCircleIcon,
  ArrowUpRightIcon,
} from '@heroicons/react/24/outline'
import Image from 'next/image'

const FetchHello = async () => {
  try {
    const response = await fetch('http://go-backend:8000/v1/say-hello')
    const data = await response.text()

    if (!response.ok) return '???/'

    return data
  } catch {
    return 'error'
  }
}

export default async function Home() {
  const helloString = await FetchHello()

  return (
    <>
      <div>
        <hr className="border-[#F7AB0A] mb-10" />
        <div className="grid grid-cols-1 md:grid-cols-2 px-10 gap-10 gap-y-16 pb-24">
          <div>
            <div className="flex flex-col group cursor-pointer">
              <div className="relative w-full h-80 drop-shadow-xl group-hover:scale-105 transition-transform duration-200 ease-out">
                <Image
                  className="object-cover object-left lg:object-center"
                  src="https://imageio.forbes.com/specials-images/imageserve/5d35eacaf1176b0008974b54/2020-Chevrolet-Corvette-Stingray/0x0.jpg?format=jpg&crop=4560,2565,x790,y784,safe&width=960"
                  alt="Image"
                  fill
                />
                <div className="absolute bottom-0 w-full bg-opacity-20 bg-black backdrop-blur-lg rounded drop-shadow-lg text-white p-5 flex justify-between">
                  <div>
                    <p className="font-bold">Title</p>
                    <p>Post time</p>
                  </div>
                  <div className="flex flex-col md:flex-row gap-y-2 md:gap-x-2 items-center">
                    <div className="bg-[#F7AB0A] text-center text-black px-3 py-1 rounded-full text-sm font-semibold">
                      <p className="font-bold text-center flex items-center">
                        Read paper
                        <ArrowUpRightIcon className="ml-1 h-2.5 w-2.5" />
                      </p>
                    </div>
                    <div className="bg-[#F7AB0A] text-center text-black px-3 py-1 rounded-full text-sm font-semibold">
                      <p className="font-bold flex justify-center items-center">
                        Report paper
                      </p>
                    </div>
                  </div>
                </div>
              </div>
              <div className="mt-5 flex-1">
                <p className="underline text-lg font-bold">Tilte</p>
                <p className="text-gray-500">Short description</p>
              </div>
            </div>
          </div>

          <div>
            <div className="flex flex-col group cursor-pointer">
              <div className="relative w-full h-80 drop-shadow-xl group-hover:scale-105 transition-transform duration-200 ease-out">
                <Image
                  className="object-cover object-left lg:object-center"
                  src="https://imageio.forbes.com/specials-images/imageserve/5d35eacaf1176b0008974b54/2020-Chevrolet-Corvette-Stingray/0x0.jpg?format=jpg&crop=4560,2565,x790,y784,safe&width=960"
                  alt="Image"
                  fill
                />
                <div className="absolute bottom-0 w-full bg-opacity-20 bg-black backdrop-blur-lg rounded drop-shadow-lg text-white p-5 flex justify-between">
                  <div>
                    <p className="font-bold">Title</p>
                    <p>Post time</p>
                  </div>
                  <div className="flex flex-col md:flex-row gap-y-2 md:gap-x-2 items-center">
                    <div className="bg-[#F7AB0A] text-center text-black px-3 py-1 rounded-full text-sm font-semibold">
                      <p className="font-bold text-center flex items-center">
                        Read paper
                        <ArrowUpRightIcon className="ml-1 h-2.5 w-2.5" />
                      </p>
                    </div>
                    <div className="bg-[#F7AB0A] text-center text-black px-3 py-1 rounded-full text-sm font-semibold">
                      <p className="font-bold flex justify-center items-center">
                        Report paper
                      </p>
                    </div>
                  </div>
                </div>
              </div>
              <div className="mt-5 flex-1">
                <p className="underline text-lg font-bold">Tilte</p>
                <p className="text-gray-500">Short description</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

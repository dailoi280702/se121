import Image from 'next/image'

export default function SeriesList({
  series,
  cars,
}: {
  series: Series[]
  cars: Car[]
}) {
  const seriesMap: Map<number, Car[]> = cars
    ? cars.reduce((map, car) => {
        const seriesId = car.series!.id
        map.set(seriesId, map.get(seriesId) || [])
        map.get(seriesId).push(car)
        return map
      }, new Map())
    : new Map()

  return (
    <div className="flex flex-col">
      {series.length &&
        series.map((s) => (
          <div key={s.id}>
            {s.name}
            {seriesMap.get(s.id) &&
              seriesMap.get(s.id)!.map((c) => <div key={c.id}>{c.name}</div>)}
            <hr />
          </div>
        ))}
    </div>
  )
}

// <div>
//   <button
//     className="flex h-10 w-10 items-center justify-center rounded-full font-bold
// hover:bg-neutral-700 hover:bg-opacity-[0.08]"
//     onClick={() => {}}
//   >
//     <PencilIcon className="h-6 w-6" />
//   </button>
//   <div
//     id="dropdown"
//     className="z-10 w-44 divide-y divide-gray-100 rounded-lg bg-white shadow"
//   >
//     <ul
//       className="py-2 text-sm text-gray-700"
//       aria-labelledby="dropdownDefaultButton"
//     >
//       <li>
//         <a href="#" className="block px-4 py-2 hover:bg-gray-100">
//           Dashboard
//         </a>
//       </li>
//       <li>
//         <a href="#" className="block px-4 py-2 hover:bg-gray-100">
//           Settings
//         </a>
//       </li>
//       <li>
//         <a href="#" className="block px-4 py-2 hover:bg-gray-100">
//           Earnings
//         </a>
//       </li>
//       <li>
//         <a href="#" className="block px-4 py-2 hover:bg-gray-100">
//           Sign out
//         </a>
//       </li>
//     </ul>
//   </div>
// </div>

const CarCard = ({ car }: { car: Car }) => {
  const { name, brand, series, imageUrl } = car
  return (
    <div className="relative rounded-md">
      {imageUrl && (
        <Image
          className="absolute z-[1]"
          src={imageUrl!}
          fill
          sizes=""
          alt={name}
        />
      )}
    </div>
  )
}

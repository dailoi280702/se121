import Image from 'next/image'

export default function BrandDetail({ brand }: { brand: Brand }) {
  return (
    <div>
      <div className="relative max-w-lg overflow-x-auto bg-white shadow-md sm:rounded-lg">
        <div className="flex items-center bg-white p-5 text-left text-lg font-semibold text-neutral-900">
          {brand.logoUrl && (
            <div className="relative mr-4 h-32 w-96">
              <Image
                className="object-contain"
                fill
                alt={`${brand.name} logo`}
                src={brand.logoUrl}
              />
            </div>
          )}
          <div>
            {brand.name}
            <p className="mt-1 text-sm font-normal text-neutral-500">
              Browse a list of Flowbite products designed to help you work and
              play, stay organized, get answers, keep in touch, grow your
              business, and more.
            </p>
          </div>
        </div>
        <hr />
        <table className="w-full text-left text-sm text-neutral-500">
          <tbody>
            {brand.countryOfOrigin && (
              <tr className="border-b bg-white">
                <th
                  scope="row"
                  className="whitespace-nowrap px-6 py-4 font-medium text-neutral-900"
                >
                  Origin
                </th>
                <td className="px-6 py-4">{brand.countryOfOrigin}</td>
              </tr>
            )}
            {brand.foundedYear && (
              <tr className="border-b bg-white">
                <th
                  scope="row"
                  className="whitespace-nowrap px-6 py-4 font-medium text-neutral-900"
                >
                  Founded year
                </th>
                <td className="px-6 py-4">White</td>
              </tr>
            )}
            {brand.websiteUrl && (
              <tr className="bg-white">
                <th
                  scope="row"
                  className="whitespace-nowrap px-6 py-4 font-medium text-neutral-900"
                >
                  Official site
                </th>
                <td className="px-6 py-4">
                  <a
                    href={brand.websiteUrl}
                    className="font-medium text-blue-600 hover:underline"
                  >
                    {brand.websiteUrl.replace('https://www.', '')}
                  </a>
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  )
}

import Image from 'next/image'

export default function BrandDetail({ brand }: { brand: Brand }) {
  return (
    <div>
      <div className="relative overflow-x-auto shadow-md sm:rounded-lg bg-white max-w-lg">
        <div className="p-5 text-lg font-semibold text-left text-neurtal-900 bg-white flex items-center">
          {brand.logoUrl && (
            <div className="h-32 relative w-96 mr-4">
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
            <p className="mt-1 text-sm font-normal text-neurtal-500">
              Browse a list of Flowbite products designed to help you work and
              play, stay organized, get answers, keep in touch, grow your
              business, and more.
            </p>
          </div>
        </div>
        <table className="w-full text-sm text-left text-neurtal-500">
          <tbody>
            {brand.countryOfOrigin && (
              <tr className="bg-white border-b">
                <th
                  scope="row"
                  className="px-6 py-4 font-medium text-neurtal-900 whitespace-nowrap"
                >
                  Origin
                </th>
                <td className="px-6 py-4">{brand.countryOfOrigin}</td>
              </tr>
            )}
            {brand.foundedYear && (
              <tr className="bg-white border-b">
                <th
                  scope="row"
                  className="px-6 py-4 font-medium text-neurtal-900 whitespace-nowrap"
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
                  className="px-6 py-4 font-medium text-neurtal-900 whitespace-nowrap"
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

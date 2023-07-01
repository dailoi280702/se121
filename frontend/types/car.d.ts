// Car.ts

type CarObject = {
  id: number
  brand?: Brand
  series?: Series
  name: string
  year?: number
  horsePower?: number
  torque?: number
  transmission?: Transmission
  fuelType?: FuelType
  imageUrl?: string
  review?: string
}

interface CarIndex {
  [key: string]: any
}

type Car = CarObject & CarIndex

// Brand.ts

interface Brand {
  id: number
  name: string
  countryOfOrigin?: string
  foundedYear?: number
  websiteUrl?: string
  logoUrl?: string
}

// Series.ts

type Series = {
  id: number
  name: string
  brandId: number
}

// SeriesDetail.ts

type SeriesDetail = {
  id: number
  name: string
  brand: Brand
}

// Transmission.ts

interface Transmission {
  id: number
  name: string
  description?: string
}

// FuelType.ts

interface FuelType {
  id: number
  name: string
  description?: string
}

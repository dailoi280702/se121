export default function SeriesList({
  series,
  cars,
}: {
  series: Series[]
  cars: Car[]
}) {
  const seriesMap: Map<number, Car[]> = cars.reduce((map, car) => {
    const seriesId = car.series!.id
    map.set(seriesId, map.get(seriesId) || [])
    map.get(seriesId).push(car)
    return map
  }, new Map())

  return (
    <div>
      {series.length && series.map((s) => <div key={s.id}>{s.name}</div>)}
    </div>
  )
}

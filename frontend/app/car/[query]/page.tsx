'use server'

export default async function page({ params }: { params: { query: string } }) {
  return <div>{params.query}</div>
}

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="mx-auto h-full py-8 sm:max-w-4xl md:px-4">{children}</div>
  )
}

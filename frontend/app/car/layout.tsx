export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="mx-auto mb-24 h-full sm:max-w-4xl md:px-4">{children}</div>
  )
}

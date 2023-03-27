export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div
      className="mx-auto overflow-hidden
      sm:my-14 sm:max-w-lg sm:rounded-2xl sm:bg-gray-50 sm:border sm:border-neutral-200"
    >
      {children}
    </div>
  )
}

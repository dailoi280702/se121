export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div
      className="mx-auto overflow-hidden
      sm:my-10 sm:max-w-lg sm:rounded-md sm:bg-gray-50 sm:border sm:border-neutral-200"
    >
      {children}
    </div>
  )
}

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div
      className="mt-10 mx-auto
      sm:max-w-lg sm:rounded-md sm:bg-gray-50 sm:border sm:border-neutral-200"
    >
      {children}
    </div>
  )
}

export default function Layout({ children }: { children: React.ReactNode }) {
  return <div className="h-full grid-flow-col-dense grid">{children}</div>
}

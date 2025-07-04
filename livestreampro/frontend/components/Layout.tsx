import { ReactNode } from 'react'
import Head from 'next/head'

interface LayoutProps {
  children: ReactNode
  title?: string
}

export default function Layout({ children, title = 'LiveStreamPro' }: LayoutProps) {
  return (
    <>
      <Head>
        <title>{title}</title>
      </Head>
      <header className="p-4 bg-gray-100">
        <h1 className="text-2xl">LiveStreamPro</h1>
      </header>
      <main className="p-4">{children}</main>
      <footer className="p-4 bg-gray-100 text-center text-sm">
        © {new Date().getFullYear()} LiveStreamPro
      </footer>
    </>
  )
}

import MainLayout from '../components/layout/MainLayout'
import Link from 'next/link'

export default function Home() {
  return (
    <div>
      <MainLayout>
        <div className="hero min-h-screen">
          <div className="hero-content text-center">
            <div className="max-w-md">
              <h1 className="text-5xl font-bold">Hello there</h1>
              <p className="py-6">Welcome to Coderahi where we aim to make your developer journey a smooth journey but with full of excitement
              just like a Rahi does.</p>
              <Link href="/courses"><button className="btn btn-primary">Get Started</button></Link>
            </div>
          </div>
        </div>
      </MainLayout>
    </div>
  )
}

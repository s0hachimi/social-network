import { redirect } from 'next/navigation'

export default async function Root() {

  const response = await fetch('http://localhost:8080')
  const data = await response.json()


  if (data.status) {
    redirect('/Posts')
  } else {
    redirect('/login')
  }
}
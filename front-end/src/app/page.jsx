import { redirect } from 'next/navigation'

export default async function Root() {

  let data
  
  try {
    const response = await fetch('http://localhost:8080/', {
      method: "GET",
      credentials: "include",
    })

    data = await response.json()

  } catch (error) {
    console.error(error)
    return
  }

  console.log(data);

  if (data.status) {
    redirect('/Posts')
  } else {
    redirect('/login')
  }
}
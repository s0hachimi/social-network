'use client'

import { useRouter } from 'next/navigation'
import { useEffect } from 'react'

export default function Root() {

  const router = useRouter()

  useEffect(() => {
    const fetchData = async () => {


      try {
        const response = await fetch('http://localhost:8080/', {
          method: 'GET',
          credentials: "include",
          headers: {
            'Content-Type': 'application/json',
          },
        })

        const data = await response.json()
        console.log(data)

        if (data.status) {
          router.push('/Posts')
        } else {
          router.push('/login')
        }
      } catch (error) {
        console.error("Error fetching data:", error)
        router.push('/login')
      }
    }

    fetchData()
  }, [])

  return null
}
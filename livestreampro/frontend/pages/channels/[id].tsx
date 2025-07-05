import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import Layout from '../../components/Layout'

interface Channel {
  id: string
  name: string
}

export default function ChannelPage() {
  const router = useRouter()
  const { id } = router.query
  const [channel, setChannel] = useState<Channel | null>(null)

  useEffect(() => {
    if (!id) return
    fetch(`http://localhost:8080/channels/${id}`)
      .then((res) => res.json())
      .then((data) => setChannel(data))
      .catch((err) => console.error(err))
  }, [id])

  return (
    <Layout title={`Channel ${id}`}> 
      {channel ? (
        <div>
          <h1 className="text-xl">{channel.name}</h1>
          <p>ID: {channel.id}</p>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </Layout>
  )
}

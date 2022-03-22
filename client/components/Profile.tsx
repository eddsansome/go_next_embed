import useSWR from 'swr'


async function fetcher(url: string){
  const resp = await fetch(url)
  return resp.text()
}

export default function Profile () {
  const { data, error } = useSWR('/api/users', fetcher)

  if (error) return <div>failed to load</div>
  if (!data) {
    return <div>loading...</div> } else {
      const user = JSON.parse(data)

     return <div><h1>hello {user.username}</h1></div>
    }
}
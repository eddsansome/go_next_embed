import type { NextPage } from 'next'
import Profile from "../components/Profile"

const Home: NextPage = () => {
  return (
    <div>
      <h1>Index page</h1>
      <Profile/>
    </div>
  )
}

export default Home

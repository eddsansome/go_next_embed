/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites(){
    return [
      {
        source: "/api/users",
        destination:  "http://localhost:4000/api/users"
      }
    ]

  },
  reactStrictMode: true,
}

module.exports = nextConfig

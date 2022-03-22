/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites(){
    return [
      {
        source: "/api",
        destination:  "http://localhost:4000/api"
      }
    ]

  },
  reactStrictMode: true,
}

module.exports = nextConfig

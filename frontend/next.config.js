/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  images: {
    domains: ['imageio.forbes.com'],
  },
}

module.exports = nextConfig

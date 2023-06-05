/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
    serverActions: true,
  },
  images: {
    domains: [
      'imageio.forbes.com',
      'www.ramtrucks.com',
      'www.tesla.com',
      'www.carlogos.org',
    ],
  },
}

module.exports = nextConfig

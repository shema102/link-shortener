import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api/shorten': {
        target: 'http://localhost:8080/api/shorten',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api\/shorten/, '')
      }
    },
    cors: false
  }
})

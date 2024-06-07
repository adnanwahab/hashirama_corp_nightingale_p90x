import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [tailwindcss()],
  server: {
    proxy: {
      '/api': {
        target: 'http://0.0.0.0:9000/hooks/',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
        //curl 0.0.0.0:9000/hooks/redeploy-webhook
      }
    }
  }


  //     server: {
  //   proxy: {
  //     // Proxy requests to your FastAPI app running on the Jetson Orin
  //     '/api': {
  //       target: 'http://<jetson_orin_ip>:<port>', // Replace <jetson_orin_ip> and <port> with the IP address and port of your Jetson Orin
  //       changeOrigin: true,
  //       rewrite: (path) => path.replace(/^\/api/, ''),
  //     },
  //     // Additional proxy configurations if needed
  //   },
  // },
})

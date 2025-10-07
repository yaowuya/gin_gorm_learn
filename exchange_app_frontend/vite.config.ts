import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueJsx(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  // 添加 server 配置
  server: {
    port: 5173, // 指定开发服务器端口
    host: '0.0.0.0', // 指定服务器应该监听哪个 IP 地址。设置为 '0.0.0.0' 或 true 将会监听所有地址，包括局域网和公网地址。
  }
})
